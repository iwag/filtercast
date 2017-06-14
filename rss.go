package main

import(
	"bytes"
	"encoding/xml"
	"net/http"
	"google.golang.org/appengine/urlfetch"
	"google.golang.org/appengine/log"
	"golang.org/x/net/context"
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
    GetRss(context.Context, string) (Rss, error)
}

type RssClient struct{} 

func (c RssClient) GetRss(ctx context.Context, url string) (Rss, error) {
	resp, err := urlfetch.Client(ctx).Get(url)
	if err != nil {
		log.Errorf(ctx, err.Error(), http.StatusInternalServerError)
		return Rss{}, err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	var xmlv Rss
	if err := xml.Unmarshal([]byte(buf.String()), &xmlv); err != nil {
		return Rss{}, err
	}

	return xmlv, nil
}
