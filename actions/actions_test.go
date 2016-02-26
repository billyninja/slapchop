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

var ac = ActionsConfig{
	HostName:      "localhost",
	Port:          "3001",
	Host:          "localhost:3001",
	UploadDir:     "/tmp/slapcho/upload",
	MaxUploadSize: int64(1024 * 1024 * 5),
	TileSize:      64,
	PuzzlerHost:   "",
}

func TestCreate(t *testing.T) {

	req, _ := NewfileUploadRequest("/chopit/billyninja", "uploadfile", "../test/gopher.jpg")

	mw := new(mockResponseWriter)
	params := httprouter.Params{
		httprouter.Param{Key: "username", Value: "testing"},
	}

	ac.Create(mw, req, params)
}

func TestReadAll(t *testing.T) {

	req, _ := http.NewRequest("GET", "/chopit/billyninja", nil)
	mw := new(mockResponseWriter)
	params := httprouter.Params{
		httprouter.Param{Key: "username", Value: "testing"},
	}

	ac.ReadAll(mw, req, params)
}

func TestRead(t *testing.T) {

	req, _ := http.NewRequest("GET", "/chopit/billyninja", nil)
	mw := new(mockResponseWriter)
	params := httprouter.Params{
		httprouter.Param{Key: "username", Value: "testing"},
		httprouter.Param{Key: "chopid", Value: "ASDQWE1111"},
	}

	ac.Read(mw, req, params)
}

func TestDelete(t *testing.T) {

	req, _ := http.NewRequest("DELETE", "/chopit/billyninja", nil)
	mw := new(mockResponseWriter)
	params := httprouter.Params{
		httprouter.Param{Key: "username", Value: "testing"},
		httprouter.Param{Key: "chopid", Value: "ASDQWE1111"},
	}

	ac.Delete(mw, req, params)
}
