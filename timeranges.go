// Code generated DO NOT EDIT
// timeranges.go
package dom

import "syscall/js"

type TimeRangesIFace interface {
	End(args ...interface{}) float64
	GetLength() float64
	Start(args ...interface{}) float64
}
type TimeRanges struct {
	Value
}

func jsValueToTimeRanges(val js.Value) TimeRanges { return TimeRanges{Value: Value{Value: val}} }
func (v Value) AsTimeRanges() TimeRanges          { return TimeRanges{Value: v} }
func (t TimeRanges) End(args ...interface{}) float64 {
	val := t.Call("end", args...)
	return val.Float()
}
func (t TimeRanges) GetLength() float64 {
	val := t.Get("length")
	return val.Float()
}
func (t TimeRanges) Start(args ...interface{}) float64 {
	val := t.Call("start", args...)
	return val.Float()
}
