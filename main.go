package main

import (
	"log"
	"net"
)

func main() {

	s := newServer()
	go s.run()

	listener, err := net.Listen("tcp", ":8888")

	if err != nil {
		log.Fatalf("unable to start server %s", err.Error())
	}

	defer listener.Close()
	log.Printf("sever start on the port :8888")

	for {

		conn, err := listener.Accept()

		if err != nil {
			log.Printf("failed to accept the connection: %s", err.Error())
			continue
		}

		c := s.newClient(conn)
		go c.readInput()
	}
}
