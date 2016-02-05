package actions

import (
	"log"
	"net/http"
	"slapchop/chopper"
	
	// 3rd party
	"github.com/julienschmidt/httprouter"
)

/* Constants */
var MaxFileSize = int64(1024 * 1024 * 5) // MB
var TileSize = 64 // pixels

/* Let's use here the CRUD standard names */
func Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

    log.Printf("Received new slapchop!")
    r.ParseMultipartForm(MaxFileSize)

    // Opening the file
    file, _, err := r.FormFile("uploadfile")
    if err != nil {
    	log.Println(err)
        return
    }
    defer file.Close() // Setting to close filehandler at the end of this function

    img, format, err := chopper.Load(file)
    tiles := chopper.Slice(*img, TileSize, format)
    chopper.SaveAll(tiles)

    w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}

func Read(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    log.Printf("Requesting slapchop")
    chop_id := ps.ByName("chopid")
    log.Println(chop_id)
}
