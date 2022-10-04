package goreq

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestResponse_String(t *testing.T) {
	expectedBody := "test"
	r := Response{
		&http.Response{
			Body: io.NopCloser(bytes.NewBuffer([]byte(expectedBody))),
		},
	}

	if r.String() != expectedBody {
		t.Fatalf("expected response body string %s but was %s", expectedBody, r.String())
	}
}

func TestResponse_Json(t *testing.T) {
	testBody := map[string]interface{}{
		"test": "true",
	}

	resBody, err := json.Marshal(testBody)
	if err != nil {
		t.Fatal(err)
	}

	r := Response{
		&http.Response{
			Body: io.NopCloser(bytes.NewBuffer(resBody)),
		},
	}

	var resp map[string]interface{}
	if err := r.Json(&resp); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(resp, testBody) {
		t.Fatal("error in decoded json body")
	}
}

func TestResponse_JsonWithDecodeError(t *testing.T) {
	r := Response{
		&http.Response{
			Body: io.NopCloser(bytes.NewBuffer([]byte("error"))),
		},
	}

	var resp map[string]interface{}
	if err := r.Json(&resp); err == nil {
		t.Fatal("expected error when decoding invalid json")
	}
}
