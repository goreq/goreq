package goreq

import (
	"net/http"
	"reflect"
	"testing"
	"time"
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

	if c.timeout != expectedTimeout {
		t.Fatalf("expected timeout of %v but was %v", expectedTimeout, c.timeout)
	}

	if c.baseURL != expectedBaseUrl {
		t.Fatalf("expected base url %s but was %s", expectedBaseUrl, c.baseURL)
	}

	if c.baseURL != expectedBaseUrl {
		t.Fatalf("expected base url %s but was %s", expectedBaseUrl, c.baseURL)
	}

	if !reflect.DeepEqual(c.baseHeader, expectedBaseHeader) {
		t.Fatalf("expected base header %s but was %s", expectedBaseHeader, c.baseHeader)
	}

	if !reflect.DeepEqual(c.temporaryBody, expectedBody) {
		t.Fatalf("expected base body %v but was %v", expectedBody, c.temporaryBody)
	}

	if c.errorHandler == nil {
		t.Fatal("unexpected error handler value")
	}

	if c.beforeRequestHandler == nil {
		t.Fatal("unexpected before request handler value")
	}

	if c.afterResponseHandler == nil {
		t.Fatal("unexpected after response handler value")
	}

}
