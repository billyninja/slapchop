package main

import (
	"flag"
	"fmt"
	"github.com/billyninja/slapchop/actions"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"runtime"
)

var FlagHost = flag.String("host", "localhost", "The host address which it will be visible")
var FlagPuzzlerHost = flag.String("puzzler", "", "Puzzler Service remote url")
var FlagPortNumber = flag.String("port", "3001", "HTTP port number")
var FlagTileSize = flag.Int("tile", 64, "Tile Size in pixels")
var FlagMaxUploadSize = flag.Int64("size", int64(1024*1024*5), "Max upload file size in BYTES")
var FlagUploadDir = flag.String("dir", "/tmp/slapchop/upload", "Local path to the uploaded files")
var FlagTemplatePath = flag.String("template", "notsetted.html", "Absolute path to the preview.html file")

func InitServer(port string) chan os.Signal {
	router := httprouter.New()

	ac := actions.ActionsConfig{
		HostName:      *FlagHost,
		Port:          *FlagPortNumber,
		Host:          fmt.Sprintf("%s:%s", *FlagHost, *FlagPortNumber),
		UploadDir:     *FlagUploadDir,
		MaxUploadSize: *FlagMaxUploadSize,
		TileSize:      *FlagTileSize,
		PuzzlerHost:   *FlagPuzzlerHost,
		TemplatePath:  *FlagTemplatePath,
	}

	router.POST("/chopit/:username", ac.Create)
	router.GET("/chopit/:username", ac.ReadAll)
	router.GET("/chopit/:username/:chopid", ac.Read)
	router.DELETE("/chopit/:username/:chopid", ac.Delete)
	router.GET("/tiled/:username/:chopid", ac.Preview)
	router.GET("/random/:username/:chopid", ac.Random)

	runtime.GOMAXPROCS(1)

	// Serving files as well! Who needs nginx?
	go func() {
		http.ListenAndServe(":8080",
			http.FileServer(http.Dir("/tmp/slapchop")))
	}()

	go func() {
		println("Running on: " + port)
		err := http.ListenAndServe(":"+port, router)
		if err != nil {
			log.Fatal(err)
		}
	}()

	LiveChan := make(chan os.Signal, 1)
	signal.Notify(LiveChan, os.Interrupt, syscall.SIGTERM)

	return LiveChan
}

func main() {
	flag.Parse()
	c := InitServer(*FlagPortNumber)
	<-c
}
