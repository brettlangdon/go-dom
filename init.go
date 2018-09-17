//go:generate go run generate/main.go
package dom

import "syscall/js"

var (
	Document *document
	Window   *window
)

func init() {
	g := js.Global()
	Document = &document{Value: g.Get("document")}
	Window = &window{Value: g.Get("window")}
}
