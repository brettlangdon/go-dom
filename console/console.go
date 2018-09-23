// Code generated DO NOT EDIT
// console/console.go
package console

import dom "github.com/brettlangdon/go-dom/v1"
import "syscall/js"

var value dom.Value

func init() { value = dom.JSValueToValue(js.Global().Get("console")) }
func Assert(args ...interface{}) {
	value.Call("assert", args...)
}
func Clear(args ...interface{}) {
	value.Call("clear", args...)
}
func Count(args ...interface{}) {
	value.Call("count", args...)
}
func CountReset(args ...interface{}) {
	value.Call("countReset", args...)
}
func Debug(args ...interface{}) {
	value.Call("debug", args...)
}
func Dir(args ...interface{}) {
	value.Call("dir", args...)
}
func Dirxml(args ...interface{}) {
	value.Call("dirxml", args...)
}
func Error(args ...interface{}) {
	value.Call("error", args...)
}
func Group(args ...interface{}) {
	value.Call("group", args...)
}
func GroupCollapsed(args ...interface{}) {
	value.Call("groupCollapsed", args...)
}
func GroupEnd(args ...interface{}) {
	value.Call("groupEnd", args...)
}
func Info(args ...interface{}) {
	value.Call("info", args...)
}
func Log(args ...interface{}) {
	value.Call("log", args...)
}
func Table(args ...interface{}) {
	value.Call("table", args...)
}
func Time(args ...interface{}) {
	value.Call("time", args...)
}
func TimeEnd(args ...interface{}) {
	value.Call("timeEnd", args...)
}
func TimeLog(args ...interface{}) {
	value.Call("timeLog", args...)
}
func Trace(args ...interface{}) {
	value.Call("trace", args...)
}
func Warn(args ...interface{}) {
	value.Call("warn", args...)
}
