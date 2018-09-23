// Code generated DO NOT EDIT
// blobcallback.go
package dom

import "syscall/js"

type BlobCallbackCallback func(blob Blob)
type BlobCallback struct {
	Callback
}

func jsValueToBlobCallback(val js.Value) BlobCallback {
	return BlobCallback{Callback: jsValueToCallback(val)}
}
func NewBlobCallback(c BlobCallbackCallback) BlobCallback {
	callback := js.NewCallback(func(args []js.Value) {
		blob := jsValueToBlob(args[0])
		c(blob)
	})
	return BlobCallback{Callback: Callback{Callback: callback}}
}
