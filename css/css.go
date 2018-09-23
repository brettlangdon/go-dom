// Code generated DO NOT EDIT
// css/css.go
package css

import dom "github.com/brettlangdon/go-dom/v1"
import "syscall/js"

var value dom.Value

func init() { value = dom.JSValueToValue(js.Global().Get("css")) }
func Escape(args ...interface{}) string {
	val := value.Call("escape", args...)
	return val.String()
}
