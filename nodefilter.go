// Code generated DO NOT EDIT
// nodefilter.go
package dom

import "syscall/js"

const (
	NodeFilterFILTER_ACCEPT               int = 1
	NodeFilterFILTER_REJECT               int = 2
	NodeFilterFILTER_SKIP                 int = 3
	NodeFilterSHOW_ALL                    int = 0xFFFFFFFF
	NodeFilterSHOW_ATTRIBUTE              int = 0x2
	NodeFilterSHOW_CDATA_SECTION          int = 0x8
	NodeFilterSHOW_COMMENT                int = 0x80
	NodeFilterSHOW_DOCUMENT               int = 0x100
	NodeFilterSHOW_DOCUMENT_FRAGMENT      int = 0x400
	NodeFilterSHOW_DOCUMENT_TYPE          int = 0x200
	NodeFilterSHOW_ELEMENT                int = 0x1
	NodeFilterSHOW_ENTITY                 int = 0x20
	NodeFilterSHOW_ENTITY_REFERENCE       int = 0x10
	NodeFilterSHOW_NOTATION               int = 0x800
	NodeFilterSHOW_PROCESSING_INSTRUCTION int = 0x40
	NodeFilterSHOW_TEXT                   int = 0x4
)

type NodeFilterAcceptNodeCallback func(node Node)
type NodeFilterAcceptNode struct {
	Callback
}

func JSValueToNodeFilterAcceptNode(val js.Value) NodeFilterAcceptNode {
	return NodeFilterAcceptNode{Callback: JSValueToCallback(val)}
}
func NewNodeFilterAcceptNode(c NodeFilterAcceptNodeCallback) NodeFilterAcceptNode {
	callback := js.NewCallback(func(args []js.Value) {
		node := JSValueToNode(args[0])
		c(node)
	})
	return NodeFilterAcceptNode{Callback: Callback{Callback: callback}}
}

type NodeFilter struct {
	Value
	AcceptNode NodeFilterAcceptNodeCallback
}

func JSValueToNodeFilter(val js.Value) NodeFilter { return NodeFilter{Value: Value{Value: val}} }
func (v Value) AsNodeFilter() NodeFilter          { return NodeFilter{Value: v} }
