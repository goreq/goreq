package goreq

import (
	"time"

	"github.com/valyala/fasthttp"
)

type Option func(*client)

func WithTimeout(timeout time.Duration) Option {
	return func(c *client) {
		c.timeout = timeout
	}
}

func WithBaseURL(url string) Option {
	return func(c *client) {
		c.baseURL = url
	}
}

func WithBaseHeader(header *Header) Option {
	return func(c *client) {
		c.baseHeader = header
	}
}

func WithHeader(header *Header) Option {
	return func(c *client) {
		newHeader := c.baseHeader
		header.VisitAll(func(key, value []byte) {
			newHeader.SetBytesKV(key, value)
		})

		c.baseHeader = newHeader
	}
}

func WithBody(body []byte) Option {
	return func(c *client) {
		c.temporaryBody = body
	}
}

func WithErrorHandler(handler ErrorHandler) Option {
	return func(c *client) {
		c.errorHandler = handler
	}
}

func WithBeforeRequestHandler(handler BeforeRequestHandler) Option {
	return func(c *client) {
		c.beforeRequestHandler = handler
	}
}

func WithAfterResponseHandler(handlers ...AfterResponseHandler) Option {
	return func(c *client) {
		c.afterResponseHandler = func(resp *fasthttp.Response) {
			for _, handler := range handlers {
				handler(resp)
			}
		}
	}
}

func WithJsonEncoder(encoder JsonEncoderFunc) Option {
	return func(c *client) {
		c.jsonEncoder = encoder
	}
}

func WithJsonDecoder(decoder JsonDecoderFunc) Option {
	return func(c *client) {
		c.jsonDecoder = decoder
	}
}
