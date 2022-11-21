package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var site = flag.String("website", "scanme.nmap.org", "The website to scan")

func program() {
	flag.Parse() // Allows to user the flag defined above
	// A simple port scanner with concurrency
	var wg sync.WaitGroup

	for port := 0; port < 65535; port++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, p))

			if err != nil {
				return
			}

			conn.Close()
			fmt.Printf("Port %d is open\n", p)
		}(port)
	}

	wg.Wait()
}

// Run the program with the flag --website=url.test
// func main() {
// 	program()
// }
