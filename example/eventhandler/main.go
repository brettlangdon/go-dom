package main

import (
	dom "github.com/brettlangdon/go-dom"
	"github.com/brettlangdon/go-dom/console"
	"github.com/brettlangdon/go-dom/document"
)

func main() {
	loop := dom.NewLoop()

	document.AddEventListener("click", dom.NewEventHandler(func(evt dom.Event) {
		console.Dir(evt)
	}))

	loop.Loop()
}
