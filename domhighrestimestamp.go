// Code generated DO NOT EDIT
// domhighrestimestamp.go
package dom

import "syscall/js"

type DOMHighResTimeStamp float64

func jsValueToDOMHighResTimeStamp(val js.Value) DOMHighResTimeStamp {
	return DOMHighResTimeStamp(val.Float())
}
