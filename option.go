package goreq

import (
	"net/http"
	"time"
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

func WithBaseHeader(header http.Header) Option {
	return func(c *client) {
		c.baseHeader = header
	}
}

func WithHeader(header http.Header) Option {
	return func(c *client) {
		newHeader := make(http.Header)
		for key, val := range c.baseHeader {
			newHeader[key] = val
		}
		for key, val := range header {
			newHeader[key] = val
		}

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
		c.afterResponseHandler = func(resp *http.Response) {
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
