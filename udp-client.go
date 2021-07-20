package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	connection_string := os.Getenv("CLIENT_CONNECTION_STRING")
	if connection_string == "" {
		log.Fatalf("CLIENT_CONNECTION_STRING can not be empty. Set e.g. export CLIENT_CONNECTION_STRING='localhost:7777'.")
	}
	conn, err := net.DialTimeout("udp", connection_string, 30*time.Second)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to %s", connection_string))
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example ", f)

	i := 1
	for {
		logger.Printf("This is a message %d.", i)
		i++
		time.Sleep(5 * time.Second)
	}
}
