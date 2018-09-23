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

func jsValueToSelectionMode(val js.Value) SelectionMode {
	return SelectionMode(val.String())
}
