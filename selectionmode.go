// Code generated DO NOT EDIT
// selectionmode.go
package dom

import "syscall/js"

type SelectionMode string

const (
	SelectionModeSelect   SelectionMode = "select"
	SelectionModeStart    SelectionMode = "start"
	SelectionModeEnd      SelectionMode = "end"
	SelectionModePreserve SelectionMode = "preserve"
)

func JSValueToSelectionMode(val js.Value) SelectionMode {
	return SelectionMode(val.String())
}
