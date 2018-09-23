// Code generated DO NOT EDIT
// timerhandler.go
package dom

import "syscall/js"

type TimerHandler Value

func jsValueToTimerHandler(val js.Value) TimerHandler {
	return TimerHandler(jsValueToValue(val))
}
