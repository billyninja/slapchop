package actions

import (
	"encoding/json"
	"fmt"
	"github.com/billyninja/slapchop/chopper"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//	"slapchop/puzzler"
	"time"

	// 3rd party
	"github.com/julienschmidt/httprouter"
)

/* Constants */
var MaxFileSize = int64(1024 * 1024 * 5) // MB
var TileSize = 64                        // pixels
var BasePath = "/tmp/slapchop/upload"

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
	path := fmt.Sprintf("%s/%s/%s", BasePath, username, chop_id)
	img, format, err := chopper.Load(file)
	tiles := chopper.Slice(*img, TileSize, format, path)
	chopper.SaveAll(tiles)

	tilesR := make([]*chopper.TileEntry, len(tiles))
	for i, t := range tiles {
		tilesR[i] = t.ToResp()
	}

	resp := chopper.CreateResponse{
		User:   "temp-todo",
		ChopId: "temp-todo",
		Href:   "temp-todo",
		Tiles:  tilesR,
	}

	json_resp, err := json.Marshal(&resp)
	if err != nil {
		log.Fatalf("errr %v", err)
	}
	println(json_resp)
	w.Write(json_resp)
	w.WriteHeader(http.StatusOK)
}

func ReadAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	username := ps.ByName("username")
	log.Printf("Requesting all slapchops for %s", username)

	path := fmt.Sprintf("%s/%s", BasePath, username)
	println(path)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}

func Read(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	username := ps.ByName("username")
	chopid := ps.ByName("chopid")
	log.Printf("Requesting %s slapchop, from %s", chopid, username)

	path := fmt.Sprintf("%s/%s/%s", BasePath, username, chopid)
	files, _ := ioutil.ReadDir(path)
	var tiles []*chopper.TileEntry
	for _, f := range files {
		fname := f.Name()
		tiles = append(tiles, &chopper.TileEntry{
			Filename: fname,
			Href:     fmt.Sprintf("%s/%s", path, fname),
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

	path := fmt.Sprintf("%s/%s/%s", BasePath, username, chopid)
	err := os.RemoveAll(path)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`ok!`))
}
