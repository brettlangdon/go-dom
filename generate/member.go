package main

import "strings"

type Member struct {
	Body            *Body            `json:"body"`
	Default         *IDLType         `json:"default"`
	Deleter         interface{}      `json:"deleter"`
	EscapedName     string           `json:"escapedName"`
	ExtraAttributes *ExtraAttributes `json:"extAttrs"`
	Getter          interface{}      `json:"getter"`
	IDLType         interface{}      `json:"idlType"`
	Inherit         interface{}      `json:"inherit"`
	Name            string           `json:"name"`
	Readonly        interface{}      `json:"readonly"`
	Required        interface{}      `json:"required"`
	Setter          interface{}      `json:"setter"`
	Static          interface{}      `json:"static"`
	Stringifier     interface{}      `json:"stringifier"`
	Trivia          interface{}      `json:"trivia"`
	Type            string           `json:"type"`
	Value           Value            `json:"value"`
}

func (m Member) Title() string {
	title := strings.Title(m.Name)
	if m.Type == "operation" && m.Body != nil {
		t := m.Body.Title()
		if t != "" {
			title = t
		}
	}
	return title
}

func (m Member) IsReadOnly() bool { return m.Readonly != nil }
