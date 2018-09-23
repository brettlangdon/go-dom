// Code generated DO NOT EDIT
// documentreadystate.go
package dom

import "syscall/js"

type DocumentReadyState string

const (
	DocumentReadyStateLoading     DocumentReadyState = "loading"
	DocumentReadyStateInteractive DocumentReadyState = "interactive"
	DocumentReadyStateComplete    DocumentReadyState = "complete"
)

func jsValueToDocumentReadyState(val js.Value) DocumentReadyState {
	return DocumentReadyState(val.String())
}
