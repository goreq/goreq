package gore

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

func WithAfterResponseHandler(handler AfterResponseHandler) Option {
	return func(c *client) {
		c.afterResponseHandler = handler
	}
}
