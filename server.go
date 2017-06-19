package main

import (
	"net/http"
	"os"
	"regexp"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type Api struct {
	client Client
}

type PostContent struct {
	Url string `form:"url" json:"url" binding:"required"`
}

type RssStatus struct {
	Status string
	Url string `form:"url" json:"url" binding:"required"`
	Id string `form:"id" json:"id" binding:"required"`
}

type Status struct {
	Status string
}

var (
	api         Api
	rssUrl      string
	matcher     *regexp.Regexp
	targetField string
)

const (
	cacheControlAge = "" // "max-age=3600"
)

func (api Api) create(c echo.Context) error {
	ctx := appengine.NewContext(c.Request())
	//c.Response().Header().Add("Access-Control-Allow-Origin", "*")

	var json PostContent
	if c.Bind(&json) == nil {
		log.Debugf(ctx, "post:%v", json)
		// db.Add("", json, ctx)
		return c.JSON(http.StatusOK, RssStatus{Status: "ok", Id: "111", Url: json.Url})
	} else {
		return c.JSON(http.StatusBadRequest, Status{Status: "parse error"})
	}
}

func (api Api) get(c echo.Context) error {
	ctx := appengine.NewContext(c.Request())
	//c.Response().Header().Add("Access-Control-Allow-Origin", "*")

	log.Debugf(ctx, "get rss")
	// db.Add("", json, ctx)
	return c.JSON(http.StatusOK, RssStatus{Status: "ok", Id:"111", Url: "http://example.com"})
}

func (api Api) getRss(c echo.Context) error {
	ctx := appengine.NewContext(c.Request())
	var xmlv Rss
	var err error
	if xmlv, err = api.client.GetRss(ctx, rssUrl); err != nil {
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
	if rssUrl == "" {
		os.Exit(1)
	}
	targetField = os.Getenv("TARGET_FIELD")
	if targetField == "" {
		os.Exit(1)
	}
	matcher = regexp.MustCompile(os.Getenv("REGEXP"))
	if matcher == nil {
		os.Exit(1)
	}

	api = Api{
		client: RssClient{},
	}

	e := echo.New()

	g := e.Group("/api/rss")
	g.Use(middleware.CORS())

	g.POST("/new", api.create)
	g.OPTIONS("/new", func(c echo.Context) error {
		c.Response().Header().Add("Access-Control-Allow-Origin", "*")
		c.Response().Header().Add("Access-Control-Allow-Methods", "POST")
		return c.JSON(http.StatusOK, Status{Status: "ok"})
	})
	g.GET("/:id", api.get)

	g2 := e.Group("/rss")
	g2.GET("/:id/feed.rss", api.getRss)

	http.Handle("/", e)
}
