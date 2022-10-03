package goreq

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	expected := "test data"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer svr.Close()

	g := New()
	resp, err := g.Get(svr.URL, nil)
	if err != nil {
		t.Errorf("non-expected err, got %v", err)
	}

	defer resp.Body.Close()

	if resp.String() != expected {
		t.Errorf("expected %s, got %s", expected, resp.String())
	}
}
