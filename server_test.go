package main

import (
	"google.golang.org/appengine/aetest"
	"testing"
)

func testNormal(t *testing.T) {

	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatalf("aetest: %v", err)
	}
	defer done()
}

