package main

import (
	"log"
	"flag"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"slapchop/actions"
)

var PortNumber = *flag.String("port", "3001", "HTTP port number")

func main() {

	router := httprouter.New()

	router.POST("/chopit/:username", actions.Create)
	router.GET("/chopit/:username/:chopid", actions.Read)
	
	/*	TODO

	router.GET("/random/:chopid", actions.Random)
	router.DELETE("/chopit/:chopid", actions.Delete)

	*/

	err := http.ListenAndServe(":" + PortNumber, router)
	if err != nil {
		log.Fatal(err)
	}
}
