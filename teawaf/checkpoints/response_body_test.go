package checkpoints

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestResponseBodyCheckpoint_ResponseValue(t *testing.T) {
	resp := new(http.Response)
	resp.StatusCode = 200
	resp.Header = http.Header{}
	resp.Header.Set("Hello", "World")
	resp.Body = ioutil.NopCloser(bytes.NewBuffer([]byte("Hello, World")))

	checkpoint := new(ResponseBodyCheckpoint)
	t.Log(checkpoint.ResponseValue(nil, resp, "", nil))

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("after read:", string(data))
}
