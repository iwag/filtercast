package main

import (
	"os"
	"bytes"
	"encoding/xml"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
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

var (
	xmlv Rss
	rssUrl string
)

func requestHttp(c echo.Context) (error, string) {
	ctx := appengine.NewContext(c.Request())
	resp, err := urlfetch.Client(ctx).Get(rssUrl)
	if err != nil {
		log.Errorf(ctx, err.Error(), http.StatusInternalServerError)
		return err, ""
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	return nil, buf.String()
}

func getRss(c echo.Context) error {
	err, body := requestHttp(c)
	if err != nil {
		return c.XML(http.StatusBadRequest, "")
	}

	if err := xml.Unmarshal([]byte(body), &xmlv); err != nil {
		return c.XML(http.StatusBadRequest, "")
	}

	return c.XML(http.StatusOK, xmlv)
}

func init() {
	rssUrl = os.Getenv("RSS_URL")

	e := echo.New()
	g := e.Group("/rss")
	g.Use(middleware.CORS())

	g.GET("", getRss)
	http.Handle("/", e)
}
