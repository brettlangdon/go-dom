// Code generated DO NOT EDIT
// css/css.go
package css

import "syscall/js"

var value Value

func init()                      { value = Value{Value: js.Global().Get("CSS")} }
func Escape(args ...interface{}) { return value.Call("escape", args...) }
