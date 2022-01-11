package gore

import (
	"net/http"
)

func Get(url string, header http.Header, opts ...Option) (*Response, error) {
	return New(opts...).Get(url, header)
}

func Post(url string, header http.Header, body []byte, opts ...Option) (*Response, error) {
	return New(opts...).Post(url, header, body)
}

func Put(url string, header http.Header, body []byte, opts ...Option) (*Response, error) {
	return New(opts...).Put(url, header, body)
}

func Patch(url string, header http.Header, body []byte, opts ...Option) (*Response, error) {
	return New(opts...).Patch(url, header, body)
}

func Delete(url string, header http.Header, body []byte, opts ...Option) (*Response, error) {
	return New(opts...).Delete(url, header)
}
