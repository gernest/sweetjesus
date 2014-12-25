package main

import (
	"fmt"

	"github.com/gernest/jesus"
)

// WorkerQeue enlists the channels in which workers are listening to
var WorkerQueue chan chan []jesus.Inbox

// StartDispatcher creates the worker pool
// The parameter nworker determines the number of workers to be started
func StartDispatcher(nworkers int) {
	// First, initialize the channel we are going to but the workers' work channels into.
	WorkerQueue = make(chan chan []jesus.Inbox, nworkers)
	// Now, create all of our workers.
	for i := 0; i < nworkers; i++ {
		fmt.Println("Starting worker", i+1)
		worker := NewWorker(i+1, WorkerQueue)
		worker.Start()
	}
	go func() {
		for {
			select {
			case work := <-WorkQueue:
				fmt.Println("Received work requeust")
				go func() {
					worker := <-WorkerQueue
					fmt.Println("Dispatching work request")
					worker <- work
				}()
			}
		}
	}()
}
