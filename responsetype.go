// Code generated DO NOT EDIT
// responsetype.go
package dom

import "syscall/js"

type ResponseType string

const (
	ResponseTypeBasic          ResponseType = "basic"
	ResponseTypeCors           ResponseType = "cors"
	ResponseTypeDefault        ResponseType = "default"
	ResponseTypeError          ResponseType = "error"
	ResponseTypeOpaque         ResponseType = "opaque"
	ResponseTypeOpaqueredirect ResponseType = "opaqueredirect"
)

func jsValueToResponseType(val js.Value) ResponseType {
	return ResponseType(val.String())
}
