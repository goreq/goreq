package gore

import (
	"errors"
	"net/http"
	"net/url"
	"time"
)

type client struct {
	client     *http.Client
	timeout    time.Duration
	baseURL    string
	baseHeader http.Header
}

func New(opts ...Option) Gore {
	c := &client{}
	for _, opt := range opts {
		opt(c)
	}
	c.client = &http.Client{
		Timeout: c.timeout,
	}
	return c
}

func (c client) buildURL(fromUrl string) (string, error) {
	toUrl, err := url.Parse(fromUrl)
	if err != nil || toUrl.Scheme == "" || toUrl.Host == "" {
		return "", errors.New("invalid URL")
	}

	return toUrl.String(), nil
}

func (c client) Get(reqUrl string, header http.Header) (*http.Response, error) {
	builtUrl, err := c.buildURL(reqUrl)
	if err != nil {
		reqUrl = c.baseURL + reqUrl
	} else {
		reqUrl = builtUrl
	}

	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header = c.baseHeader
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
