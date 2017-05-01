package main

import (
	"encoding/xml"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	Image struct {
		Url   string `xml:"url"`
		Title string `xml:"title"`
		Link  string `xml:"link"`
	}
	Enclosure struct {
		Url string `xml:"url,attr"`
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
		Items       Item     `xml:"item"`
	}
	Rss struct {
		XMLName xml.Name `xml:"rss"`
		Channel Channel  `xml:"channel"`
	}
)

var (
	xmlv Rss
)

func getRss(c echo.Context) error {
	data := `
<rss xmlns:itunes="http://www.itunes.com/DTDs/Podcast-1.0.dtd" version="2.0">
	<channel>
		<title>title title</title>
		<link>http://example.com/</link>
		<language>ja</language>
		<copyright>iwag</copyright>
		<description>aaaaa</description>
		<image>
			<url>http://example.com/icon.jpg</url>
			<title>title title title</title>
			<link>http://example.com/</link>
		</image>
<item>
<link>
http://example.com/20170404/1
</link>
<title>
<![CDATA[title title title]]>
</title>
<description>
</description>
<pubDate>Sun, 04 Apr 2017 00:00:00 +0000</pubDate>
<copyright>iwag</copyright>
<enclosure url="http://example.com/201704041.mp3" length="0" type="audio/mpeg"/>
<itunes:summary>
</itunes:summary>
<itunes:author>iwag</itunes:author>
</item>
</channel>
</rss>
`

	if err := xml.Unmarshal([]byte(data), &xmlv); err != nil {
		return c.XML(http.StatusBadRequest, "")
	}

	return c.XML(http.StatusOK, xmlv)
}

func init() {
	e := echo.New()
	g := e.Group("/rss")
	g.Use(middleware.CORS())

	g.GET("", getRss)
	http.Handle("/", e)
}


