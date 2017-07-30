package main

import(
	"bytes"
	"encoding/xml"
	"net/http"
	"google.golang.org/appengine/urlfetch"
	"google.golang.org/appengine/log"
	"golang.org/x/net/context"
	"strconv"
)

type (
	Image struct {
		Url   string `xml:"url"`
		Title string `xml:"title"`
		Link  string `xml:"link"`
	}
	Enclosure struct {
		Url    string `xml:"url,attr"`
		Type   string `xml:"type,attr"`
		Length string `xml:"length,attr"`
	}
	Item struct {
		XMLName     xml.Name  `xml:"item"`
		Link        string    `xml:"link"`
		Title       string    `xml:"title"`
		Description string    `xml:"description"`
		Copyright   string    `xml:"copyright"`
		PubDate     string    `xml:"pubDate"`
		Enclosure   Enclosure `xml:"enclosure"`
	}
	Channel struct {
		XMLName     xml.Name `xml:"channel"`
		Title       string   `xml:"title"`
		Link        string   `xml:"link"`
		Language    string   `xml:"language"`
		Copyright   string   `xml:"copyright"`
		Description string   `xml:"description"`
		Image       Image    `xml:"image"`
		Author      string   `xml:"author"`
		Items       []Item   `xml:"item"`
	}
	Rss struct {
		XMLName xml.Name `xml:"rss"`
		Channel Channel  `xml:"channel"`
	}
)


type Client interface {
    GetRss(context.Context, string, string) (Rss, error)
}

type RssClient struct{}

func (c RssClient) GetRss(ctx context.Context, url string, lastDate string) (Rss, error) {
	log.Debugf(ctx, "get rss:%v", "aaa")

	resp, err := urlfetch.Client(ctx).Get(url)
	if err != nil {
		log.Errorf(ctx, err.Error(), http.StatusInternalServerError)
		return Rss{}, err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	var rssv Rss
	if err := xml.Unmarshal([]byte(buf.String()), &rssv); err != nil {
		return Rss{}, err
	}

	if lastDate != "" {
		rssv.Channel.Items = rssv.ListBeforeDate(lastDate)
	}

	return rssv, nil
}

func (rss Rss) ListBeforeDate(lastDate string) []Item {
	items := []Item{}
	append_ := false
	for _, it := range rss.Channel.Items {
		if lastDate == it.PubDate {
			append_ = true
		} else if append_ {
			items = append(items, it)
		}
	}
	// TODO use copy(items[:i], items)
	return items
}

func (rss Rss) ListFromHistory(history_ids []string) ([]Item, []Item) {
	new_items := []Item{}
	remainder := []Item{}

	for i, it := range rss.Channel.Items {
		contain := false
		for _, hs := range history_ids {
			hi, err := strconv.Atoi(hs)
			if err==nil && hi == i { // contain
			    new_items = append(new_items, it)
				contain = true
				break
			}
		}
		if !contain {
			remainder = append(remainder , it)
		}
	}

	return new_items, remainder
}
