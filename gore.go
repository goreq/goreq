package goreq

import (
	"io"

	"github.com/valyala/fasthttp"
)

type Gore interface {
	JsonEncode(data io.Writer, v interface{}) error
	Get(url string, opts ...Option) (*Response, error)
	Post(url string, opts ...Option) (*Response, error)
	Put(url string, opts ...Option) (*Response, error)
	Patch(url string, opts ...Option) (*Response, error)
	Delete(url string, opts ...Option) (*Response, error)
	Do() (*Response, error)
}

type ErrorHandler func(err error)
type BeforeRequestHandler func(req *Request)
type AfterResponseHandler func(resp *Response)

type Request struct {
	*fasthttp.Request
}

type Header struct {
	*fasthttp.RequestHeader
}
