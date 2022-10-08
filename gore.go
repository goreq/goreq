package goreq

import (
	"io"
	"net/http"
)

type Gore interface {
	JsonEncode(data io.Writer, v interface{}) error
	Get(url string, opts ...Option) (*Response, error)
	Post(url string, opts ...Option) (*Response, error)
	Put(url string, opts ...Option) (*Response, error)
	Patch(url string, opts ...Option) (*Response, error)
	Delete(url string, opts ...Option) (*Response, error)
	Do(*http.Request) (*Response, error)
}

type ErrorHandler func(err error)
type BeforeRequestHandler func(req *http.Request)
type AfterResponseHandler func(resp *http.Response)
