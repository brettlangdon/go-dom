package main

import (
	"fmt"
	"syscall/js"

	dom "github.com/brettlangdon/go-dom/v1"
)

func onLoad(evt js.Value) {
	fmt.Println("LOADED")
	fmt.Printf("%#v\r\n", evt)
}

func main() {
	dom.Document.AddEventListener("DOMContentLoaded", js.NewEventCallback(0, onLoad))
	fmt.Printf("%#v\r\n", dom.Document)
	body := dom.Document.GetBody()
	for i := 0; i < 50; i += 1 {
		body.AppendChild(dom.Document.CreateElement(fmt.Sprintf("TEST-%d", i)))
	}
}
