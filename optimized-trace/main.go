package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		time.Sleep(10 * time.Millisecond)
		trace.Log(context.Background(), "Worker Processing", fmt.Sprintf("Worker %d processing job %d", id, job))
	}
}

func main() {
	f, err := os.Create("trace_optimized.out")
	if err != nil {
		fmt.Println("Error creating trace file:", err)
		return
	}
	defer f.Close()

	trace.Start(f)
	defer trace.Stop()

	jobs := make(chan int, 5) // With buffer → Reduces contention
	var wg sync.WaitGroup

	workerCount := runtime.NumCPU() // Uses available cores
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	ctx, task := trace.NewTask(context.Background(), "Sending Jobs")
	defer task.End()

	for j := 0; j < 10; j++ {
		trace.Log(ctx, "Job Sent", fmt.Sprintf("Job %d sent", j))
		jobs <- j
	}
	close(jobs)

	wg.Wait()

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Tracing completed. Run 'go tool trace trace_optimized.out' to analyze.")
}
