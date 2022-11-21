package main

import (
	"fmt"
	"sync"
	"time"
)

func ExpensiveFibonacci(n int) int {
	fmt.Printf("Calculating the expensive fibonacci of %d\n", n)
	time.Sleep(5 * time.Second)
	return n
}

type Service struct {
	InProgress map[int]bool       // Map of jobs that are in progress
	IsPending  map[int][]chan int // Queue of jobs that are waiting for the calculation to be done
	Lock       sync.RWMutex       // Lock for the map
}

func (s *Service) Work(job int) {
	s.Lock.RLock()

	calculationIsInProgress := s.InProgress[job]

	if calculationIsInProgress { // Add listener to the pending list
		s.Lock.RUnlock()
		response := make(chan int)
		defer close(response)

		s.Lock.Lock()
		s.IsPending[job] = append(s.IsPending[job], response) // Appennd the thread to wait the calculation to be done
		s.Lock.Unlock()

		fmt.Printf("Waiting for job %d to finish\n", job)

		resp := <-response
		fmt.Printf("Job %d finished with result %d\n", job, resp)
		return
	}

	s.Lock.RUnlock()

	s.Lock.Lock()
	s.InProgress[job] = true
	s.Lock.Unlock()

	fmt.Printf("Calculating fibonacci for job %d\n", job)
	result := ExpensiveFibonacci(job)

	s.Lock.Lock()
	pendingWorkers, exists := s.IsPending[job]
	s.Lock.Unlock()

	if exists {
		for _, pendingWorker := range pendingWorkers {
			pendingWorker <- result // Notifies the threads that the calculation is done
		}
		fmt.Printf("Results sent - pending workers ready: %d\n", result)
	}

	s.Lock.Lock()
	s.InProgress[job] = false
	s.IsPending[job] = make([]chan int, 0)
	s.Lock.Unlock()
}

func NewService() *Service {
	return &Service{
		InProgress: make(map[int]bool),
		IsPending:  make(map[int][]chan int),
	}
}

func main() {
	service := NewService()
	jobs := []int{3, 4, 5, 5, 4, 8, 8, 8}

	var wg sync.WaitGroup
	wg.Add(len(jobs))
	for _, job := range jobs {
		go func(num int) {
			defer wg.Done()
			service.Work(num)
		}(job)
	}
	wg.Wait()

}
