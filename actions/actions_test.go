package actions

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"testing"
)

type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header) {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockResponseWriter) WriteHeader(int) {}

func TestCreate(t *testing.T) {

	req, _ := NewfileUploadRequest("/chopit/billyninja", "uploadfile", "../test/gopher.jpg")

	mw := new(mockResponseWriter)
	params := httprouter.Params{
		httprouter.Param{Key: "username", Value: "testing"},
	}

	Create(mw, req, params)
}

func TestReadAll(t *testing.T) {

	req, _ := http.NewRequest("GET", "/chopit/billyninja", nil)
	mw := new(mockResponseWriter)
	params := httprouter.Params{
		httprouter.Param{Key: "username", Value: "testing"},
	}

	ReadAll(mw, req, params)
}

func TestRead(t *testing.T) {

	req, _ := http.NewRequest("GET", "/chopit/billyninja", nil)
	mw := new(mockResponseWriter)
	params := httprouter.Params{
		httprouter.Param{Key: "username", Value: "testing"},
		httprouter.Param{Key: "chopid", Value: "ASDQWE1111"},
	}

	Read(mw, req, params)
}

func TestDelete(t *testing.T) {

	req, _ := http.NewRequest("DELETE", "/chopit/billyninja", nil)
	mw := new(mockResponseWriter)
	params := httprouter.Params{
		httprouter.Param{Key: "username", Value: "testing"},
		httprouter.Param{Key: "chopid", Value: "ASDQWE1111"},
	}

	Delete(mw, req, params)
}
