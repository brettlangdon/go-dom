package main

type Argument struct {
	Default         interface{} `json:"default"`
	EscapedName     string      `json:"escapedName"`
	ExtraAttributes interface{} `json:"extAttrs"`
	IDLType         IDLType     `json:"idlType"`
	Name            string      `json:"name"`
	Optional        interface{} `json:"optional"`
	Separator       interface{} `json:"separator"`
	Trivia          interface{} `json:"trivia"`
	Variadic        interface{} `json:"variadic"`
}
