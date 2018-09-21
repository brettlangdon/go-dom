package main

import (
	dom "github.com/brettlangdon/go-dom/v1"
	"github.com/brettlangdon/go-dom/v1/console"
	"github.com/brettlangdon/go-dom/v1/document"
)

func main() {
	app := document.GetElementById("app")
	shadow := app.AttachShadow(dom.ShadowRootInit{Mode: "open"})
	console.Dir(shadow)
}
