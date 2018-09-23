package main

type Value struct {
	Separator interface{} `json:"separator"`
	Trivia    interface{} `json:"trivia"`
	Type      string      `json:"type"`
	Value     string      `json:"value"`
}
