// Code generated DO NOT EDIT
// binarytype.go
package dom

import "syscall/js"

type BinaryType string

const (
	BinaryTypeBlob        BinaryType = "blob"
	BinaryTypeArraybuffer BinaryType = "arraybuffer"
)

func jsValueToBinaryType(val js.Value) BinaryType {
	return BinaryType(val.String())
}
