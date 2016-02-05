package actions

import (
	"os"
	"log"
	"fmt"
	"net/http"
	"image"
	"image/jpeg"
	"image/draw"
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

    // Decoding the image from the file
    img, format, err := image.Decode(file)
    if err != nil {
    	log.Println(err)
        return
    }
    
	// Getting img dimensions, that we can decide how to slice it
    bounds := img.Bounds()

    n_x := bounds.Max.X/TileSize // Number of tiles in the X axis
    n_y := bounds.Max.Y/TileSize // Number of tiles in the Y axis
    log.Println("N. Tiles: ",  n_x, n_y)
    log.Println("Succesfully loaded image!", &img, format)

    // Drawing the multiple rectangles
    for x := 0; x < n_x; x++ {
    	for y := 0; y < n_y; y++ {
    		rxi := x * TileSize
    		ryi := y * TileSize
    		rxj := rxi + TileSize
			ryj := ryi + TileSize

    		cropR := image.Rect(rxi, ryi, rxj, ryj)
    		cropR = img.Bounds().Intersect(cropR)
    		log.Println(cropR.Min.X, cropR.Min.Y, cropR.Max.X, cropR.Max.Y)

    		subImg := image.NewRGBA(image.Rect(0, 0, TileSize, TileSize))
    		draw.Draw(subImg, subImg.Bounds(), img, cropR.Min, draw.Src)

    		tile_file_name := fmt.Sprintf("upload/%d_%d.jpg", y, x)
    		toimg, _ := os.Create(tile_file_name)
    		defer toimg.Close()

    		jpeg.Encode(toimg, subImg, &jpeg.Options{jpeg.DefaultQuality})
		}
    }

}

func Read(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    log.Printf("Requesting slapchop")
    chop_id := ps.ByName("chopid")
    log.Println(chop_id)
}
