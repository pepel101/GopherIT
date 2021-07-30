package main

type Room struct {
	Id    string
	Desc  string
	Links []*RoomLink

	Characters []*Character
}
type RoomLink struct {
	Verb   string
	RoomId string
}

func (r *Room) AddCharacter(character *Character) {
	r.Characters = append(r.Characters, character)
	character.Room = r
}

func (r *Room) RemoveCharacter(character *Character) {
	character.Room = nil

	var characters []*Character
	for _, c := range r.Characters {
		if c != character {
			characters = append(characters, c)
		}
	}
	r.Characters = characters
}

func (r *Room) roomInit(Id string, Desc string) {
	r.Id = Id
	r.Desc = Desc
}

func (r *Room) roomInitId(Id string) {
	r.Id = Id
}

func (r *Room) roomInitDesc(Desc string) {
	r.Desc = Desc
}

func (r *Room) addLink(Verb string, RoomId string) {
	var Link *RoomLink
	Link.RoomId = RoomId
	Link.Verb = Verb
	r.Links = append(r.Links, Link)
}
