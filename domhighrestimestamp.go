// Code generated DO NOT EDIT
// domhighrestimestamp.go
package dom

import "syscall/js"

type DOMHighResTimeStamp float64

func JSValueToDOMHighResTimeStamp(val js.Value) DOMHighResTimeStamp {
	return DOMHighResTimeStamp(val.Float())
}
