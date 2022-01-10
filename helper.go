package gore

import "net/http"

func Get(url string, header http.Header, opts ...Option) (*http.Response, error) {
	return New(opts...).Get(url, header)
}
