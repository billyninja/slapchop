package main

import (
	"flag"
	"github.com/billyninja/slapchop/actions"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

var PortNumber = *flag.String("port", "3001", "HTTP port number")

func main() {

	router := httprouter.New()

	router.POST("/chopit/:username", actions.Create)
	router.GET("/chopit/:username", actions.ReadAll)
	router.GET("/chopit/:username/:chopid", actions.Read)
	router.DELETE("/chopit/:username/:chopid", actions.Delete)

	/*	TODO
		router.GET("/random/:chopid", actions.Random)
	*/

	err := http.ListenAndServe(":"+PortNumber, router)
	if err != nil {
		log.Fatal(err)
	}
}
