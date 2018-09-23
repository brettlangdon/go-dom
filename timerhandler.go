// Code generated DO NOT EDIT
// timerhandler.go
package dom

import "syscall/js"

type TimerHandler Value

func JSValueToTimerHandler(val js.Value) TimerHandler {
	return TimerHandler(JSValueToValue(val))
}
