// Code generated DO NOT EDIT
// nodefilter.go
package dom

import "syscall/js"

const (
	NodeFilterFILTER_ACCEPT               int     = 1
	NodeFilterFILTER_REJECT               int     = 2
	NodeFilterFILTER_SKIP                 int     = 3
	NodeFilterSHOW_ALL                    float64 = 0xFFFFFFFF
	NodeFilterSHOW_ATTRIBUTE              float64 = 0x2
	NodeFilterSHOW_CDATA_SECTION          float64 = 0x8
	NodeFilterSHOW_COMMENT                float64 = 0x80
	NodeFilterSHOW_DOCUMENT               float64 = 0x100
	NodeFilterSHOW_DOCUMENT_FRAGMENT      float64 = 0x400
	NodeFilterSHOW_DOCUMENT_TYPE          float64 = 0x200
	NodeFilterSHOW_ELEMENT                float64 = 0x1
	NodeFilterSHOW_ENTITY                 float64 = 0x20
	NodeFilterSHOW_ENTITY_REFERENCE       float64 = 0x10
	NodeFilterSHOW_NOTATION               float64 = 0x800
	NodeFilterSHOW_PROCESSING_INSTRUCTION float64 = 0x40
	NodeFilterSHOW_TEXT                   float64 = 0x4
)

type NodeFilterAcceptNodeCallback func(node Node)
type NodeFilterAcceptNode struct {
	Callback
}

func jsValueToNodeFilterAcceptNode(val js.Value) NodeFilterAcceptNode {
	return NodeFilterAcceptNode{Callback: jsValueToCallback(val)}
}
func NewNodeFilterAcceptNode(c NodeFilterAcceptNodeCallback) NodeFilterAcceptNode {
	callback := js.NewCallback(func(args []js.Value) {
		node := jsValueToNode(args[0])
		c(node)
	})
	return NodeFilterAcceptNode{Callback: Callback{Callback: callback}}
}

type NodeFilter struct {
	Value
	AcceptNode NodeFilterAcceptNodeCallback
}

func jsValueToNodeFilter(val js.Value) NodeFilter { return NodeFilter{Value: Value{Value: val}} }
func (v Value) AsNodeFilter() NodeFilter          { return NodeFilter{Value: v} }
