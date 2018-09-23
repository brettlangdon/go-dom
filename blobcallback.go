// Code generated DO NOT EDIT
// blobcallback.go
package dom

import "syscall/js"

type BlobCallbackCallback func(blob Blob)
type BlobCallback struct {
	Callback
}

func JSValueToBlobCallback(val js.Value) BlobCallback {
	return BlobCallback{Callback: JSValueToCallback(val)}
}
func NewBlobCallback(c BlobCallbackCallback) BlobCallback {
	callback := js.NewCallback(func(args []js.Value) {
		blob := JSValueToBlob(args[0])
		c(blob)
	})
	return BlobCallback{Callback: Callback{Callback: callback}}
}
