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
	Id         string         `datastore:"-" goon:"id"`
	Url        string         `datastore:"url"`
	History    string         `datastore:"history"`
	CreatedAt  time.Time      `datastore:"created_at"`
	UpdatedAt  time.Time      `datastore:"updated_at"`
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

	v := Content{
		Id:         key,
		Url:        w.Url,
		History:    w.History,
		CreatedAt:  w.CreatedAt,
		UpdatedAt:  w.UpdatedAt,
	}

	return v, nil
}

func (db *ContentDb) GetAll(uid string, is_review bool, duration_s string, c context.Context) ([]Content, error) {

	filter := datastore.NewQuery("Content")

	if duration_s != "" {
		_, err := time.ParseDuration(duration_s)
		if err != nil {
			log.Debugf(c, "%v duration:%v", err, duration_s)
			return []Content{}, err
		}
	}

	filter = filter.Order("-created_at").Limit(100).Offset(0)

	contents := []Content{}
	g := goon.FromContext(c)
	if _, err := g.GetAll(filter, &contents); err != nil {
		log.Debugf(c, "%v", err)
		return []Content{}, err
	}

	ws := []Content{}
	for _, w := range contents {
		v := Content{
			Id:         w.Id,
			Url:        w.Url,
			History:    w.History,
			CreatedAt:  w.CreatedAt,
			UpdatedAt:  w.UpdatedAt,
		}
		ws = append(ws, v)
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

func (db *ContentDb) Add(uid string, w PostContent, c context.Context) (string, error) {

	if w.Url == "" {
		return "", errors.New("empty")
	}

	key, err1 := db.GenId(w.Url, c)
	if err1 != nil {
		log.Debugf(c, "%v", err1)
		return "", err1
	}

	g := goon.FromContext(c)

	wg := Content{
		Id:         key,
		Url:        w.Url,
		History:    "",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if _, err := g.Put(&wg); err != nil {
		log.Debugf(c, "%v", err)
		return "", err
	}
	log.Debugf(c, "%v", wg)

	return key, nil
}

// func (db *ContentDb) Edit(id string, uid string, ew EditContent, c context.Context) (Content, error) {
//
// 	g := goon.FromContext(c)
//
// 	w := Content{
// 		Id: id,
// 	}
// 	if err := g.Get(w); err != nil {
// 		log.Debugf(c, "edit:%v", err)
// 		return Content{}, err
// 	}
//
// 	if ew.Kind != "url" {
// 		ew.Url = w.Url
// 	}
// 	if ew.Kind != "history" {
// 		ew.History = w.History
// 	}
//
// 	wg := Content{
// 		Id:         id,
// 		Url:        w.Url,
// 		History:    w.History,
// 		CreatedAt:  w.CreatedAt,
// 		UpdatedAt:  time.Now(),
// 	}
//
// 	if _, err := g.Put(&wg); err != nil {
// 		log.Debugf(c, "%v", err)
// 		return Content{}, err
// 	}
//
// 	w2, err := db.Get(id, uid, c)
// 	log.Debugf(c, "updated:%v", w2)
// 	return w2, err
// }

func (db *ContentDb) Delete(id string, uid string, c context.Context) error {
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
