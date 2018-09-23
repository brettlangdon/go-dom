// Code generated DO NOT EDIT
// binarytype.go
package dom

import "syscall/js"

type BinaryType string

const (
	BinaryTypeBlob        BinaryType = "blob"
	BinaryTypeArraybuffer BinaryType = "arraybuffer"
)

func JSValueToBinaryType(val js.Value) BinaryType {
	return BinaryType(val.String())
}
