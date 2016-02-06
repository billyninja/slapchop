package actions

import (
	"log"
	"fmt"
	"time"
	"net/http"
    "io/ioutil"
    "slapchop/chopper"
	
	// 3rd party
	"github.com/julienschmidt/httprouter"
)

/* Constants */
var MaxFileSize = int64(1024 * 1024 * 5) // MB
var TileSize = 64 // pixels
var BasePath = "upload"


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

    // TODO write proper JSON output
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(""))
}


func ReadAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

    username := ps.ByName("username")
    log.Printf("Requesting all slapchops for %s", username)

    w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}


func Read(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

    username := ps.ByName("username")
    chopid := ps.ByName("chopid")
    log.Printf("Requesting %s slapchop, from %s", chopid, username)

    path = fmt.Sprintf("%s/%s/%s", BasePath, username, chopid)

    files, _ := ioutil.ReadDir(path)
    for _, f := range files {
            fmt.Println(f.Name())
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(""))
}
