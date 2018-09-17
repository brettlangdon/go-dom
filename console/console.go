// DO NOT EDIT - generated file
package console

import "syscall/js"
import dom "github.com/brettlangdon/go-dom/v1"

var c *dom.Console

func init() {
	c = dom.NewConsole(js.Global().Get("console"))
}
func Log(v ...interface{}) dom.Value {
	return c.Log(v...)
}
func Error(v ...interface{}) dom.Value {
	return c.Error(v...)
}
func Dir(v dom.JSValue) dom.Value {
	return c.Dir(v)
}
