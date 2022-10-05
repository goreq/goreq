package goreq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-must/must"
)

func TestClientGet(t *testing.T) {
	expected := "test data"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer svr.Close()

	must := must.New(t)

	g := New()
	resp, err := g.Get(svr.URL)

	must.Nil(err)
	defer resp.Body.Close()

	must.Equal(resp.String(), expected)
}

func TestClientPost(t *testing.T) {
	expected := "test data"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer svr.Close()

	must := must.New(t)

	g := New()
	resp, err := g.Post(svr.URL)

	must.Nil(err)
	defer resp.Body.Close()

	must.Equal(resp.String(), expected)
}

func TestClientPut(t *testing.T) {
	expected := "test data"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer svr.Close()

	must := must.New(t)

	g := New()
	resp, err := g.Put(svr.URL)

	must.Nil(err)
	defer resp.Body.Close()

	must.Equal(resp.String(), expected)
}

func TestClientPatch(t *testing.T) {
	expected := "test data"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer svr.Close()

	must := must.New(t)

	g := New()
	resp, err := g.Patch(svr.URL)

	must.Nil(err)
	defer resp.Body.Close()

	must.Equal(resp.String(), expected)
}

func TestClientDelete(t *testing.T) {
	expected := "test data"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer svr.Close()

	must := must.New(t)

	g := New()
	resp, err := g.Delete(svr.URL)

	must.Nil(err)
	defer resp.Body.Close()

	must.Equal(resp.String(), expected)
}

func TestClientDo(t *testing.T) {
	expected := "test data"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer svr.Close()

	must := must.New(t)

	body, err := json.Marshal(map[string]string{
		"test": "data",
	})
	must.Nil(err)

	req, err := http.NewRequest(http.MethodPost, svr.URL, bytes.NewBuffer(body))
	must.Nil(err)

	g := New()
	resp, err := g.Do(req)
	must.Nil(err)

	defer resp.Body.Close()

	must.Equal(resp.String(), expected)
}

func TestClientValidateURL(t *testing.T) {

	g := &client{}
	must := must.New(t)

	err := g.validateURL("")
	must.NotNil(err)

}

func TestClientBuildURL(t *testing.T) {

	g := &client{
		baseURL: "base.com",
	}
	must := must.New(t)

	result := g.buildURL("/api")
	must.Equal(result, "base.com/api")

}

func TestClientReq(t *testing.T) {

	expected := "test data"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer svr.Close()

	g := &client{
		client: &http.Client{
			Timeout: time.Minute,
		},
	}
	must := must.New(t)
	header := http.Header{
		"x-token": []string{"adjlajsdkexample"},
	}

	WithBaseHeader(header)(g)
	WithErrorHandler(func(err error) {})(g)
	WithBeforeRequestHandler(func(req *http.Request) {})(g)
	WithAfterResponseHandler(func(resp *http.Response) {})(g)

	res, err := g.req(svr.URL, http.MethodGet, header, []byte(""))
	must.Nil(err)
	must.Equal(res.String(), expected)
}
