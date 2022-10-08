package goreq

import (
	"io/ioutil"
	"net/http"
)

type Response struct {
	*http.Response

	jsonEncoder JsonEncoderFunc
	jsonDecoder JsonDecoderFunc
}

func (r Response) Json(v interface{}) error {
	return r.jsonDecoder(r.Body).Decode(v)
}

func (r Response) String() string {
	body, _ := ioutil.ReadAll(r.Body)
	return string(body)
}
