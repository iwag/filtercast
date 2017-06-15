package main

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/aetest"
	"testing"
	"errors"
)

type (
	TestNormalClient struct{}
	TestErrorClient struct{}
)

func (c TestNormalClient) GetRss(ctx context.Context, url string) (Rss, error) {
	return Rss{}, nil
}

func (c TestErrorClient) GetRss(ctx context.Context, url string) (Rss, error) {
	return Rss{}, errors.New("errrrorrrr")
}

func testNormal(t *testing.T) {

	_, done, err := aetest.NewContext()
	if err != nil {
		t.Fatalf("aetest: %v", err)
	}
	defer done()
}
