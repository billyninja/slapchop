package actions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/billyninja/slapchop/chopper"
	"github.com/billyninja/slapchop/puzzler"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	// 3rd party
	"github.com/julienschmidt/httprouter"
)

/* Constants */
var MaxFileSize = int64(1024 * 1024 * 5) // MB
var TileSize = 64                        // pixels

var UploadDir = "/tmp/slapchop/upload"

/* Let's use here the CRUD standard names */
func Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	username := ps.ByName("username")
	log.Printf("Received new slapchop for user: %s!", username)

	r.ParseMultipartForm(MaxFileSize)
	// Opening the file
	file, _, err := r.FormFile("uploadfile")
	if err != nil {
		log.Println(err)
		return
	}
	// Setting to close filehandler at the end of this function
	defer file.Close()
	chop_id := time.Now().Format("020106150405")
	path := fmt.Sprintf("%s/%s/%s", UploadDir, username, chop_id)
	img, format, err := chopper.Load(file)
	tiles := chopper.Slice(*img, TileSize, format, path)
	chopper.SaveAll(tiles)

	href_base := fmt.Sprintf("/upload/%s/%s", username, chop_id)

	tilesR := make([]*chopper.TileEntry, len(tiles))
	for i, t := range tiles {
		tilesR[i] = t.ToResp(href_base)
	}

	resp := chopper.CreateResponse{
		User:   username,
		ChopId: chop_id,
		Href:   fmt.Sprintf("/chopit/%s/%s", username, chop_id),
		Tiles:  tilesR,
	}

	json_resp, err := json.Marshal(&resp)
	if err != nil {
		log.Fatalf("errr %v", err)
	}

	status, response, err := puzzler.CreatePuzzle(username, tilesR)
	println(status, response, err)

	w.Write(json_resp)
	w.WriteHeader(http.StatusOK)
}

func ReadAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	username := ps.ByName("username")
	log.Printf("Requesting all slapchops for %s", username)

	path := fmt.Sprintf("%s/%s", UploadDir, username)
	println(path)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}

func Read(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	username := ps.ByName("username")
	chopid := ps.ByName("chopid")
	log.Printf("Requesting %s slapchop, from %s", chopid, username)

	path := fmt.Sprintf("%s/%s/%s", UploadDir, username, chopid)
	files, _ := ioutil.ReadDir(path)
	var tiles []*chopper.TileEntry
	for _, f := range files {
		fname := f.Name()
		tiles = append(tiles, &chopper.TileEntry{
			Filename: fname,
			Href:     fmt.Sprintf("%s/%s/%s", UploadDir, username, chopid, fname),
		})
	}

	resp := chopper.ReadResponse{
		User:  username,
		Id:    chopid,
		Tiles: tiles,
	}

	w.WriteHeader(http.StatusOK)
	json_resp, _ := json.Marshal(&resp)
	w.Write(json_resp)
}

func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	username := ps.ByName("username")
	chopid := ps.ByName("chopid")
	log.Printf("Deleting %s slapchop, from %s", chopid, username)

	path := fmt.Sprintf("%s/%s/%s", UploadDir, username, chopid)
	err := os.RemoveAll(path)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`ok!`))
}

//TEST HELPER: Creates a new file upload http request with optional extra params
func NewfileUploadRequest(uri string, paramName, path string) (*http.Request, error) {

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
