package main

import (
	"flag"
	"github.com/billyninja/slapchop/actions"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

var FlagPortNumber = flag.String("port", "3001", "HTTP port number")

func main() {
	flag.Parse()
	router := httprouter.New()

	router.POST("/chopit/:username", actions.Create)
	router.GET("/chopit/:username", actions.ReadAll)
	router.GET("/chopit/:username/:chopid", actions.Read)
	router.DELETE("/chopit/:username/:chopid", actions.Delete)

	/*	TODO
		router.GET("/random/:chopid", actions.Random)
	*/

	// Serving files as well! Who needs nginx?
	go func() {
		http.ListenAndServe(":8080",
			http.FileServer(http.Dir("/tmp/slapchop")))
	}()

	println("Running on: " + *FlagPortNumber)
	err := http.ListenAndServe(":"+*FlagPortNumber, router)
	if err != nil {
		log.Fatal(err)
	}

}
