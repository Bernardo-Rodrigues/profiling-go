package main

import (
	"bytes"
	"log"
	"math"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var globalStore [][]byte

func main() {
	go leakMemory()
	go startCPUHeavyRoutine()

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

func startCPUHeavyRoutine() {
	for {
		doHeavyComputation()
		time.Sleep(500 * time.Millisecond)
	}
}

func doHeavyComputation() {
	sum := 0.0
	for i := 0; i < 1000000; i++ {
		sum += math.Sin(float64(i))
	}
}
