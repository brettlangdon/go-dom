package dom

import "syscall/js"

func ToValue(v interface{}) js.Value {
	switch t := v.(type) {
	case *Element:
		return t.Value
	case *document:
		return t.Value
	default:
		return js.ValueOf(v)
	}
}
