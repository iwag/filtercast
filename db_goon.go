// https://godoc.org/github.com/mjibson/goon
package main

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mjibson/goon"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"time"
)

type Content struct {
	Id             string    `datastore:"-" goon:"id"`
	Url            string    `datastore:"url"`
	History        string    `datastore:"history"`
	PublishWay     string    `datastore:"publish_way"` // random or firstout
	CreatedAt      time.Time `datastore:"created_at"`
	UpdatedAt      time.Time `datastore:"updated_at"`
	LastLatestDate string    `datastore:"latest_date"`
	Duration       string    `datastore:"duration"`
}

type ContentDb struct {
}

func (db *ContentDb) Get(key string, c context.Context) (Content, error) {
	g := goon.FromContext(c)

	w := Content{
		Id: key,
	}
	if err := g.Get(&w); err != nil {
		log.Debugf(c, "%v", err)
		return Content{}, err
	}

	return w, nil
}

func (db *ContentDb) GetAll(max int, c context.Context) ([]Content, error) {

	filter := datastore.NewQuery("Content")

	filter = filter.Order("-created_at").Limit(max).Offset(0)

	contents := []Content{}
	g := goon.FromContext(c)
	if _, err := g.GetAll(filter, &contents); err != nil {
		log.Debugf(c, "%v", err)
		return []Content{}, err
	}

	ws := []Content{}
	for _, w := range contents {
		// if wanna use custom filter, write here
		ws = append(ws, w)
	}

	return ws, nil
}

func (db *ContentDb) GenId(content string, c context.Context) (string, error) {
	uuid, err1 := uuid.NewUUID()
	if err1 != nil {
		log.Debugf(c, "%v", err1)
		return "", err1
	}
	key := string(uuid.String()[0:12])

	return key, nil
}

func (db *ContentDb) Add(w PostContent, c context.Context) (string, error) {

	if w.Url == "" {
		return "", errors.New("empty")
	}
	if w.PublishWay != "firstout" && w.PublishWay != "random" {
		return "", errors.New("empty")
	}

	key, err1 := db.GenId(w.Url, c)
	if err1 != nil {
		log.Debugf(c, "%v", err1)
		return "", err1
	}

	g := goon.FromContext(c)

	wg := Content{
		Id:             key,
		Url:            w.Url,
		History:        "",
		PublishWay:     w.PublishWay, // "firstout or random"
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(), // todo should be null
		LastLatestDate: w.Date,
		Duration:       w.Duration,
	}

	if _, err := g.Put(&wg); err != nil {
		log.Debugf(c, "%v", err)
		return "", err
	}
	log.Debugf(c, "%v", wg)

	return key, nil
}

func (db *ContentDb) Edit(id string, ew EditContent, c context.Context) (Content, error) {

	g := goon.FromContext(c)

	w := Content{
		Id: id,
	}
	if err := g.Get(&w); err != nil {
		log.Debugf(c, "edit:%v", err)
		return Content{}, err
	}

	if ew.Kind == "history" {
		w.History = ew.History
	} else {
		return w, nil
	}

	wg := Content{
		Id:             id,
		Url:            w.Url,
		History:        ew.History,
		PublishWay:     w.PublishWay,
		CreatedAt:      w.CreatedAt,
		UpdatedAt:      time.Now(),
		LastLatestDate: w.LastLatestDate,
		Duration:       w.Duration,
	}

	if _, err := g.Put(&wg); err != nil {
		log.Debugf(c, "%v", err)
		return Content{}, err
	}

	w2, err := db.Get(id, c)
	log.Debugf(c, "updated:%v", w2)
	return w2, err
}

func (db *ContentDb) Delete(id string, c context.Context) error {
	g := goon.FromContext(c)

	w := Content{
		Id: id,
	}
	if err := g.Get(w); err != nil {
		log.Debugf(c, "couldn't find:%v", err)
		return err
	}

	wkey := new(Content)
	wkey.Id = id
	key, err := g.KeyError(wkey)
	if err != nil {
		return err
	}

	err2 := g.Delete(key)
	return err2
}
