package http

import (
	"encoding/json"
	"errors"
	"io"
	nethttp "net/http"
	"testing"
)

func TestParseRequest_ReadError(t *testing.T) {
	expectedErr := errors.New("an error")
	req := &nethttp.Request{
		Body: newErrorReadCloser(expectedErr),
	}
	data := struct{}{}
	err := ParseRequest(req, &data)

	if err == nil {
		t.Fatalf("an error was expected to be returned when failing to read from the request body")
	}
	if err != expectedErr {
		t.Fatalf("error \"%s\" was expected but got \"%s\"", expectedErr.Error(), err.Error())
	}
}

func TestParseRequest_UnmarshalError(t *testing.T) {
	expectedErr := errors.New("json: cannot unmarshal bool into Go struct field .field of type int")
	bodyData := struct {
		Field bool `json:"field"`
	}{
		Field: true,
	}

	req := &nethttp.Request{
		Body: newDataReadCloser(bodyData),
	}
	data := struct {
		Field int `json:"field"`
	}{}
	err := ParseRequest(req, &data)

	if err == nil {
		t.Fatalf("an error was expected to be returned when failing to read from the request body")
	}
	if err.Error() != expectedErr.Error() {
		t.Fatalf("error \"%s\" was expected but got \"%s\"", expectedErr.Error(), err.Error())
	}
}

func TestParseRequest(t *testing.T) {
	bodyData := struct {
		Field int `json:"field"`
	}{
		Field: 5,
	}

	req := &nethttp.Request{
		Body: newDataReadCloser(bodyData),
	}
	data := struct {
		Field int `json:"field"`
	}{}
	err := ParseRequest(req, &data)

	if err != nil {
		t.Fatalf("unexpected error returned \"%s\"", err.Error())
	}
	if data.Field != 5 {
		t.Fatalf("field was not correctly parsed from request")
	}
}

func newErrorReadCloser(err error) io.ReadCloser {
	return &errReadCloser{err: err}
}

type errReadCloser struct {
	err error
}

func (c *errReadCloser) Read(_ []byte) (n int, err error) {
	return 0, c.err
}

func (c *errReadCloser) Close() error {
	return c.err
}

func newDataReadCloser(data interface{}) io.ReadCloser {
	return &dataReadCloser{data: data}
}

type dataReadCloser struct {
	data    interface{}
	written bool
}

func (c *dataReadCloser) Read(out []byte) (n int, err error) {
	if c.written {
		return 0, io.EOF
	}
	serialized, _ := json.Marshal(c.data)
	bytesWritten := copy(out, serialized)
	if bytesWritten >= len(serialized) {
		c.written = true
	}

	return bytesWritten, nil
}

func (c *dataReadCloser) Close() error {
	return nil
}
