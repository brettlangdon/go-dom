package document

import (
	"syscall/js"

	dom "github.com/brettlangdon/go-dom/v1"
)

var value dom.Document

func init() {
	value = dom.Value{Value: js.Global().Get("document")}.AsDocument()
}
