package goreq

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type client struct {
	client     *http.Client
	timeout    time.Duration
	baseURL    string
	baseHeader http.Header

	errorHandler         ErrorHandler
	beforeRequestHandler BeforeRequestHandler
	afterResponseHandler AfterResponseHandler
}

func New(opts ...Option) Gore {
	c := &client{}
	resolveOptions(c, opts...)
	c.client = &http.Client{
		Timeout: c.timeout,
	}

	return c
}

func resolveOptions(c *client, opts ...Option) {
	for _, opt := range opts {
		if opt != nil {
			opt(c)
		}
	}
}

func (c client) validateURL(fromUrl string) error {
	toUrl, err := url.Parse(fromUrl)
	if err != nil || toUrl.Scheme == "" || toUrl.Host == "" {
		return errors.New("invalid URL")
	}

	return nil
}

func (c client) buildURL(fromUrl string) string {
	err := c.validateURL(fromUrl)
	if err != nil {
		fromUrl = c.baseURL + fromUrl
	}

	return fromUrl
}

func (c client) req(reqUrl string, method string, header http.Header, body []byte) (*Response, error) {
	reqUrl = c.buildURL(reqUrl)
	req, err := http.NewRequest(method, reqUrl, bytes.NewBuffer(body))
	if err != nil {
		if c.errorHandler != nil {
			c.errorHandler(err)
		}
		return nil, err
	}

	if c.baseHeader != nil {
		req.Header = c.baseHeader
	}

	if c.beforeRequestHandler != nil {
		c.beforeRequestHandler(req)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		if c.errorHandler != nil {
			c.errorHandler(err)
		}
		return nil, err
	}

	if c.afterResponseHandler != nil {
		c.afterResponseHandler(resp)
	}

	return &Response{resp}, nil
}

func (c client) Get(reqUrl string, opts ...Option) (*Response, error) {
	cCopy := c
	resolveOptions(&cCopy, opts...)
	return c.req(reqUrl, http.MethodGet, cCopy.baseHeader, nil)
}

func (c client) Post(reqUrl string, body []byte, opts ...Option) (*Response, error) {
	cCopy := c
	resolveOptions(&cCopy, opts...)
	return c.req(reqUrl, http.MethodPost, cCopy.baseHeader, body)
}

func (c client) Put(reqUrl string, body []byte, opts ...Option) (*Response, error) {
	cCopy := c
	resolveOptions(&cCopy, opts...)
	return c.req(reqUrl, http.MethodPut, cCopy.baseHeader, body)
}

func (c client) Patch(reqUrl string, body []byte, opts ...Option) (*Response, error) {
	cCopy := c
	resolveOptions(&cCopy, opts...)
	return c.req(reqUrl, http.MethodPatch, cCopy.baseHeader, body)
}

func (c client) Delete(reqUrl string, opts ...Option) (*Response, error) {
	cCopy := c
	resolveOptions(&cCopy, opts...)
	return c.req(reqUrl, http.MethodDelete, cCopy.baseHeader, nil)
}

func (c client) Do(req *http.Request) (*Response, error) {
	var body []byte
	var err error

	if req.Body != nil {
		body, err = ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
	}
	return c.req(req.URL.String(), req.Method, req.Header, body)
}
