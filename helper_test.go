package goreq

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-must/must"
)

func TestGet(t *testing.T) {
	expected := "test data"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer svr.Close()

	must := must.New(t)

	g := New()
	resp, err := g.Get(svr.URL, nil)

	must.Nil(err)
	defer resp.Body.Close()

	must.Equal(resp.String(), expected)
}
