package twotierchannel

import "fmt"


// Job represents the job to be run
type Job struct {
	Flag    int
	Payload Payload
}

// A buffered channel that we can send work requests on.
var JobQueue chan Job

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	fmt.Println(fmt.Sprintf("%p", workerPool))
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case job, ok := <-w.JobChannel:
				if !ok {
					return
				}
			// we have received a work request.
				if err := job.Payload.UploadToS3(); err != nil {

				}

			case <-w.quit:
			// we have received a signal to stop
				return
			}
		}
	}()
}


// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
