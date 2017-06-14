package main

import (
	"net/http"
	"os"
	"regexp"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"google.golang.org/appengine"
)

var (
	client  Client
	rssUrl  string
	matcher *regexp.Regexp
	targetField string
)

const (
	cacheControlAge = "" // "max-age=3600"
)

func getRss(c echo.Context) error {
	ctx := appengine.NewContext(c.Request())
	var xmlv Rss
	var err error
	if xmlv, err = client.GetRss(ctx, rssUrl); err != nil {
		return c.XML(http.StatusBadRequest, "")
	}

	items := []Item{}
	for _, i := range xmlv.Channel.Items {
		if targetField == "title" && matcher.MatchString(i.Title) {
			items = append(items, i)
		} else if targetField == "description" && matcher.MatchString(i.Description) {
			items = append(items, i)
		}
	}

	xmlv.Channel.Items = items

	if cacheControlAge != "" {
		c.Response().Header().Set("Cache-Control", cacheControlAge)
	}

	return c.XML(http.StatusOK, xmlv)
}

func init() {
	rssUrl = os.Getenv("RSS_URL")
	if (rssUrl == "") {
		os.Exit(1)
	}
	targetField = os.Getenv("TARGET_FIELD")
	if (targetField == "") {
		os.Exit(1)
	}
	matcher = regexp.MustCompile(os.Getenv("REGEXP"))
	if (matcher == nil) {
		os.Exit(1)
	}

	client = RssClient{}

	e := echo.New()
	g := e.Group("/rss")
	g.Use(middleware.CORS())

	g.GET("", getRss)
	http.Handle("/", e)
}
