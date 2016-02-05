package actions

import (
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"image"
)

import _ "image/jpeg"

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

    // Decoding the image from the file
    img, format, err := image.Decode(file)
    if err != nil {
    	log.Println(err)
        return
    }
    
	// Getting img dimensions, that we can decide how to slice it
    bounds := img.Bounds()
    log.Println(bounds.Min.X, bounds.Max.X, bounds.Min.X, bounds.Max.Y)

    n_x := bounds.Max.X/TileSize // Number of tiles in the X axis
    n_y := bounds.Max.Y/TileSize // Number of tiles in the Y axis
    log.Println("N. Tiles: ",  n_x, n_y)
    log.Println("Succesfully loaded image!", &img, format)

}

func Read(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    log.Printf("Requesting slapchop")
    chop_id := ps.ByName("chopid")
    log.Println(chop_id)
}
