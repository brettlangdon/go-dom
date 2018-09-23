package main

import (
	"syscall/js"

	"github.com/brettlangdon/go-dom/v1/document"
)

func main() {
	console := js.Global().Get("console")
	nodes := document.Document.QuerySelectorAll("div")
	var i float64 = 0
	for ; i < nodes.GetLength(); i++ {
		console.Call("dir", nodes.Item(i).JSValue())
	}
}
