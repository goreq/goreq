package gore

import "net/http"

type Gore interface {
	Get(url string, header http.Header) (*http.Response, error)
}
