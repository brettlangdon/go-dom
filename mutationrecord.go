// Code generated DO NOT EDIT
// mutationrecord.go
package dom

import "syscall/js"

type MutationRecordIFace interface {
	GetAddedNodes() NodeList
	GetAttributeName() string
	GetAttributeNamespace() string
	GetNextSibling() Node
	GetOldValue() string
	GetPreviousSibling() Node
	GetRemovedNodes() NodeList
	GetTarget() Node
	GetType() string
}
type MutationRecord struct {
	Value
}

func jsValueToMutationRecord(val js.Value) MutationRecord {
	return MutationRecord{Value: Value{Value: val}}
}
func (v Value) AsMutationRecord() MutationRecord { return MutationRecord{Value: v} }
func (m MutationRecord) GetAddedNodes() NodeList {
	val := m.Get("addedNodes")
	return jsValueToNodeList(val.JSValue())
}
func (m MutationRecord) GetAttributeName() string {
	val := m.Get("attributeName")
	return val.String()
}
func (m MutationRecord) GetAttributeNamespace() string {
	val := m.Get("attributeNamespace")
	return val.String()
}
func (m MutationRecord) GetNextSibling() Node {
	val := m.Get("nextSibling")
	return jsValueToNode(val.JSValue())
}
func (m MutationRecord) GetOldValue() string {
	val := m.Get("oldValue")
	return val.String()
}
func (m MutationRecord) GetPreviousSibling() Node {
	val := m.Get("previousSibling")
	return jsValueToNode(val.JSValue())
}
func (m MutationRecord) GetRemovedNodes() NodeList {
	val := m.Get("removedNodes")
	return jsValueToNodeList(val.JSValue())
}
func (m MutationRecord) GetTarget() Node {
	val := m.Get("target")
	return jsValueToNode(val.JSValue())
}
func (m MutationRecord) GetType() string {
	val := m.Get("type")
	return val.String()
}
