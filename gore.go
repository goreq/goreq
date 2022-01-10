package gore

import "net/http"

type Client interface {
	Get(url string, header http.Header) (*http.Response, error)
}
