package gore

import (
	"bytes"
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

func (c client) validateURL(fromUrl string) (string, error) {
	toUrl, err := url.Parse(fromUrl)
	if err != nil || toUrl.Scheme == "" || toUrl.Host == "" {
		return "", errors.New("invalid URL")
	}

	return toUrl.String(), nil
}

func (c client) buildURL(fromUrl string) string {
	builtUrl, err := c.validateURL(fromUrl)
	if err != nil {
		fromUrl = c.baseURL + fromUrl
	} else {
		fromUrl = builtUrl
	}

	return fromUrl
}

func (c client) req(reqUrl string, method string, header http.Header, body []byte) (*http.Response, error) {
	reqUrl = c.buildURL(reqUrl)
	req, err := http.NewRequest(method, reqUrl, bytes.NewBuffer(body))
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

func (c client) Get(reqUrl string, header http.Header) (*http.Response, error) {
	return c.req(reqUrl, http.MethodGet, header, nil)
}

func (c client) Post(reqUrl string, header http.Header, body []byte) (*http.Response, error) {
	return c.req(reqUrl, http.MethodPost, header, body)
}
