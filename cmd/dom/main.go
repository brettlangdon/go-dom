package main

import (
	dom "github.com/brettlangdon/go-dom/v1"
	console "github.com/brettlangdon/go-dom/v1/console"
	document "github.com/brettlangdon/go-dom/v1/document"
)

func onClick(evt *dom.Event) {
	elm := evt.GetTarget().ToElement()
	console.Log(evt, elm)
}

func main() {
	id := "app"
	app := document.GetElementById(id)
	if app == nil {
		console.Error("Could not find element with id %s\r\n", id)
		return
	}

	document.AddEventListener("click", dom.NewEventCallback(onClick))
	stop := make(chan int)
	<-stop
}
