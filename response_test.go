package goreq

import (
	"encoding/json"
	"testing"

	"github.com/golang-must/must"
	"github.com/valyala/fasthttp"
)

func TestResponse_String(t *testing.T) {
	expectedBody := "test"
	r := Response{
		&fasthttp.Response{},
		defaultJsonEncoder,
		defaultJsonDecoder,
	}
	r.SetBody([]byte(expectedBody))

	must := must.New(t)
	must.Equal(r.String(), expectedBody)
}

func TestResponse_Json(t *testing.T) {
	testBody := map[string]interface{}{
		"test": "true",
	}

	must := must.New(t)

	resBody, err := json.Marshal(testBody)
	must.Nil(err)

	r := Response{
		&fasthttp.Response{},
		defaultJsonEncoder,
		defaultJsonDecoder,
	}
	r.SetBody([]byte(resBody))

	var resp map[string]interface{}
	err = r.Json(&resp)

	must.Nil(err)
	must.Equal(resp, testBody)
}

func TestResponse_JsonWithDecodeError(t *testing.T) {
	r := Response{
		&fasthttp.Response{},
		defaultJsonEncoder,
		defaultJsonDecoder,
	}
	r.SetBody([]byte("resBody"))

	must := must.New(t)

	var resp map[string]interface{}
	err := r.Json(&resp)
	must.NotNil(err)
}
