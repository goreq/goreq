package goreq

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/valyala/fasthttp"
)

type client struct {
	client        *fasthttp.Client
	timeout       time.Duration
	baseURL       string
	baseHeader    *fasthttp.RequestHeader
	temporaryBody []byte
	baseRequest   *fasthttp.Request
	baseResponse  *fasthttp.Response

	errorHandler         ErrorHandler
	beforeRequestHandler BeforeRequestHandler
	afterResponseHandler AfterResponseHandler
	jsonEncoder          JsonEncoderFunc
	jsonDecoder          JsonDecoderFunc
}

type JsonEncoderFunc func(io.Writer) JsonEncoder
type JsonDecoderFunc func(io.Reader) JsonDecoder

var defaultClient = New()

func New(opts ...Option) Gore {
	c := &client{
		jsonEncoder:  defaultJsonEncoder,
		jsonDecoder:  defaultJsonDecoder,
		baseRequest:  fasthttp.AcquireRequest(),
		baseResponse: fasthttp.AcquireResponse(),
		baseHeader:   &fasthttp.RequestHeader{},
	}

	resolveOptions(c, opts...)
	c.client = &fasthttp.Client{
		// Timeout: c.timeout,
		ReadTimeout:  c.timeout,
		WriteTimeout: c.timeout,
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

func defaultJsonEncoder(w io.Writer) JsonEncoder {
	return json.NewEncoder(w)
}

func defaultJsonDecoder(r io.Reader) JsonDecoder {
	return json.NewDecoder(r)
}

func (c *client) validateURL(fromUrl string) error {
	toUrl, err := url.Parse(fromUrl)
	if err != nil || toUrl.Scheme == "" || toUrl.Host == "" {
		return errors.New("invalid URL")
	}

	return nil
}

func (c *client) buildURL(fromUrl string) *fasthttp.URI {
	err := c.validateURL(fromUrl)
	if err != nil {
		fromUrl = c.baseURL + fromUrl
	}

	uri := &fasthttp.URI{}
	uri.Update(fromUrl)

	return uri
}

func (c *client) buildReq(reqUrl string, method string, body []byte) {
	reqUrlParsed := c.buildURL(reqUrl)

	req := c.baseRequest
	req.SetURI(reqUrlParsed)
	req.Header = *c.baseHeader
	req.Header.SetMethod(method)
	req.SetBody(body)
}

func (c *client) JsonEncode(w io.Writer, v interface{}) error {
	return c.jsonEncoder(w).Encode(v)
}

func (c *client) Get(reqUrl string, opts ...Option) (*Response, error) {
	resolveOptions(c, opts...)
	c.buildReq(reqUrl, http.MethodGet, nil)
	res, err := c.Do()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) Post(reqUrl string, opts ...Option) (*Response, error) {
	resolveOptions(c, opts...)
	c.buildReq(reqUrl, http.MethodGet, nil)
	res, err := c.Do()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) Put(reqUrl string, opts ...Option) (*Response, error) {
	resolveOptions(c, opts...)
	c.buildReq(reqUrl, http.MethodGet, nil)
	res, err := c.Do()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) Patch(reqUrl string, opts ...Option) (*Response, error) {
	resolveOptions(c, opts...)
	c.buildReq(reqUrl, http.MethodGet, nil)
	res, err := c.Do()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) Delete(reqUrl string, opts ...Option) (*Response, error) {
	resolveOptions(c, opts...)
	c.buildReq(reqUrl, http.MethodGet, nil)
	res, err := c.Do()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c client) Do() (*Response, error) {
	if c.beforeRequestHandler != nil {
		c.beforeRequestHandler(c.baseRequest)
	}

	err := c.client.Do(c.baseRequest, c.baseResponse)
	if err != nil {
		if c.errorHandler != nil {
			c.errorHandler(err)
		}
		return nil, err
	}

	if c.afterResponseHandler != nil {
		c.afterResponseHandler(c.baseResponse)
	}

	return newResponse(&c), nil
}
