// +build  appengine
package main

import (
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type Api struct {
	client Client
	db     ContentDb
}

type PostContent struct {
	Url        string `form:"url" json:"url" binding:"required"`
	PublishWay string `form:"publish_way" json:"publish_way"`
	Date       string
	Duration   string
}

type EditContent struct {
	Kind    string
	History string
}

type RssStatus struct {
	Status  string
	Url     string `form:"url" json:"url" binding:"required"`
	Id      string `form:"id" json:"id" binding:"required"`
	History string `form:"id" json:"history"`
}

type Status struct {
	Status string
	Debug  string
}

var (
	api         Api
	matcher     *regexp.Regexp
	targetField string
)

const (
	cacheControlAge = "max-age=21600" // 3 * 3600
	defaultDuration = "12h"
	HOST_URL        = "https://tweakpods.appspot.com"
)

func (api Api) create(c echo.Context) error {
	ctx := appengine.NewContext(c.Request())
	//c.Response().Header().Add("Access-Control-Allow-Origin", "*")

	var json PostContent
	if c.Bind(&json) == nil {
		log.Debugf(ctx, "post:%v", json)
		if json.PublishWay == "" {
			json.PublishWay = "firstout"
		}
		if json.Duration == "" {
			json.Duration = defaultDuration
		}

		var rssv Rss
		var err error
		if rssv, err = api.client.GetRss(ctx, json.Url, ""); err != nil {
			return c.XML(http.StatusBadRequest, Status{Status: "couldn't request"})
		}
		if len(rssv.Channel.Items) <= 0 {
			return c.XML(http.StatusBadRequest, Status{Status: "couldn't find rss"})
		}
		if rssv.Channel.Items[0].PubDate == "" {
			return c.XML(http.StatusBadRequest, Status{Status: "couldn't find pubDate in latest item"})
		}
		json.Date = rssv.Channel.Items[0].PubDate

		if id, err := api.db.Add(json, ctx); err != nil {
			log.Debugf(ctx, "create:%v", err)
			return c.JSON(http.StatusBadRequest, Status{Status: "parse error"})
		} else {
			return c.JSON(http.StatusOK, RssStatus{Status: "ok", Id: id, Url: json.Url})
		}
	} else {
		return c.JSON(http.StatusBadRequest, Status{Status: "parse error"})
	}
}

func (api Api) get(c echo.Context) error {
	ctx := appengine.NewContext(c.Request())
	//c.Response().Header().Add("Access-Control-Allow-Origin", "*")

	log.Debugf(ctx, "get rss %v", c.Param("id"))
	if rss, err := api.db.Get(c.Param("id"), ctx); err != nil {
		return c.JSON(http.StatusBadRequest, Status{Status: "parse error"})
	} else {
		return c.JSON(http.StatusOK, RssStatus{Status: "ok", Id: rss.Id, Url: rss.Url, History: rss.History})
	}
}

func (api Api) getAll(c echo.Context) error {
	ctx := appengine.NewContext(c.Request())
	//c.Response().Header().Add("Access-Control-Allow-Origin", "*")

	if rsses, err := api.db.GetAll(100, ctx); err != nil {
		return c.JSON(http.StatusBadRequest, Status{Status: "parse error"})
	} else {
		return c.JSON(http.StatusOK, rsses)
	}
}

func (api Api) publish(c echo.Context) error {
	ctx := appengine.NewContext(c.Request())
	var rssv Rss
	var err error
	var stored Content

	if stored, err = api.db.Get(c.Param("id"), ctx); err != nil {
		return c.JSON(http.StatusBadRequest, Status{Status: "parse error"})
	}

	if rssv, err = api.client.GetRss(ctx, stored.Url, stored.LastLatestDate); err != nil {
		return c.XML(http.StatusBadRequest, "")
	}

	items := rssv.Channel.Items

	// pick up
	var p = 0
	if stored.PublishWay == "random" {
		p = rand.Intn(len(items))
	} else {
		p = len(items) - len(strings.Split(stored.History, ",")) - 1
	}

	if p >= len(items) || p < 0 {
		return c.JSON(http.StatusInternalServerError, Status{Status: "server error", Debug: fmt.Sprintf("out of index items:%v p:%v", items, p)})
	}

	// add picked up item to history
	added := stored.History + strconv.Itoa(p) + ","
	edited := EditContent{
		Kind:    "history",
		History: added,
	}
	if _, err := api.db.Edit(stored.Id, edited, ctx); err != nil {
		return c.JSON(http.StatusBadRequest, Status{Status: "edited error"})
	}

	return c.JSON(http.StatusOK, Status{Status: "ok", Debug: added})
}

func (api Api) publishAll(c echo.Context) error {
	ctx := appengine.NewContext(c.Request())
	//c.Response().Header().Add("Access-Control-Allow-Origin", "*")

	if rsses, err := api.db.GetAll(100, ctx); err != nil {
		return c.JSON(http.StatusBadRequest, Status{Status: "parse error"})
	} else {
		var errcount = 0
		for _, it := range rsses {
			d, err := time.ParseDuration(it.Duration)
			if err != nil {
				errcount = errcount + 1
				continue
			}

			if time.Now().Before(it.UpdatedAt.Add(d)) {
				var rssv Rss
				if rssv, err = api.client.GetRss(ctx, it.Url, it.LastLatestDate); err != nil {
					errcount = errcount + 1
					continue
				}

				items := rssv.Channel.Items

				// pick up
				var p = 0
				if it.PublishWay == "random" {
					p = rand.Intn(len(items))
				} else {
					p = len(items) - len(strings.Split(it.History, ",")) - 1
				}

				// add picked up item to history
				added := it.History + strconv.Itoa(p) + ","
				edited := EditContent{
					Kind:    "history",
					History: added,
				}
				if _, err := api.db.Edit(it.Id, edited, ctx); err != nil {
					errcount = errcount + 1
					continue
				}
			}
		}
		if errcount > 0 {
			return c.JSON(http.StatusBadRequest, Status{Status: "error", Debug: fmt.Sprintf("%v", errcount)})
		} else {
			return c.JSON(http.StatusOK, Status{Status: "ok", Debug: ""})
		}
	}
}

func (api Api) getRss(c echo.Context) error {
	ctx := appengine.NewContext(c.Request())
	var rssv Rss
	var err error
	var stored Content

	if stored, err = api.db.Get(c.Param("id"), ctx); err != nil {
		return c.JSON(http.StatusBadRequest, Status{Status: "parse error"})
	}

	if rssv, err = api.client.GetRss(ctx, stored.Url, stored.LastLatestDate); err != nil {
		return c.XML(http.StatusBadRequest, "")
	}

	history_ids := strings.Split(stored.History, ",")

	items := rssv.Channel.Items
	new_items, _ := rssv.ListFromHistory(history_ids)

	for i, it := range new_items {
		new_items[i].Title = "(Rebroadcast) " + it.Title
		new_items[i].Description = "(Rebroadcast) " + it.Description
	}

	d, err := time.ParseDuration(stored.Duration)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Status{Status: "duration error"})
	}

	if time.Now().After(stored.UpdatedAt.Add(d)) {
		// pick up
		var p = 0
		if stored.PublishWay == "random" {
			p = rand.Intn(len(items))
		} else {
			p = len(items) - len(strings.Split(stored.History, ",")) - 1
		}
		if p >= len(items) || p < 0 {
			return c.JSON(http.StatusInternalServerError, Status{Status: "server error", Debug: fmt.Sprintf("out of index items:%v p:%v", items, p)})
		}

		//		items[p].PubDate = time.Now().Format(time.RFC1123Z)
		items[p].Title = "(Rebroadcast) " + items[p].Title
		items[p].Description = "(Rebroadcast) " + items[p].Description

		new_items = append(new_items, items[p])
		// add picked up item to history
		added := stored.History + strconv.Itoa(p) + ","
		edited := EditContent{
			Kind:    "history",
			History: added,
		}
		if _, err := api.db.Edit(stored.Id, edited, ctx); err != nil {
			return c.JSON(http.StatusBadRequest, Status{Status: "edited error"})
		}
	}

	rssv.Channel.Items = new_items
	rssv.Channel.Link = HOST_URL + "/rss" + stored.Id

	c.Response().Header().Set("Content-Type", "application/rss+xml; charset=UTF-8")
	if cacheControlAge != "" {
		c.Response().Header().Set("Cache-Control", cacheControlAge)
	}

	return c.XML(http.StatusOK, rssv)
}

func createMux() *echo.Echo {
	api = Api{
		client: RssClient{},
		db:     ContentDb{},
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
	g.GET("/:id/publish", api.publish)
	g.GET("/all", api.getAll)
	g.GET("/all/publish", api.publishAll)

	g2 := e.Group("/rss")
	g2.GET("/:id/feed.rss", api.getRss)

	http.Handle("/", e)

	return e
}

func init() {
	targetField = os.Getenv("TARGET_FIELD")
	if targetField == "" {
		os.Exit(1)
	}
	matcher = regexp.MustCompile(os.Getenv("REGEXP"))
	if matcher == nil {
		os.Exit(1)
	}
}
