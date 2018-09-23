package main

import "strings"

type Body struct {
	Arguments []Argument  `json:"arguments"`
	IDLType   IDLType     `json:"idlType"`
	Name      Name        `json:"Name"`
	Trivia    interface{} `json:"trivia"`
}

func (b Body) Title() string { return strings.Title(b.Name.Value) }
