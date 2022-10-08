package goreq

func Get(url string, opts ...Option) (*Response, error) {
	return defaultClient.Get(url, opts...)
}

func Post(url string, opts ...Option) (*Response, error) {
	return defaultClient.Post(url, opts...)
}

func Put(url string, opts ...Option) (*Response, error) {
	return defaultClient.Put(url, opts...)
}

func Patch(url string, opts ...Option) (*Response, error) {
	return defaultClient.Patch(url, opts...)
}

func Delete(url string, opts ...Option) (*Response, error) {
	return defaultClient.Delete(url, opts...)
}
