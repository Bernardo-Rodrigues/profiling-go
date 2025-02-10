package main

import (
	"bytes"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var globalStore [][]byte

func main() {
	go leakMemory()

	log.Println("Server running on :6060")
	http.ListenAndServe(":6060", nil)
}

func leakMemory() {
	for {
		data := bytes.Repeat([]byte("x"), 1_000_000)
		globalStore = append(globalStore, data)

		time.Sleep(500 * time.Millisecond)
	}
}
