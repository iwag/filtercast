package main

import (
	"errors"
	"github.com/labstack/echo"
	"golang.org/x/net/context"
	"google.golang.org/appengine/aetest"
	"net/http/httptest"
	"testing"
)

type (
	TestNormalClient struct{}
	TestErrorClient  struct{}
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

	api = Api{
		client: TestNormalClient{},
	}

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = api.getRss(c)
	if err != nil {
		t.Fatalf("status: %v", err)
	}
}
