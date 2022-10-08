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

func TestClientEncode(t *testing.T) {
	input := map[string]int{"id": 1}

	must := must.New(t)

	g := New()

	var expected = new(bytes.Buffer)
	json.NewEncoder(expected).Encode(input)

	var result = new(bytes.Buffer)
	err := g.JsonEncode(result, input)
	must.Nil(err)
	must.Equal(expected, result)
}

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

func BenchmarkClientDefaultJsonEncoder(b *testing.B) {
	expected := `[{"id": 1},{"id": 2},{"id": 3},{"id": 4},{"id": 5}]`
	g := New()

	type Data struct {
		ID int `json:"id"`
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var data []Data
		g.JsonEncode(bytes.NewBuffer([]byte(expected)), &data)
	}

}

func BenchmarkClientDefaultJsonDecoder(b *testing.B) {
	expected := `[{"id": 1},{"id": 2},{"id": 3},{"id": 4},{"id": 5}]`
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer svr.Close()

	type Data struct {
		ID int `json:"id"`
	}

	g := New()
	res, _ := g.Get(svr.URL)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var data []Data
		res.Json(&data)
	}

}
