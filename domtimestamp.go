// Code generated DO NOT EDIT
// domtimestamp.go
package dom

import "syscall/js"

type DOMTimeStamp float64

func jsValueToDOMTimeStamp(val js.Value) DOMTimeStamp {
	return DOMTimeStamp(val.Float())
}
