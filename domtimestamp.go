// Code generated DO NOT EDIT
// domtimestamp.go
package dom

import "syscall/js"

type DOMTimeStamp int

func JSValueToDOMTimeStamp(val js.Value) DOMTimeStamp {
	return DOMTimeStamp(val.Int())
}
