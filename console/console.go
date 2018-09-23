// Code generated DO NOT EDIT
// console/console.go
package console

import "syscall/js"

var value Value

func init()                              { value = Value{Value: js.Global().Get("console")} }
func Clear(args ...interface{})          { return value.Call("clear", args...) }
func Table(args ...interface{})          { return value.Call("table", args...) }
func Warn(args ...interface{})           { return value.Call("warn", args...) }
func Dirxml(args ...interface{})         { return value.Call("dirxml", args...) }
func Group(args ...interface{})          { return value.Call("group", args...) }
func Time(args ...interface{})           { return value.Call("time", args...) }
func Trace(args ...interface{})          { return value.Call("trace", args...) }
func CountReset(args ...interface{})     { return value.Call("countReset", args...) }
func GroupEnd(args ...interface{})       { return value.Call("groupEnd", args...) }
func Assert(args ...interface{})         { return value.Call("assert", args...) }
func Debug(args ...interface{})          { return value.Call("debug", args...) }
func Log(args ...interface{})            { return value.Call("log", args...) }
func Dir(args ...interface{})            { return value.Call("dir", args...) }
func TimeLog(args ...interface{})        { return value.Call("timeLog", args...) }
func Error(args ...interface{})          { return value.Call("error", args...) }
func Info(args ...interface{})           { return value.Call("info", args...) }
func Count(args ...interface{})          { return value.Call("count", args...) }
func GroupCollapsed(args ...interface{}) { return value.Call("groupCollapsed", args...) }
func TimeEnd(args ...interface{})        { return value.Call("timeEnd", args...) }
