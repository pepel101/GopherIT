package main

import (
	"encoding/xml"
	"fmt"
	"math/rand"
	"strings"
)

type SessionEvent struct {
	Session *Session
	Event   interface{}
}

type SessionCreatedEvent struct{}

type SessionDisconnectEvent struct{}

type SessionInputEvent struct {
	input string
}

type Entity struct {
	entityId string
}

func (e *Entity) EntityId() string {
	return e.entityId
}

type User struct {
	Session   *Session
	Character *Character
}

type Character struct {
	Name       string
	User       *User
	Room       *Room
	XMLName    xml.Name    `xml:"player"`
	Nickname   string      `xml:"nickname,attr"`
	Gamename   string      `xml:"name"`
	Position   string      `xml:"position,attr"`
	PlayerType string      `xml:"type"`
	Ch         chan string `xml:"-"`
	ActionLog  []string    `xml:"actions>action"`
	//Attributes []Attribute `xml:"attributes>attribute"`
}

func (c *Character) SendMessage(msg string) {
	c.User.Session.WriteLine(msg)
}

func (c *Character) LogAction(action string) {
	if !c.HasAction(action) {
		c.ActionLog = append(c.ActionLog, strings.ToLower(action))
	}
}

func (c *Character) HasAction(action string) bool {
	for _, a := range c.ActionLog {
		if strings.ToLower(a) == strings.ToLower(action) {
			return true
		}
	}
	return false
}

func generateName() string {
	return fmt.Sprintf("User %d", rand.Intn(100)+1)
}

type MessageEvent struct {
	msg string
}

type MoveEvent struct {
	dir string
}

type UserJoinedEvent struct {
}
