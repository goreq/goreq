package gore

import (
	"net/http"
)

type Gore interface {
	Get(url string, header http.Header) (*Response, error)
	Post(url string, header http.Header, body []byte) (*Response, error)
	Put(url string, header http.Header, body []byte) (*Response, error)
	Patch(url string, header http.Header, body []byte) (*Response, error)
	Delete(url string, header http.Header) (*Response, error)
	Do(*http.Request) (*Response, error)
}

type ErrorHandler func(err error)
type BeforeRequestHandler func(req *http.Request)
type AfterResponseHandler func(resp *http.Response)
