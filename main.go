package main

import (
	"log"
	"net"
)

func handleConnection(conn net.Conn) error {
	log.Println("Got a straggler in the MUD")

	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("error reading from buffer", err)
			return err
		}
		if n == 0 {
			log.Println("Zero bytes, closing connection")
			break
		}
		msg := string(buf)
		log.Println("received msg", msg)
	}

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
