package main

import (
	"flag"
	"github.com/billyninja/slapchop/actions"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"runtime"
)

var FlagPortNumber = flag.String("port", "3001", "HTTP port number")

func InitServer(port string) chan os.Signal {
	router := httprouter.New()

	router.POST("/chopit/:username", actions.Create)
	router.GET("/chopit/:username", actions.ReadAll)
	router.GET("/chopit/:username/:chopid", actions.Read)
	router.DELETE("/chopit/:username/:chopid", actions.Delete)

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
