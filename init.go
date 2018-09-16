//go:generate go run generate/main.go
package dom

import "syscall/js"

var (
	Document *document
)

func init() {
	Document = &document{Value: js.Global().Get("document")}
}
