package dom

import "syscall/js"

type JSValue interface {
	JSValue() js.Value
}

type Value struct {
	js.Value
}

func JSValueToValue(val js.Value) Value { return Value{Value: val} }

func (v Value) JSValue() js.Value  { return v.Value }
func (v Value) Get(p string) Value { return JSValueToValue(v.Value.Get(p)) }
func (v Value) Index(i int) Value  { return JSValueToValue(v.Value.Index(i)) }
func (v Value) Call(m string, args ...interface{}) Value {
	for i, a := range args {
		args[i] = GetJSValue(a)
	}
	return JSValueToValue(v.Value.Call(m, args...))
}
func (v Value) Invoke(args ...interface{}) Value {
	for i, a := range args {
		args[i] = GetJSValue(a)
	}
	return JSValueToValue(v.Value.Invoke(args...))
}
func (v Value) New(args ...interface{}) Value {
	for i, a := range args {
		args[i] = GetJSValue(a)
	}
	return JSValueToValue(v.Value.New(args...))
}

func GetJSValue(v interface{}) interface{} {
	t, ok := v.(JSValue)
	if ok {
		return t.JSValue()
	}
	return v
}

func (v Value) IsUndefined() bool       { return v.Type() == js.TypeUndefined }
func (v Value) IsNull() bool            { return v.Type() == js.TypeNull }
func (v Value) IsNullOrUndefined() bool { return v.IsNull() || v.IsUndefined() }
func (v Value) IsBoolean() bool         { return v.Type() == js.TypeBoolean }
func (v Value) IsNumber() bool          { return v.Type() == js.TypeNumber }
func (v Value) IsString() bool          { return v.Type() == js.TypeString }
func (v Value) IsSymbol() bool          { return v.Type() == js.TypeSymbol }
func (v Value) IsObject() bool          { return v.Type() == js.TypeObject }
func (v Value) IsFunction() bool        { return v.Type() == js.TypeFunction }
