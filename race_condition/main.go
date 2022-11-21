package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	balance = 100
)

// To avoid race condition, we need to use mutex
func Deposit(amount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()
	lock.Lock() // Blocks the access to shared resource.
	// The code inside the lock will be executed by one thread at a time
	b := balance
	time.Sleep(time.Duration(amount) * time.Millisecond)
	balance = b + amount
	lock.Unlock()
}

func Balance(lock *sync.RWMutex) int {
	lock.RLock() // Especial lock that allos to read even when there is a write lock
	b := balance
	lock.RUnlock()
	return b
}

// To see if there is any possibility os race condition in the code, run the command: go build --race <.go file> and run the executable.
//                                                                                                        This will show warnings about.

func main() {
	var wg sync.WaitGroup
	var lock sync.RWMutex
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &lock)
		fmt.Println(Balance(&lock))
	}
	wg.Wait()
}
