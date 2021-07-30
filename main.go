package main

import (
	"log"
)

func main() {
	ch := make(chan SessionEvent)
	world := NewWorld()
	world.Init()

	rooms, err1 := roomsInit()
	if err1 != nil {
		log.Fatal(err1)
	}
	world.rooms = append(world.rooms, rooms...)
	sessionHandler := NewSessionHandler(world, ch)
	go sessionHandler.Start()

	err := startServer(ch)
	if err != nil {
		log.Fatal(err)
	}
}
