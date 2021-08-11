package main

type Room struct {
	Id         string
	Desc       string
	Links      []*RoomLink
	Key        string       `xml:"key,attr"`
	Tag        string       `xml:"tag,attr"`
	Name       string       `xml:"name"`
	Directions []*Direction `xml:"directions>direction"`
	Actions    []*Action    `xml:"actions>action"`
	Messages   []*Message   `xml:"messages>message"`
	Intro      string       `xml:"desc"`
	//Asciimation Asciimation `xml:"asciimation"`

	Characters []*Character
}

type Message struct {
	Text         string       `xml:"text"`
	Dependencies []Dependency `xml:"dep"`
}

type Action struct {
	Name         string       `xml:"name,attr"`
	Hidden       string       `xml:"hidden,attr"`
	Dependencies []Dependency `xml:"dep"`
	Answer       string       `xml:"ok"`
}

type Direction struct {
	Station      string       `xml:"room"`
	Hidden       bool         `xml:"hidden,attr"`
	Dependencies []Dependency `xml:"dep"`
	Direction    string       `xml:"name"`
	DirKey       string       `xml:"key"`
}

type Dependency struct {
	Key         string `xml:"key,attr"`
	Type        string `xml:"type,attr"`
	MinValue    string `xml:"minValue"`
	MaxValue    string `xml:"maxValue"`
	OkMessage   string `xml:"okMessage"`
	FailMessage string `xml:"failMessage"`
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
	Link := &RoomLink{}
	Link.RoomId = RoomId
	Link.Verb = Verb
	r.Links = append(r.Links, Link)
}

func (r *Room) directions() []*Direction {
	return r.Directions
}

func (r *Room) actions() []*Action {
	return r.Actions
}
