package goreq

import (
	"net/http"
)

func Get(url string, opts ...Option) (*Response, error) {
	return New(opts...).Get(url)
}

func Post(url string, header http.Header, opts ...Option) (*Response, error) {
	return New(opts...).Post(url, opts...)
}

func Put(url string, header http.Header, opts ...Option) (*Response, error) {
	return New(opts...).Put(url, opts...)
}

func Patch(url string, header http.Header, opts ...Option) (*Response, error) {
	return New(opts...).Patch(url, opts...)
}

func Delete(url string, header http.Header, opts ...Option) (*Response, error) {
	return New(opts...).Delete(url, opts...)
}
