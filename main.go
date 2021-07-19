package main

import (
	"log"
	"net"
)

func handleConnection(conn net.Conn) error {
	log.Println("Got a straggler in the MUD")
	return nil
}

func startServer() error {
	log.Print("starting server")

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("error accepting connection")
		}
		go func() {
			if err := handleConnection(conn); err != nil {
				log.Println("Error handling connection", err)
				return
			}
		}()
	}

}

func main() {

	err := startServer()
	if err != nil {
		log.Fatal(err)
	}
}
