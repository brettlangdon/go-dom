// Code generated DO NOT EDIT
// messagechannel.go
package dom

import "syscall/js"

type MessageChannelIFace interface {
	GetPort1() MessagePort
	GetPort2() MessagePort
}
type MessageChannel struct {
	Value
}

func JSValueToMessageChannel(val js.Value) MessageChannel {
	return MessageChannel{Value: JSValueToValue(val)}
}
func (v Value) AsMessageChannel() MessageChannel { return MessageChannel{Value: v} }
func NewMessageChannel(args ...interface{}) MessageChannel {
	return MessageChannel{Value: JSValueToValue(js.Global().Get("MessageChannel").New(args...))}
}
func (m MessageChannel) GetPort1() MessagePort {
	val := m.Get("port1")
	return JSValueToMessagePort(val.JSValue())
}
func (m MessageChannel) GetPort2() MessagePort {
	val := m.Get("port2")
	return JSValueToMessagePort(val.JSValue())
}
