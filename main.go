package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
)

type User struct {
	name    string
	session *Session
}

type Session struct {
	conn net.Conn
}

func (s *Session) WriteLine(str string) error {

	_, err := s.conn.Write([]byte(str + "\r\n"))
	return err
}

type World struct {
	users []*User
}

type UserJoinedEvent struct {
}

type MessageEvent struct {
	msg string
}

type ClientInput struct {
	event interface{}
	user  *User
}

func generateName() string {
	return fmt.Sprintf("User %d", rand.Intn(100)+1)
}

func handleConnection(conn net.Conn, inputChannel chan ClientInput) error {
	log.Println("Got a straggler in the MUD")

	session := &Session{conn}
	user := &User{name: generateName(), session: session}
	inputChannel <- ClientInput{

		&UserJoinedEvent{},
		user,
	}
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
		msg := string(buf[:n])
		//log.Println(buf)
		log.Println("received msg", msg)

		e := ClientInput{&MessageEvent{msg}, user}
		inputChannel <- e

	}

	return nil
}

func startServer(eventChannel chan ClientInput) error {
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
			if err := handleConnection(conn, eventChannel); err != nil {
				log.Println("Error handling connection", err)
				return
			}
		}()
	}

}

func startGameLoop(clientInputChannel <-chan ClientInput) {
	w := &World{}
	for input := range clientInputChannel {
		switch event := input.event.(type) {
		case *MessageEvent:
			fmt.Println("received message ", event.msg)
			input.user.session.WriteLine(fmt.Sprintf("You said, \"%s\"\r\n", event.msg))
			for _, user := range w.users {
				if user != input.user {
					user.session.WriteLine(fmt.Sprintf(input.user.name+" said \"%s\"", event.msg))
				}
			}

		case *UserJoinedEvent:
			fmt.Println("USer joined;", input.user.name)
			//input.session.conn.Write("")
			w.users = append(w.users, input.user)

			input.user.session.WriteLine("Welcome, " + input.user.name)
			for _, user := range w.users {
				if user != input.user {
					user.session.WriteLine(fmt.Sprintf(input.user.name + " entered the world"))
				}
			}
			/*n, err = conn.Write([]byte(resp))
			if err != nil {
				log.Println("error reading from buffer", err)
				return err
			}
			if n == 0 {
				log.Println("Zero bytes, closing connection")
				break
			}*/
		}

		fmt.Println("received input", input)
		fmt.Println("received event ", input.event)
	}
}

func main() {

	ch := make(chan ClientInput)

	go startGameLoop(ch)

	err := startServer(ch)
	if err != nil {
		panic(err)
		//log.Fatal(err)
	}

}
