package main

import (
	"log"
)

func main() {
	ch := make(chan SessionEvent)
	world := NewWorld()
	world.Init()

	world.roomsInit()
	sessionHandler := NewSessionHandler(world, ch)
	go sessionHandler.Start()

	err := startServer(ch)
	if err != nil {
		log.Fatal(err)
	}
}
