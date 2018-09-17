// DO NOT EDIT - generated file
package window

import "syscall/js"
import dom "github.com/brettlangdon/go-dom/v1"

var w *dom.Window

func init() {
	w = dom.NewWindow(js.Global().Get("window"))
}
