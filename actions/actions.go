package actions

import (
	"log"
	"fmt"
	"time"
	"net/http"
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

    w.WriteHeader(http.StatusOK)

    // TODO write proper JSON output
	w.Write([]byte(""))
}

func Read(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    log.Printf("Requesting slapchop")
    username := ps.ByName("username")
    chop_id := ps.ByName("chopid")
    log.Println(chop_id, username)
}
