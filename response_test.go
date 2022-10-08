package goreq

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/golang-must/must"
)

func TestResponse_String(t *testing.T) {
	expectedBody := "test"
	r := Response{
		&http.Response{
			Body: io.NopCloser(bytes.NewBuffer([]byte(expectedBody))),
		},
		defaultJsonEncoder,
		defaultJsonDecoder,
	}

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
		&http.Response{
			Body: io.NopCloser(bytes.NewBuffer(resBody)),
		},
		defaultJsonEncoder,
		defaultJsonDecoder,
	}

	var resp map[string]interface{}
	err = r.Json(&resp)

	must.Nil(err)
	must.Equal(resp, testBody)
}

func TestResponse_JsonWithDecodeError(t *testing.T) {
	r := Response{
		&http.Response{
			Body: io.NopCloser(bytes.NewBuffer([]byte("error"))),
		},
		defaultJsonEncoder,
		defaultJsonDecoder,
	}

	must := must.New(t)

	var resp map[string]interface{}
	err := r.Json(&resp)
	must.NotNil(err)
}
