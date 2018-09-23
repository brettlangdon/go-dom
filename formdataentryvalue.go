// Code generated DO NOT EDIT
// formdataentryvalue.go
package dom

import "syscall/js"

type FormDataEntryValue Value

func jsValueToFormDataEntryValue(val js.Value) FormDataEntryValue {
	return FormDataEntryValue(jsValueToValue(val))
}
