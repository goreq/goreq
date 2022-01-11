package gore

import (
	"net/http"
)

type Gore interface {
	Get(url string, header http.Header) (*Response, error)
	Post(url string, header http.Header, body []byte) (*Response, error)
}

type ErrorHandler func(err error)
type BeforeRequestHandler func(req *http.Request)
type AfterResponseHandler func(resp *http.Response)
