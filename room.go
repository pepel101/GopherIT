package main

import "fmt"

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
	End          string       `xml:"end"`
}

func (m *Message) Ended(c *Character) bool {
	return c.HasAction(m.End)
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
	OkMessage   string `xml:"ok"`
	FailMessage string `xml:"fail"`
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

func (r *Room) directions() string {
	dirs := ""
	for _, dir := range r.Directions {
		dirs += dir.Direction + ", leading to " + dir.Station + ", "
	}
	return dirs
}

func (r *Room) actions() string {
	acts := ""

	for _, act := range r.Actions {
		acts += act.Name + " ,"
	}
	return acts
}

func (r *Room) getAction(action string) (string, bool) {
	for _, a := range r.Actions {
		if action == a.Name {
			return fmt.Sprintf("%s:%s", r.Key, a.Name), true
		}
	}

	return fmt.Sprintf(""), false
}

func (r *Room) CanDoAction(action Action, character Character) (bool, string) {
	return CheckDependencies(action.Dependencies, character, action.Name)
}

func (r *Room) actionByName(action string) (Action, error) {
	for _, a := range r.Actions {
		if a.Name == action {
			return *a, nil
		}
	}
	return Action{}, nil
}

func CheckDependencies(dependencies []Dependency, character Character, defaultAnswer string) (bool, string) {
	if len(dependencies) == 0 {
		return true, defaultAnswer
	}

	lastOkMessage := ""
	for _, d := range dependencies {
		switch d.Type {
		case "":
			fallthrough
		case "action":
			if !character.HasAction(d.Key) {
				return false, d.FailMessage
			}

			lastOkMessage = d.OkMessage
			/*case "attribute":
			playerAttribute := player.GetAttribute(d.Key)

			minValue, errMin := strconv.ParseInt(d.MinValue, 10, 64)
			if errMin == nil && d.MinValue != "" && playerAttribute < minValue {
				return false, d.FailMessage
			}

			maxValue, errMax := strconv.ParseInt(d.MinValue, 10, 64)
			if errMax == nil && d.MaxValue != "" && playerAttribute > maxValue {
				return false, d.FailMessage
			}

			lastOkMessage = d.OkMessage
			*/
		}
	}
	/*if defaultAnswer != "" {
		lastOkMessage = defaultAnswer
	}*/
	return true, lastOkMessage
}
