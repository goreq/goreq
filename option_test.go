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
	WithTimeout(expectedTimeout)(c)
	WithBaseURL(expectedBaseUrl)(c)
	WithBaseHeader(expectedBaseHeader)(c)
	WithBody(expectedBody)(c)
	WithErrorHandler(expectedErrHandler)(c)
	WithBeforeRequestHandler(expectedBeforeRequestHandler)(c)
	WithAfterResponseHandler(expectedAfterRequestHandler)(c)

	must := must.New(t)
	must.Equal(c.timeout, expectedTimeout)
	must.Equal(c.baseURL, expectedBaseUrl)
	must.Equal(c.baseHeader, expectedBaseHeader)
	must.Equal(c.temporaryBody, expectedBody)
	must.NotNil(c.errorHandler)
	must.NotNil(c.beforeRequestHandler)
	must.NotNil(c.afterResponseHandler)

}
