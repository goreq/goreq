package goreq

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-must/must"
)

func TestHelperGet(t *testing.T) {
	expected := "test data"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer svr.Close()

	must := must.New(t)

	resp, err := Get(svr.URL)

	must.Nil(err)
	defer resp.Body.Close()

	must.Equal(resp.String(), expected)
}

func BenchmarkHelperGet(b *testing.B) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "test data")
	}))
	defer svr.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Get(svr.URL)
	}

}

func TestHelperPost(t *testing.T) {
	expected := "test data"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer svr.Close()

	must := must.New(t)

	resp, err := Post(svr.URL)

	must.Nil(err)
	defer resp.Body.Close()

	must.Equal(resp.String(), expected)
}

func BenchmarkHelperPost(b *testing.B) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "test data")
	}))
	defer svr.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Post(svr.URL)
	}

}

func TestHelperPut(t *testing.T) {
	expected := "test data"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer svr.Close()

	must := must.New(t)

	resp, err := Put(svr.URL)

	must.Nil(err)
	defer resp.Body.Close()

	must.Equal(resp.String(), expected)
}

func BenchmarkHelperPut(b *testing.B) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "test data")
	}))
	defer svr.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Put(svr.URL)
	}

}

func TestHelperPatch(t *testing.T) {
	expected := "test data"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer svr.Close()

	must := must.New(t)

	resp, err := Patch(svr.URL)

	must.Nil(err)
	defer resp.Body.Close()

	must.Equal(resp.String(), expected)
}

func BenchmarkHelperPatch(b *testing.B) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "test data")
	}))
	defer svr.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Patch(svr.URL)
	}

}

func TestHelperDelete(t *testing.T) {
	expected := "test data"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer svr.Close()

	must := must.New(t)

	resp, err := Delete(svr.URL)

	must.Nil(err)
	defer resp.Body.Close()

	must.Equal(resp.String(), expected)
}

func BenchmarkHelperDelete(b *testing.B) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "test data")
	}))
	defer svr.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Delete(svr.URL)
	}

}
