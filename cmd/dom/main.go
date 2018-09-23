package main

import (
	"github.com/brettlangdon/go-dom/v1/console"
	"github.com/brettlangdon/go-dom/v1/document"
	"github.com/brettlangdon/go-dom/v1/window"
)

func main() {
	loc := window.GetLocation()
	console.Dir(loc.JSValue())

	nodes := document.QuerySelectorAll("div")
	var i float64 = 0
	for ; i < nodes.GetLength(); i++ {
		console.Dir(nodes.Item(i).JSValue())
	}
}
