package main

import (
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatalf("SERVER_PORT variable is not correct format:%v", err)
	}
	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: port})
	if err != nil {
		log.Fatalf("Udp Service listen report udp fail:%v", err)
	}
	log.Printf("Running UDP server on port %d\n", port)
	defer conn.Close()
	var data = make([]byte, 1024*4)
	var raw []byte
	for {
		n, _, err := conn.ReadFromUDP(data)
		if err == nil {
			raw = make([]byte, n)
			copy(raw, data[:n])
			go func() {
				// ... use raw to do something
				log.Printf("%s", raw)
			}()
		}
	}
}
