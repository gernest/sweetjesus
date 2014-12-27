package main

import (
	"fmt"

	"github.com/gernest/jesus"
	"github.com/kr/pretty"
)

// NewWorker is for ibitializing a worker object
func NewWorker(id int, workerQueue chan chan []jesus.Inbox) Worker {
	// Create, and return the worker.
	worker := Worker{
		ID:          id,
		Work:        make(chan []jesus.Inbox),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool)}
	return worker
}

// Worker perfoms simple tasts
type Worker struct {
	ID          int
	Work        chan []jesus.Inbox
	WorkerQueue chan chan []jesus.Inbox
	QuitChan    chan bool
}

// Start runs a go routine which is ready for accepting tasks
func (w Worker) Start() {
	go func() {
		for {
			// Add ourselves into the worker queue.
			w.WorkerQueue <- w.Work
			select {
			case work := <-w.Work:
				// Receive a work request.
				var uProfile jesus.UProfile
				remoteDBConn, err := RemoteDB()
				pretty.Println(work)
				if err != nil {
					w.Stop()
				}
				for _, v := range work {
					fmt.Println("working on deposits for ", v.SenderNumber)
					uProfile = jesus.UProfile{}
					err = remoteDBConn.Where(&jesus.UProfile{Phone: v.SenderNumber}).First(&uProfile).Error
					if err != nil {
						w.Stop()
					}
					uProfile.PrepareDeposit(&v)
					remoteDBConn.Save(&uProfile)
					fmt.Println("====done===")
				}

			case <-w.QuitChan:
				fmt.Printf("worker%d stopping\n", w.ID)
				return
			}
		}
	}()
}

// Stop tells the worker to stop listening for work requests.
//
// Note that the worker will only stop *after* it has finished its work.
func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}
