// Code generated DO NOT EDIT
// validitystate.go
package dom

import "syscall/js"

type ValidityStateIFace interface {
	GetBadInput() bool
	GetCustomError() bool
	GetPatternMismatch() bool
	GetRangeOverflow() bool
	GetRangeUnderflow() bool
	GetStepMismatch() bool
	GetTooLong() bool
	GetTooShort() bool
	GetTypeMismatch() bool
	GetValid() bool
	GetValueMissing() bool
}
type ValidityState struct {
	Value
}

func JSValueToValidityState(val js.Value) ValidityState {
	return ValidityState{Value: Value{Value: val}}
}
func (v Value) AsValidityState() ValidityState { return ValidityState{Value: v} }
func (v ValidityState) GetBadInput() bool {
	val := v.Get("badInput")
	return val.Bool()
}
func (v ValidityState) GetCustomError() bool {
	val := v.Get("customError")
	return val.Bool()
}
func (v ValidityState) GetPatternMismatch() bool {
	val := v.Get("patternMismatch")
	return val.Bool()
}
func (v ValidityState) GetRangeOverflow() bool {
	val := v.Get("rangeOverflow")
	return val.Bool()
}
func (v ValidityState) GetRangeUnderflow() bool {
	val := v.Get("rangeUnderflow")
	return val.Bool()
}
func (v ValidityState) GetStepMismatch() bool {
	val := v.Get("stepMismatch")
	return val.Bool()
}
func (v ValidityState) GetTooLong() bool {
	val := v.Get("tooLong")
	return val.Bool()
}
func (v ValidityState) GetTooShort() bool {
	val := v.Get("tooShort")
	return val.Bool()
}
func (v ValidityState) GetTypeMismatch() bool {
	val := v.Get("typeMismatch")
	return val.Bool()
}
func (v ValidityState) GetValid() bool {
	val := v.Get("valid")
	return val.Bool()
}
func (v ValidityState) GetValueMissing() bool {
	val := v.Get("valueMissing")
	return val.Bool()
}
