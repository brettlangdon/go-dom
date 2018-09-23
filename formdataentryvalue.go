// Code generated DO NOT EDIT
// formdataentryvalue.go
package dom

import "syscall/js"

type FormDataEntryValue Value

func JSValueToFormDataEntryValue(val js.Value) FormDataEntryValue {
	return FormDataEntryValue(JSValueToValue(val))
}
