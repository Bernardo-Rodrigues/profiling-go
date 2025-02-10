package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	log.Println("Starting server on :6060")
	log.Println("pprof is available at http://localhost:6060/debug/pprof/")
	if err := http.ListenAndServe(":6060", nil); err != nil {
		log.Fatal(err)
	}
}
