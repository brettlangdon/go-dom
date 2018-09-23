package main

import (
	"github.com/brettlangdon/go-dom/console"
	"github.com/brettlangdon/go-dom/document"
	"github.com/brettlangdon/go-dom/window"
)

func main() {
	loc := window.GetLocation()
	console.Dir(loc)

	nodes := document.QuerySelectorAll("div")
	for i := 0; i < nodes.GetLength(); i++ {
		console.Dir(nodes.Item(i))
	}
}
