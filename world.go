package main

import "fmt"

type World struct {
	characters []*Character
	rooms      []*Room
}

func NewWorld() *World {
	return &World{}
}

func (w *World) Init() {
	w.rooms = []*Room{
		{
			Id:   "A",
			Desc: "This is a room with a sign that has the letter A written on it.",
			Links: []*RoomLink{
				{
					Verb:   "east",
					RoomId: "B",
				},
			},
		},
		{
			Id:   "B",
			Desc: "This is a room with a sign that has the letter B written on it.",
			Links: []*RoomLink{
				{
					Verb:   "west",
					RoomId: "A",
				},
			},
		},
	}
}

func (w *World) HandleCharacterJoined(character *Character) {
	w.rooms[0].AddCharacter(character)

	character.SendMessage("Welcome!")
	character.SendMessage("")
	character.SendMessage(character.Room.Desc)
	for _, other := range character.Room.Characters {
		if other != character {
			other.SendMessage(fmt.Sprintf(character.Name + " joined the room"))
		}
	}

}

func (w *World) GetRoomById(id string) *Room {
	for _, r := range w.rooms {
		if r.Id == id {
			return r
		}
	}
	return nil
}

func (w *World) HandleCharacterInput(character *Character, input string) {
	room := character.Room
	for _, link := range room.Links {
		if link.Verb == input {
			target := w.GetRoomById(link.RoomId)
			if target != nil {
				w.MoveCharacter(character, target)
				return
			}
		}
	}

	character.SendMessage(fmt.Sprintf("You said " + input))

	for _, other := range character.Room.Characters {
		if other != character {
			other.SendMessage(fmt.Sprintf(character.Name + " said, " + input))
		}
	}
}

func (world *World) MoveCharacter(character *Character, to *Room) {
	character.Room.RemoveCharacter(character)
	to.AddCharacter(character)
	character.SendMessage(to.Desc)
}
