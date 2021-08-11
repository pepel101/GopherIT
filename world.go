package main

import (
	"fmt"
	"strings"
)

type World struct {
	characters []*Character
	rooms      []*Room
}

func NewWorld() *World {
	return &World{}
}

func (w *World) Init() {
	w.rooms = []*Room{
		/*{
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
		},*/
	}
}

func (w *World) addRoom(room *Room) {
	w.rooms = append(w.rooms, room)
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

func (w *World) HandleCharacterMoved(character *Character) {
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
	//room := character.Room
	/*for _, link := range room.Links {
		if runtime.GOOS == "windows" {
			input = strings.TrimRight(input, "\r\n")
		} else {
			input = strings.TrimRight(input, "\n")
		}
		input = strings.TrimSpace(input)
		if link.Verb == input {
			eq := true
			fmt.Println("aaaa", eq)
			target := w.GetRoomById(link.RoomId)
			if target != nil {
				w.MoveCharacter(character, target)
				w.HandleCharacterMoved(character)
				return
			}
		} else {
			eq := false
			fmt.Println("aaa", eq)
		}

	}
	*/

	inputParts := strings.SplitN(input, " ", 2)

	var command, commandText string
	if len(inputParts) > 1 {
		command = inputParts[0]
		commandText = inputParts[1]

		switch command {
		case "say":
			fallthrough
		case "speak":
			character.SendMessage(fmt.Sprintf("You said " + commandText))

			for _, other := range character.Room.Characters {
				if other != character {
					other.SendMessage(fmt.Sprintf(character.Name + " said " + commandText))
				}
			}
		case "go":
			fallthrough
		case "exit":
			var toRoom *Room
			for _, dir := range character.Room.Directions {
				if strings.ToLower(strings.TrimSpace(commandText)) == dir.Direction {
					toRoom = w.getRoomByKey(dir.DirKey)
				}
			}
			w.MoveCharacter(character, toRoom)
		}

	} else {
		character.SendMessage(fmt.Sprintf("You said " + input))

		for _, other := range character.Room.Characters {
			if other != character {
				other.SendMessage(fmt.Sprintf(character.Name + " said " + input))
			}
		}
	}
}

func (world *World) MoveCharacter(character *Character, to *Room) {
	character.Room.RemoveCharacter(character)
	to.AddCharacter(character)
	character.SendMessage(to.Intro)
}

func (w *World) getRoomByKey(key string) *Room {
	for _, room := range w.rooms {
		if room.Key == key {
			return room
		}
	}
	return nil
}
