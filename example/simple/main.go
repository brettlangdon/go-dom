package main

import (
	"github.com/brettlangdon/go-dom/console"
	"github.com/brettlangdon/go-dom/document"
)

func main() {
	nodes := document.QuerySelectorAll("div")
	for i := 0; i < nodes.GetLength(); i++ {
		console.Dir(nodes.Item(0))
	}
}
