package main

import (
	"fmt"
	"time"
)

func worker(id int, tasks chan string, results chan string) {
	task:=<-tasks
		fmt.Printf("Worker %d processing task: %s\n", id, task)
		time.Sleep(time.Second) // Simulate task processing delay
		results <- fmt.Sprintf("Worker %d finished task: %s", id, task) // sending result back to results channel
	
}

func deadlock() {
	tasks := make(chan string)  // Unbuffered channel for tasks
	results := make(chan string) // Unbuffered channel for results

	// Launch 3 workers
	for i := 1; i <= 5; i++ {
		go worker(i, tasks, results)
	}

	// Send 5 tasks to workers
	for i := 1; i <= 5; i++ {
		task := fmt.Sprintf("Task %d", i)
		fmt.Println("Sending", task)
		tasks <- task // Send task to workers (will block until a worker receives it)

	}

	// Try to receive the results from the workers
	for i := 1; i <= 5; i++ {

		fmt.Println(<-results) // Receive results from workers
	}
	close(tasks) // Close the tasks channel (will signal the workers to stop)
}

// Breakdown of the Deadlock:
// Unbuffered Channels:

// tasks is an unbuffered channel, meaning that when you send a task (tasks <- task), the send will block until one of the workers receives the task.

// results is also an unbuffered channel, so workers sending results (results <-) will block until the main goroutine receives the result.


// 3 Workers vs. 5 Tasks:

// You launch 3 workers, but you send 5 tasks.

// The first 3 tasks will be picked up by the 3 workers. At this point, the main goroutine is blocked because it is sending the 4th task to the tasks channel. However, all 3 workers are busy processing the first 3 tasks, so there is no worker available to pick up the 4th task.

// Deadlock happens here: the main goroutine is blocked trying to send the 4th task, while the workers are blocked trying to send their results to the results channel (because the main goroutine hasn’t started receiving the results yet).

// Main Goroutine Blocking:

// After the main goroutine sends the 3rd task, it immediately tries to send the 4th task. But since all workers are busy processing the first 3 tasks, the main goroutine can’t send the 4th task until one of the workers becomes free.

// Meanwhile, the workers can't send their results because the main goroutine is still blocked trying to send the 4th task, so it never reaches the point where it starts receiving the results.

// Both are blocked: workers are blocked waiting for the main goroutine to receive results, and the main goroutine is blocked trying to send more tasks.

// This circular blocking leads to deadlock.