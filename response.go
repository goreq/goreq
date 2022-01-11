package gore

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	*http.Response
}

func (r Response) Json(v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
