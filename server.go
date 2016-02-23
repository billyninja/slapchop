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

	println("Running on: " + *FlagPortNumber)
	err := http.ListenAndServe(":"+*FlagPortNumber, router)
	if err != nil {
		log.Fatal(err)
	}
}
