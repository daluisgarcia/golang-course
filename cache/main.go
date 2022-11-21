package main

import (
	"fmt"
	"sync"
	"time"
)

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

type Memory struct {
	f     Function               // Function to cache
	cache map[int]FunctionResult // Cache of the function results
	lock  sync.Mutex             // Mutex to protect the cache
}

func NewCache(f Function) *Memory { // Memory constructor
	return &Memory{
		f:     f,
		cache: make(map[int]FunctionResult),
	}
}

func (mem *Memory) Get(key int) (interface{}, error) {
	mem.lock.Lock() // Locks the access to the cache by others threads
	result, exists := mem.cache[key]
	mem.lock.Unlock()

	start := time.Now() // Starts the timer to calculate the time of the value calculation

	if !exists {
		mem.lock.Lock() // Locks the access to the cache by others threads
		result.value, result.err = mem.f(key)
		mem.cache[key] = result
		mem.lock.Unlock()
	}

	fmt.Printf("%d, %s, %d\n", key, time.Since(start), result.value) // Prints the time of the value calculation

	return result.value, result.err
}

type Function func(int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err   error
}

func GetFibonacci(n int) (interface{}, error) {
	return Fibonacci(n), nil
}

// func main() {
// 	cache := NewCache(GetFibonacci)
// 	fibo := []int{42, 40, 41, 42, 38}

// 	var wg sync.WaitGroup

// 	for _, n := range fibo {
// 		wg.Add(1)
// 		go func(num int) {
// 			defer wg.Done()

// 			_, err := cache.Get(num)

// 			if err != nil {
// 				log.Println(err)
// 			}

// 		}(n)
// 	}
// 	wg.Wait()
// }
