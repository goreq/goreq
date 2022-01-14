package goreq

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	*http.Response
}

func (r Response) Json(v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func (r Response) String() string {
	body, _ := ioutil.ReadAll(r.Body)
	return string(body)
}
