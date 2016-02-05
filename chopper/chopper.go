package chopper

import(
	"image"
	"image/jpeg"
	"image/draw"
	"mime/multipart"
	//"image/png"
	"log"
	"fmt"
	"os"
)

type Tile struct {
	filename string
	image *image.RGBA
	format string
}

func Load(file multipart.File) (*image.Image, string, error) {

	// Decoding the image from the file
	img, format, err := image.Decode(file)
    if err != nil {
    	log.Println(err)
        return nil, "", err
    }
    return &img, format, err
}


func (t *Tile) Save() {
	toimg, _ := os.Create(t.filename)
    defer toimg.Close()
    // TODO: Switch t.format
   	jpeg.Encode(toimg, t.image, &jpeg.Options{jpeg.DefaultQuality})
}

func SaveAll(tiles []*Tile) {
	for _, t := range tiles {
		t.Save()
	}
}

func Slice(original image.Image, tileSize int, format string) []*Tile {

	// Getting original dimensions, that we can decide how to slice it
    bounds := original.Bounds()

    n_x := bounds.Max.X/tileSize // Number of tiles in the X axis
    n_y := bounds.Max.Y/tileSize // Number of tiles in the Y axis
    log.Println("N. Tiles: ",  n_x, n_y)

    /* Since we now upfront how many tiles gonna be, we can `make` the tiles array 
       with the perfetc size. Great optimization over the dynamic append() method */
    tiles := make([]*Tile, n_x * n_y)

    // Drawing the multiple rectangles
    var idx = 0
    for x := 0; x < n_x; x++ {
    	for y := 0; y < n_y; y++ {
    		rxi := x * tileSize
    		ryi := y * tileSize
    		rxj := rxi + tileSize
			ryj := ryi + tileSize

			// Setting the cropped area
    		cropR := image.Rect(rxi, ryi, rxj, ryj)
    		cropR = original.Bounds().Intersect(cropR)
    		log.Println(cropR.Min.X, cropR.Min.Y, cropR.Max.X, cropR.Max.Y)

    		// Initializing an empty RGBA tile
    		subImg := image.NewRGBA(image.Rect(0, 0, tileSize, tileSize))
    		// Filling the new img with the source area
    		draw.Draw(subImg, subImg.Bounds(), original, cropR.Min, draw.Src)
    		
    		tiles[idx] = &Tile{
    			filename: fmt.Sprintf("upload/%d_%d.jpg", y, x),
    			image: subImg,
    			format: format,
    		}

    		idx += 1
		}
    }

    return tiles
}
