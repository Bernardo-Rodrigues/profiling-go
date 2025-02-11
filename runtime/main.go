package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

var globalStore [][]byte

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close()

	fm, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer fm.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	runtime.GC()

	leakMemory()
	expensiveFunc()

	if err := pprof.WriteHeapProfile(fm); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}

	fmt.Println("Profiling completed. Run 'go tool pprof cpu.prof' and 'go tool pprof mem.prof' to analyze.")

}

//go:noinline
func expensiveFunc() {
	label := pprof.Labels("expensiveFunc", "sum of values at length of 10m")

	pprof.Do(context.Background(), label, func(ctx context.Context) {
		var sum float64
		for i := 0; i < 10_000_000; i++ {
			sum += rand.Float64()
		}
	})

	anotherExpensiveFunc()
}

//go:noinline
func anotherExpensiveFunc() {
	var sum int
	for i := 0; i < 1_000_000; i++ {
		sum += rand.Intn(10)
	}
}

func leakMemory() {
	for i := 0; i < 10; i++ {
		data := bytes.Repeat([]byte("x"), 1_000_000)
		globalStore = append(globalStore, data)

		time.Sleep(500 * time.Millisecond)
	}
}
