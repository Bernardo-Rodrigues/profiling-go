package main

import (
	"context"
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		time.Sleep(10 * time.Millisecond) // Simulated work
		fmt.Printf("Worker %d processed job %d\n", id, job)
	}
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		fmt.Println("Error creating trace file:", err)
		return
	}
	defer f.Close()

	trace.Start(f)
	defer trace.Stop()

	jobs := make(chan int) // Without buffer → Block producers and consumers
	var wg sync.WaitGroup

	// Only 2 fixed workers
	for i := 0; i < 2; i++ {
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

	fmt.Println("Tracing completed. Run 'go tool trace trace.out' to analyze.")
}
