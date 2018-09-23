package main

type Name struct {
	Value   string      `json:"value"`
	Escaped string      `json:"escaped"`
	Trivia  interface{} `json:"trivia"`
}
