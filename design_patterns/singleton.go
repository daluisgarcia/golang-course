package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct{}

func (Database) GetConnection() {
	fmt.Println("Creating database connection...")
	time.Sleep(2 * time.Second)
	fmt.Println("Database connection created.")
}

var db *Database
var lock sync.Mutex

func GetDatabaseInstance() *Database {
	lock.Lock() // Avoiding race condition in database instanciation
	defer lock.Unlock()
	if db == nil {
		fmt.Println("Creating unique instance of database...")
		db = &Database{}
		db.GetConnection()
	} else {
		fmt.Println("Using existing instance of database")
	}
	return db
}

// func main() {
// 	var wg sync.WaitGroup
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go func() {
// 			GetDatabaseInstance()
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// }
