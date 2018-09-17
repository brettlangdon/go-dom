package dom

import "syscall/js"

type JSValue interface {
	JSValue() js.Value
}

func ToJSValue(x interface{}) js.Value {
	if v, ok := x.(JSValue); ok {
		return v.JSValue()
	}
	return js.ValueOf(x)
}

type Value struct {
	js.Value
}

func (v Value) JSValue() js.Value { return v.Value }

func (v Value) IsUndefined() bool { return v.Type() == js.TypeUndefined }
func (v Value) IsNull() bool      { return v.Type() == js.TypeNull }
func (v Value) IsBoolean() bool   { return v.Type() == js.TypeBoolean }
func (v Value) IsNumber() bool    { return v.Type() == js.TypeNumber }
func (v Value) IsString() bool    { return v.Type() == js.TypeString }
func (v Value) IsSymbol() bool    { return v.Type() == js.TypeSymbol }
func (v Value) IsObject() bool    { return v.Type() == js.TypeObject }
func (v Value) IsFunction() bool  { return v.Type() == js.TypeFunction }
