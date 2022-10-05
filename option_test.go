package goreq

import (
	"net/http"
	"testing"
	"time"

	"github.com/golang-must/must"
)

func TestOptions(t *testing.T) {
	var (
		expectedTimeout              = 1 * time.Second
		expectedBaseUrl              = "http://127.0.0.1/"
		expectedBaseHeader           = http.Header{"test": {"true"}}
		expectedBody                 = []byte("test")
		expectedErrHandler           = ErrorHandler(func(err error) {})
		expectedBeforeRequestHandler = BeforeRequestHandler(func(req *http.Request) {})
		expectedAfterRequestHandler  = AfterResponseHandler(func(resp *http.Response) {})
	)

	c := &client{}
	must := must.New(t)

	WithTimeout(expectedTimeout)(c)
	must.Equal(c.timeout, expectedTimeout)

	WithBaseURL(expectedBaseUrl)(c)
	must.Equal(c.baseURL, expectedBaseUrl)

	WithBaseHeader(expectedBaseHeader)(c)
	WithHeader(expectedBaseHeader)(c)
	newHeader := make(http.Header)
	for key, val := range expectedBaseHeader {
		newHeader[key] = val
	}
	for key, val := range expectedBaseHeader {
		newHeader[key] = val
	}
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
		expectedBaseHeader           = http.Header{"test": {"true"}}
		expectedBody                 = []byte("test")
		expectedErrHandler           = ErrorHandler(func(err error) {})
		expectedBeforeRequestHandler = BeforeRequestHandler(func(req *http.Request) {})
		expectedAfterRequestHandler  = AfterResponseHandler(func(resp *http.Response) {})
	)

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
	)

	must.Equal(c.timeout, expectedTimeout)
	must.Equal(c.baseURL, expectedBaseUrl)
	must.Equal(c.temporaryBody, expectedBody)
	must.NotNil(c.beforeRequestHandler)
	must.NotNil(c.afterResponseHandler)
	must.NotNil(c.errorHandler)

	newHeader := make(http.Header)
	for key, val := range expectedBaseHeader {
		newHeader[key] = val
	}
	for key, val := range expectedBaseHeader {
		newHeader[key] = val
	}
	must.Equal(c.baseHeader, newHeader)

}
