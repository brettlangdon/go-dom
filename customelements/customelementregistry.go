// DO NOT EDIT - generated file
package customelements

import "syscall/js"
import dom "github.com/brettlangdon/go-dom/v1"

var c *dom.CustomElementRegistry

func init() {
	c = dom.NewCustomElementRegistry(js.Global().Get("customElements"))
}
func Define(name string, constructor interface{}) dom.Value {
	return c.Define(name, constructor)
}
func Get(name string) dom.Value {
	return c.Get(name)
}
func WhenDefined(name string) *dom.Promise {
	return c.WhenDefined(name)
}
