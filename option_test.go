package goreq

import (
	"testing"
	"time"

	"github.com/golang-must/must"
	"github.com/valyala/fasthttp"
)

func TestOptions(t *testing.T) {
	var (
		expectedTimeout              = 1 * time.Second
		expectedBaseUrl              = "http://127.0.0.1/"
		expectedBaseHeader           = &fasthttp.RequestHeader{}
		expectedBody                 = []byte("test")
		expectedErrHandler           = ErrorHandler(func(err error) {})
		expectedBeforeRequestHandler = BeforeRequestHandler(func(req *fasthttp.Request) {})
		expectedAfterRequestHandler  = AfterResponseHandler(func(resp *fasthttp.Response) {})
	)

	c := &client{}
	must := must.New(t)

	WithTimeout(expectedTimeout)(c)
	must.Equal(c.timeout, expectedTimeout)

	WithBaseURL(expectedBaseUrl)(c)
	must.Equal(c.baseURL, expectedBaseUrl)

	WithBaseHeader(expectedBaseHeader)(c)
	WithHeader(expectedBaseHeader)(c)
	newHeader := expectedBaseHeader
	expectedBaseHeader.VisitAll(func(key, value []byte) {
		newHeader.Set(string(key), string(value))
	})
	must.Equal(c.baseHeader, newHeader)

	WithBody(expectedBody)(c)
	must.Equal(c.temporaryBody, expectedBody)

	WithBeforeRequestHandler(expectedBeforeRequestHandler)(c)
	must.NotNil(c.beforeRequestHandler)

	WithAfterResponseHandler(expectedAfterRequestHandler)(c)
	must.NotNil(c.afterResponseHandler)

	WithHeader(expectedBaseHeader)(c)

	WithErrorHandler(expectedErrHandler)(c)
	must.NotNil(c.errorHandler)

}

func TestResolveOptions(t *testing.T) {
	var (
		expectedTimeout              = 1 * time.Second
		expectedBaseUrl              = "http://127.0.0.1/"
		expectedBaseHeader           = &fasthttp.RequestHeader{}
		expectedBody                 = []byte("test")
		expectedErrHandler           = ErrorHandler(func(err error) {})
		expectedBeforeRequestHandler = BeforeRequestHandler(func(req *fasthttp.Request) {})
		expectedAfterRequestHandler  = AfterResponseHandler(func(resp *fasthttp.Response) {})
	)

	expectedBaseHeader.Add("x-token", "akldsasklkhtrue")

	c := &client{}
	must := must.New(t)

	resolveOptions(
		c,
		WithTimeout(expectedTimeout),
		WithBaseURL(expectedBaseUrl),
		WithBaseHeader(expectedBaseHeader),
		WithHeader(expectedBaseHeader),
		WithBody(expectedBody),
		WithBeforeRequestHandler(expectedBeforeRequestHandler),
		WithAfterResponseHandler(expectedAfterRequestHandler),
		WithHeader(expectedBaseHeader),
		WithErrorHandler(expectedErrHandler),
		WithJsonEncoder(defaultJsonEncoder),
		WithJsonDecoder(defaultJsonDecoder),
	)

	must.Equal(c.timeout, expectedTimeout)
	must.Equal(c.baseURL, expectedBaseUrl)
	must.Equal(c.temporaryBody, expectedBody)
	must.NotNil(c.beforeRequestHandler)
	must.NotNil(c.afterResponseHandler)
	must.NotNil(c.errorHandler)
	must.NotNil(c.jsonEncoder)
	must.NotNil(c.jsonDecoder)

	newHeader := expectedBaseHeader
	expectedBaseHeader.VisitAll(func(key, value []byte) {
		newHeader.Set(string(key), string(value))
	})

	must.Equal(c.baseHeader, newHeader)

}
