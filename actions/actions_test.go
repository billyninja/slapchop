package actions

import (
	"bytes"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
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

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, fi.Name())
	if err != nil {
		return nil, err
	}
	part.Write(fileContents)
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uri, body)
	request.Header.Add("Content-Type", writer.FormDataContentType())

	return request, err
}

func TestCreate(t *testing.T) {

	req, _ := newfileUploadRequest("/chopit/billyninja", "uploadfile", "../test/gopher.jpg")

	mw := new(mockResponseWriter)
	params := httprouter.Params{
		httprouter.Param{"username", "testing"},
	}

	Create(mw, req, params)
}

func TestReadAll(t *testing.T) {

	req, _ := http.NewRequest("GET", "/chopit/billyninja", nil)
	mw := new(mockResponseWriter)
	params := httprouter.Params{
		httprouter.Param{"username", "testing"},
	}

	ReadAll(mw, req, params)
}

func TestRead(t *testing.T) {

	req, _ := http.NewRequest("GET", "/chopit/billyninja", nil)
	mw := new(mockResponseWriter)
	params := httprouter.Params{
		httprouter.Param{"username", "testing"},
		httprouter.Param{"chopid", "ASDQWE1111"},
	}

	Read(mw, req, params)
}

func TestDelete(t *testing.T) {

	req, _ := http.NewRequest("DELETE", "/chopit/billyninja", nil)
	mw := new(mockResponseWriter)
	params := httprouter.Params{
		httprouter.Param{"username", "testing"},
		httprouter.Param{"chopid", "ASDQWE1111"},
	}

	Delete(mw, req, params)
}
