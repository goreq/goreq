package goreq

import (
	"bytes"
	"io/ioutil"

	"github.com/valyala/fasthttp"
)

type Response struct {
	*fasthttp.Response

	jsonEncoder JsonEncoderFunc
	jsonDecoder JsonDecoderFunc
}

func newResponse(c *client) *Response {
	return &Response{
		c.baseResponse,
		c.jsonEncoder,
		c.jsonDecoder,
	}
}

func (r Response) Json(v interface{}) error {
	bodyReader := bytes.NewBuffer(r.Body())
	return r.jsonDecoder(bodyReader).Decode(v)
}

func (r Response) String() string {
	bodyReader := bytes.NewBuffer(r.Body())
	body, _ := ioutil.ReadAll(bodyReader)
	return string(body)
}
