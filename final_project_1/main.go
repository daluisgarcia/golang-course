package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type Job struct {
	Name   string
	Delay  time.Duration
	Number int
}

type Worker struct {
	Id         int
	JobQueue   chan Job
	WorkerPool chan chan Job
	QuitChan   chan bool
}

func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		Id:         id,
		JobQueue:   make(chan Job),
		WorkerPool: workerPool,
		QuitChan:   make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobQueue // The worker is adding its job queue to the worker pool. This worker pool is the same that the dispatcher has

			select {
			case job := <-w.JobQueue: // receive a job request
				fmt.Printf("Worker with id %d started %s\n", w.Id, job.Name)
				fib := Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("Worker with id %d finished with result %d\n", w.Id, fib)
			case <-w.QuitChan: // we have received a signal to stop
				fmt.Printf("Worker with id %d stopped\n", w.Id)
			}

		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true // Sends a signal to stop the worker
	}()
}

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
	JobQueue   chan Job
}

func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	return &Dispatcher{
		WorkerPool: make(chan chan Job, maxWorkers), // Limits the number of workers to be executed at the same time
		MaxWorkers: maxWorkers,
		JobQueue:   jobQueue,
	}
}

func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			go func() {
				workerJobQueue := <-d.WorkerPool
				workerJobQueue <- job
			}()
		}
	}
}

func (d *Dispatcher) Run() {
	for i := 1; i <= d.MaxWorkers; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()
	}

	go d.Dispatch()
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

type JobRequest struct {
	Delay string
	Name  string
	Value int
}

func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	body, _ := io.ReadAll(r.Body)

	var jobRequest JobRequest

	if err := json.Unmarshal(body, &jobRequest); err != nil {
		errorSolutionInfo := strings.Split(strings.Split(err.Error(), ".")[1], " ") // Getting important info from error message
		http.Error(w, errorSolutionInfo[0]+" field must be of type "+errorSolutionInfo[3], http.StatusBadRequest)
		return
	}

	delay, err := time.ParseDuration(jobRequest.Delay)
	if err != nil {
		http.Error(w, "Invalid delay format", http.StatusBadRequest)
		return
	}

	if jobRequest.Name == "" {
		http.Error(w, "Invalid name length", http.StatusBadRequest)
		return
	}

	job := Job{
		Name:   jobRequest.Name,
		Delay:  delay,
		Number: jobRequest.Value,
	}
	jobQueue <- job
	w.WriteHeader(http.StatusCreated)

}

func main() {
	const (
		maxWorkers   = 4
		maxQueueSize = 20
		port         = ":8081"
	)

	jobQueue := make(chan Job, maxQueueSize)
	dispatcher := NewDispatcher(jobQueue, maxWorkers)

	dispatcher.Run()

	// ​http://127.0.0.1:8081/fib
	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, jobQueue)
	})
	log.Fatal(http.ListenAndServe(port, nil))
}
