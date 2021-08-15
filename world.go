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
	w.rooms = []*Room{}
}

func (w *World) addRoom(room *Room) {
	w.rooms = append(w.rooms, room)
}

func (w *World) HandleCharacterJoined(character *Character) {
	w.rooms[0].AddCharacter(character)

	character.SendMessage("Welcome!")
	character.SendMessage("")
	character.SendMessage(character.Room.Intro)
	character.LogAction("316:start")
	character.SendMessage(fmt.Sprintf("You see the following exits: %s", character.Room.directions()))
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
	} else {
		command = strings.TrimSpace(inputParts[0])
	}

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
		val := false
		for _, dir := range character.Room.Directions {
			if strings.ToLower(strings.TrimSpace(commandText)) == dir.Direction {
				toRoom = w.getRoomByKey(dir.DirKey)
				w.MoveCharacter(character, toRoom)
				val = true
			}
		}
		if !val {
			character.SendMessage("please specify the direction you want to go")
		}
	default:
		action, gotAction := character.Room.getAction(command)
		if gotAction {
			aAction, err := character.Room.actionByName(command)
			if err != nil {
				character.SendMessage("No such action")
			}
			allowed, msg := character.Room.CanDoAction(aAction, *character)
			if msg != "" {
				character.SendMessage(msg)
			}
			if allowed {
				character.LogAction(action)

			}

		} else {
			character.SendMessage("No such action")
		}
	}

	/* else {
		character.SendMessage(fmt.Sprintf("You said " + input))

		for _, other := range character.Room.Characters {
			if other != character {
				other.SendMessage(fmt.Sprintf(character.Name + " said " + input))
			}
		}
	}*/
}

func (world *World) MoveCharacter(character *Character, to *Room) {
	character.Room.RemoveCharacter(character)
	to.AddCharacter(character)
	character.SendMessage(to.Intro)
	character.SendMessage(fmt.Sprintf("You see the following exits: %s", character.Room.directions()))
	if len(to.Messages) > 0 {
		for _, m := range to.Messages {
			if len(m.Dependencies) > 0 {
				ok, _ := CheckDependencies(m.Dependencies, *character, "No messages")
				if ok {
					if !m.Ended(character) {
						character.SendMessage(m.Text)
					}

				}
			} else {
				if !m.Ended(character) {
					character.SendMessage(m.Text)
				}
			}
		}
	}

}

func (w *World) getRoomByKey(key string) *Room {
	for _, room := range w.rooms {
		if room.Key == key {
			return room
		}
	}
	return nil
}
