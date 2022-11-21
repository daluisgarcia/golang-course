package main

import (
	"flag"
	"io"
	"log"
)

var (
	port = flag.Int("port", 3090, "The port to scan")
	host = flag.String("host", "localhost", "The host to scan")
)

func CopyContent(dct io.Writer, src io.Reader) {
	_, err := io.Copy(dct, src)

	if err != nil {
		log.Fatal(err)
	}
}

// func main() { // Client main function for the chat - DECOMMENT TO RUN
// 	flag.Parse()
// 	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	done := make(chan struct{}) // Control channel

// 	go func() {
// 		io.Copy(os.Stdout, conn)
// 		done <- struct{}{}
// 	}()

// 	CopyContent(conn, os.Stdin)
// 	conn.Close()
// 	<-done
// }
