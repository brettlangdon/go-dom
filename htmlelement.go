// Code generated DO NOT EDIT
// htmlelement.go
package dom

import "syscall/js"

type HTMLElementIFace interface {
	GetAccessKey() string
	SetAccessKey(string)
	GetAccessKeyLabel() string
	AddEventListener(args ...interface{})
	AppendChild(args ...interface{}) Node
	AttachShadow(args ...interface{}) ShadowRoot
	GetAttributes() NamedNodeMap
	GetAutocapitalize() string
	SetAutocapitalize(string)
	GetBaseURI() string
	Blur(args ...interface{})
	GetChildNodes() NodeList
	GetClassList() DOMTokenList
	GetClassName() string
	SetClassName(string)
	Click(args ...interface{})
	CloneNode(args ...interface{}) Node
	Closest(args ...interface{}) Element
	CompareDocumentPosition(args ...interface{}) int
	Contains(args ...interface{}) bool
	GetContentEditable() string
	SetContentEditable(string)
	GetDataset() DOMStringMap
	GetDir() string
	SetDir(string)
	DispatchEvent(args ...interface{}) bool
	GetDraggable() bool
	SetDraggable(bool)
	GetEnterKeyHint() string
	SetEnterKeyHint(string)
	GetFirstChild() Node
	Focus(args ...interface{})
	GetAttribute(args ...interface{}) string
	GetAttributeNS(args ...interface{}) string
	GetAttributeNames(args ...interface{})
	GetAttributeNode(args ...interface{}) Attr
	GetAttributeNodeNS(args ...interface{}) Attr
	GetElementsByClassName(args ...interface{}) HTMLCollection
	GetElementsByTagName(args ...interface{}) HTMLCollection
	GetElementsByTagNameNS(args ...interface{}) HTMLCollection
	GetRootNode(args ...interface{}) Node
	HasAttribute(args ...interface{}) bool
	HasAttributeNS(args ...interface{}) bool
	HasAttributes(args ...interface{}) bool
	HasChildNodes(args ...interface{}) bool
	GetHidden() bool
	SetHidden(bool)
	GetId() string
	SetId(string)
	GetInnerText() string
	SetInnerText(string)
	GetInputMode() string
	SetInputMode(string)
	InsertAdjacentElement(args ...interface{}) Element
	InsertAdjacentText(args ...interface{})
	InsertBefore(args ...interface{}) Node
	GetIsConnected() bool
	GetIsContentEditable() bool
	IsDefaultNamespace(args ...interface{}) bool
	IsEqualNode(args ...interface{}) bool
	IsSameNode(args ...interface{}) bool
	GetLang() string
	SetLang(string)
	GetLastChild() Node
	GetLocalName() string
	LookupNamespaceURI(args ...interface{}) string
	LookupPrefix(args ...interface{}) string
	Matches(args ...interface{}) bool
	GetNamespaceURI() string
	GetNextSibling() Node
	GetNodeName() string
	GetNodeType() int
	GetNodeValue() string
	SetNodeValue(string)
	GetNonce() string
	SetNonce(string)
	Normalize(args ...interface{})
	GetOnabort() EventHandler
	SetOnabort(EventHandler)
	GetOnauxclick() EventHandler
	SetOnauxclick(EventHandler)
	GetOnblur() EventHandler
	SetOnblur(EventHandler)
	GetOncancel() EventHandler
	SetOncancel(EventHandler)
	GetOncanplay() EventHandler
	SetOncanplay(EventHandler)
	GetOncanplaythrough() EventHandler
	SetOncanplaythrough(EventHandler)
	GetOnchange() EventHandler
	SetOnchange(EventHandler)
	GetOnclick() EventHandler
	SetOnclick(EventHandler)
	GetOnclose() EventHandler
	SetOnclose(EventHandler)
	GetOncontextmenu() EventHandler
	SetOncontextmenu(EventHandler)
	GetOncopy() EventHandler
	SetOncopy(EventHandler)
	GetOncuechange() EventHandler
	SetOncuechange(EventHandler)
	GetOncut() EventHandler
	SetOncut(EventHandler)
	GetOndblclick() EventHandler
	SetOndblclick(EventHandler)
	GetOndrag() EventHandler
	SetOndrag(EventHandler)
	GetOndragend() EventHandler
	SetOndragend(EventHandler)
	GetOndragenter() EventHandler
	SetOndragenter(EventHandler)
	GetOndragexit() EventHandler
	SetOndragexit(EventHandler)
	GetOndragleave() EventHandler
	SetOndragleave(EventHandler)
	GetOndragover() EventHandler
	SetOndragover(EventHandler)
	GetOndragstart() EventHandler
	SetOndragstart(EventHandler)
	GetOndrop() EventHandler
	SetOndrop(EventHandler)
	GetOndurationchange() EventHandler
	SetOndurationchange(EventHandler)
	GetOnemptied() EventHandler
	SetOnemptied(EventHandler)
	GetOnended() EventHandler
	SetOnended(EventHandler)
	GetOnerror() OnErrorEventHandler
	SetOnerror(OnErrorEventHandler)
	GetOnfocus() EventHandler
	SetOnfocus(EventHandler)
	GetOninput() EventHandler
	SetOninput(EventHandler)
	GetOninvalid() EventHandler
	SetOninvalid(EventHandler)
	GetOnkeydown() EventHandler
	SetOnkeydown(EventHandler)
	GetOnkeypress() EventHandler
	SetOnkeypress(EventHandler)
	GetOnkeyup() EventHandler
	SetOnkeyup(EventHandler)
	GetOnload() EventHandler
	SetOnload(EventHandler)
	GetOnloadeddata() EventHandler
	SetOnloadeddata(EventHandler)
	GetOnloadedmetadata() EventHandler
	SetOnloadedmetadata(EventHandler)
	GetOnloadend() EventHandler
	SetOnloadend(EventHandler)
	GetOnloadstart() EventHandler
	SetOnloadstart(EventHandler)
	GetOnmousedown() EventHandler
	SetOnmousedown(EventHandler)
	GetOnmouseenter() EventHandler
	SetOnmouseenter(EventHandler)
	GetOnmouseleave() EventHandler
	SetOnmouseleave(EventHandler)
	GetOnmousemove() EventHandler
	SetOnmousemove(EventHandler)
	GetOnmouseout() EventHandler
	SetOnmouseout(EventHandler)
	GetOnmouseover() EventHandler
	SetOnmouseover(EventHandler)
	GetOnmouseup() EventHandler
	SetOnmouseup(EventHandler)
	GetOnpaste() EventHandler
	SetOnpaste(EventHandler)
	GetOnpause() EventHandler
	SetOnpause(EventHandler)
	GetOnplay() EventHandler
	SetOnplay(EventHandler)
	GetOnplaying() EventHandler
	SetOnplaying(EventHandler)
	GetOnprogress() EventHandler
	SetOnprogress(EventHandler)
	GetOnratechange() EventHandler
	SetOnratechange(EventHandler)
	GetOnreset() EventHandler
	SetOnreset(EventHandler)
	GetOnresize() EventHandler
	SetOnresize(EventHandler)
	GetOnscroll() EventHandler
	SetOnscroll(EventHandler)
	GetOnsecuritypolicyviolation() EventHandler
	SetOnsecuritypolicyviolation(EventHandler)
	GetOnseeked() EventHandler
	SetOnseeked(EventHandler)
	GetOnseeking() EventHandler
	SetOnseeking(EventHandler)
	GetOnselect() EventHandler
	SetOnselect(EventHandler)
	GetOnstalled() EventHandler
	SetOnstalled(EventHandler)
	GetOnsubmit() EventHandler
	SetOnsubmit(EventHandler)
	GetOnsuspend() EventHandler
	SetOnsuspend(EventHandler)
	GetOntimeupdate() EventHandler
	SetOntimeupdate(EventHandler)
	GetOntoggle() EventHandler
	SetOntoggle(EventHandler)
	GetOnvolumechange() EventHandler
	SetOnvolumechange(EventHandler)
	GetOnwaiting() EventHandler
	SetOnwaiting(EventHandler)
	GetOnwheel() EventHandler
	SetOnwheel(EventHandler)
	GetOwnerDocument() Document
	GetParentElement() Element
	GetParentNode() Node
	GetPrefix() string
	GetPreviousSibling() Node
	RemoveAttribute(args ...interface{})
	RemoveAttributeNS(args ...interface{})
	RemoveAttributeNode(args ...interface{}) Attr
	RemoveChild(args ...interface{}) Node
	RemoveEventListener(args ...interface{})
	ReplaceChild(args ...interface{}) Node
	SetAttribute(args ...interface{})
	SetAttributeNS(args ...interface{})
	SetAttributeNode(args ...interface{}) Attr
	SetAttributeNodeNS(args ...interface{}) Attr
	GetShadowRoot() ShadowRoot
	GetSlot() string
	SetSlot(string)
	GetSpellcheck() bool
	SetSpellcheck(bool)
	GetStyle() CSSStyleDeclaration
	GetTabIndex() float64
	SetTabIndex(float64)
	GetTagName() string
	GetTextContent() string
	SetTextContent(string)
	GetTitle() string
	SetTitle(string)
	ToggleAttribute(args ...interface{}) bool
	GetTranslate() bool
	SetTranslate(bool)
	WebkitMatchesSelector(args ...interface{}) bool
}
type HTMLElement struct {
	Value
	Element
	Node
	EventTarget
}

func JSValueToHTMLElement(val js.Value) HTMLElement { return HTMLElement{Value: Value{Value: val}} }
func (v Value) AsHTMLElement() HTMLElement          { return HTMLElement{Value: v} }
func (h HTMLElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLElement) Blur(args ...interface{}) {
	h.Call("blur", args...)
}
func (h HTMLElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLElement) GetContentEditable() string {
	val := h.Get("contentEditable")
	return val.String()
}
func (h HTMLElement) SetContentEditable(val string) {
	h.Set("contentEditable", val)
}
func (h HTMLElement) GetDataset() DOMStringMap {
	val := h.Get("dataset")
	return JSValueToDOMStringMap(val.JSValue())
}
func (h HTMLElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLElement) GetEnterKeyHint() string {
	val := h.Get("enterKeyHint")
	return val.String()
}
func (h HTMLElement) SetEnterKeyHint(val string) {
	h.Set("enterKeyHint", val)
}
func (h HTMLElement) Focus(args ...interface{}) {
	h.Call("focus", args...)
}
func (h HTMLElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLElement) GetInputMode() string {
	val := h.Get("inputMode")
	return val.String()
}
func (h HTMLElement) SetInputMode(val string) {
	h.Set("inputMode", val)
}
func (h HTMLElement) GetIsContentEditable() bool {
	val := h.Get("isContentEditable")
	return val.Bool()
}
func (h HTMLElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLElement) GetNonce() string {
	val := h.Get("nonce")
	return val.String()
}
func (h HTMLElement) SetNonce(val string) {
	h.Set("nonce", val)
}
func (h HTMLElement) GetOnabort() EventHandler {
	val := h.Get("onabort")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnabort(val EventHandler) {
	h.Set("onabort", val)
}
func (h HTMLElement) GetOnauxclick() EventHandler {
	val := h.Get("onauxclick")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnauxclick(val EventHandler) {
	h.Set("onauxclick", val)
}
func (h HTMLElement) GetOnblur() EventHandler {
	val := h.Get("onblur")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnblur(val EventHandler) {
	h.Set("onblur", val)
}
func (h HTMLElement) GetOncancel() EventHandler {
	val := h.Get("oncancel")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOncancel(val EventHandler) {
	h.Set("oncancel", val)
}
func (h HTMLElement) GetOncanplay() EventHandler {
	val := h.Get("oncanplay")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOncanplay(val EventHandler) {
	h.Set("oncanplay", val)
}
func (h HTMLElement) GetOncanplaythrough() EventHandler {
	val := h.Get("oncanplaythrough")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOncanplaythrough(val EventHandler) {
	h.Set("oncanplaythrough", val)
}
func (h HTMLElement) GetOnchange() EventHandler {
	val := h.Get("onchange")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnchange(val EventHandler) {
	h.Set("onchange", val)
}
func (h HTMLElement) GetOnclick() EventHandler {
	val := h.Get("onclick")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnclick(val EventHandler) {
	h.Set("onclick", val)
}
func (h HTMLElement) GetOnclose() EventHandler {
	val := h.Get("onclose")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnclose(val EventHandler) {
	h.Set("onclose", val)
}
func (h HTMLElement) GetOncontextmenu() EventHandler {
	val := h.Get("oncontextmenu")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOncontextmenu(val EventHandler) {
	h.Set("oncontextmenu", val)
}
func (h HTMLElement) GetOncopy() EventHandler {
	val := h.Get("oncopy")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOncopy(val EventHandler) {
	h.Set("oncopy", val)
}
func (h HTMLElement) GetOncuechange() EventHandler {
	val := h.Get("oncuechange")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOncuechange(val EventHandler) {
	h.Set("oncuechange", val)
}
func (h HTMLElement) GetOncut() EventHandler {
	val := h.Get("oncut")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOncut(val EventHandler) {
	h.Set("oncut", val)
}
func (h HTMLElement) GetOndblclick() EventHandler {
	val := h.Get("ondblclick")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOndblclick(val EventHandler) {
	h.Set("ondblclick", val)
}
func (h HTMLElement) GetOndrag() EventHandler {
	val := h.Get("ondrag")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOndrag(val EventHandler) {
	h.Set("ondrag", val)
}
func (h HTMLElement) GetOndragend() EventHandler {
	val := h.Get("ondragend")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOndragend(val EventHandler) {
	h.Set("ondragend", val)
}
func (h HTMLElement) GetOndragenter() EventHandler {
	val := h.Get("ondragenter")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOndragenter(val EventHandler) {
	h.Set("ondragenter", val)
}
func (h HTMLElement) GetOndragexit() EventHandler {
	val := h.Get("ondragexit")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOndragexit(val EventHandler) {
	h.Set("ondragexit", val)
}
func (h HTMLElement) GetOndragleave() EventHandler {
	val := h.Get("ondragleave")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOndragleave(val EventHandler) {
	h.Set("ondragleave", val)
}
func (h HTMLElement) GetOndragover() EventHandler {
	val := h.Get("ondragover")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOndragover(val EventHandler) {
	h.Set("ondragover", val)
}
func (h HTMLElement) GetOndragstart() EventHandler {
	val := h.Get("ondragstart")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOndragstart(val EventHandler) {
	h.Set("ondragstart", val)
}
func (h HTMLElement) GetOndrop() EventHandler {
	val := h.Get("ondrop")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOndrop(val EventHandler) {
	h.Set("ondrop", val)
}
func (h HTMLElement) GetOndurationchange() EventHandler {
	val := h.Get("ondurationchange")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOndurationchange(val EventHandler) {
	h.Set("ondurationchange", val)
}
func (h HTMLElement) GetOnemptied() EventHandler {
	val := h.Get("onemptied")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnemptied(val EventHandler) {
	h.Set("onemptied", val)
}
func (h HTMLElement) GetOnended() EventHandler {
	val := h.Get("onended")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnended(val EventHandler) {
	h.Set("onended", val)
}
func (h HTMLElement) GetOnerror() OnErrorEventHandler {
	val := h.Get("onerror")
	return JSValueToOnErrorEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnerror(val OnErrorEventHandler) {
	h.Set("onerror", val)
}
func (h HTMLElement) GetOnfocus() EventHandler {
	val := h.Get("onfocus")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnfocus(val EventHandler) {
	h.Set("onfocus", val)
}
func (h HTMLElement) GetOninput() EventHandler {
	val := h.Get("oninput")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOninput(val EventHandler) {
	h.Set("oninput", val)
}
func (h HTMLElement) GetOninvalid() EventHandler {
	val := h.Get("oninvalid")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOninvalid(val EventHandler) {
	h.Set("oninvalid", val)
}
func (h HTMLElement) GetOnkeydown() EventHandler {
	val := h.Get("onkeydown")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnkeydown(val EventHandler) {
	h.Set("onkeydown", val)
}
func (h HTMLElement) GetOnkeypress() EventHandler {
	val := h.Get("onkeypress")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnkeypress(val EventHandler) {
	h.Set("onkeypress", val)
}
func (h HTMLElement) GetOnkeyup() EventHandler {
	val := h.Get("onkeyup")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnkeyup(val EventHandler) {
	h.Set("onkeyup", val)
}
func (h HTMLElement) GetOnload() EventHandler {
	val := h.Get("onload")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnload(val EventHandler) {
	h.Set("onload", val)
}
func (h HTMLElement) GetOnloadeddata() EventHandler {
	val := h.Get("onloadeddata")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnloadeddata(val EventHandler) {
	h.Set("onloadeddata", val)
}
func (h HTMLElement) GetOnloadedmetadata() EventHandler {
	val := h.Get("onloadedmetadata")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnloadedmetadata(val EventHandler) {
	h.Set("onloadedmetadata", val)
}
func (h HTMLElement) GetOnloadend() EventHandler {
	val := h.Get("onloadend")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnloadend(val EventHandler) {
	h.Set("onloadend", val)
}
func (h HTMLElement) GetOnloadstart() EventHandler {
	val := h.Get("onloadstart")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnloadstart(val EventHandler) {
	h.Set("onloadstart", val)
}
func (h HTMLElement) GetOnmousedown() EventHandler {
	val := h.Get("onmousedown")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnmousedown(val EventHandler) {
	h.Set("onmousedown", val)
}
func (h HTMLElement) GetOnmouseenter() EventHandler {
	val := h.Get("onmouseenter")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnmouseenter(val EventHandler) {
	h.Set("onmouseenter", val)
}
func (h HTMLElement) GetOnmouseleave() EventHandler {
	val := h.Get("onmouseleave")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnmouseleave(val EventHandler) {
	h.Set("onmouseleave", val)
}
func (h HTMLElement) GetOnmousemove() EventHandler {
	val := h.Get("onmousemove")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnmousemove(val EventHandler) {
	h.Set("onmousemove", val)
}
func (h HTMLElement) GetOnmouseout() EventHandler {
	val := h.Get("onmouseout")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnmouseout(val EventHandler) {
	h.Set("onmouseout", val)
}
func (h HTMLElement) GetOnmouseover() EventHandler {
	val := h.Get("onmouseover")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnmouseover(val EventHandler) {
	h.Set("onmouseover", val)
}
func (h HTMLElement) GetOnmouseup() EventHandler {
	val := h.Get("onmouseup")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnmouseup(val EventHandler) {
	h.Set("onmouseup", val)
}
func (h HTMLElement) GetOnpaste() EventHandler {
	val := h.Get("onpaste")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnpaste(val EventHandler) {
	h.Set("onpaste", val)
}
func (h HTMLElement) GetOnpause() EventHandler {
	val := h.Get("onpause")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnpause(val EventHandler) {
	h.Set("onpause", val)
}
func (h HTMLElement) GetOnplay() EventHandler {
	val := h.Get("onplay")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnplay(val EventHandler) {
	h.Set("onplay", val)
}
func (h HTMLElement) GetOnplaying() EventHandler {
	val := h.Get("onplaying")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnplaying(val EventHandler) {
	h.Set("onplaying", val)
}
func (h HTMLElement) GetOnprogress() EventHandler {
	val := h.Get("onprogress")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnprogress(val EventHandler) {
	h.Set("onprogress", val)
}
func (h HTMLElement) GetOnratechange() EventHandler {
	val := h.Get("onratechange")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnratechange(val EventHandler) {
	h.Set("onratechange", val)
}
func (h HTMLElement) GetOnreset() EventHandler {
	val := h.Get("onreset")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnreset(val EventHandler) {
	h.Set("onreset", val)
}
func (h HTMLElement) GetOnresize() EventHandler {
	val := h.Get("onresize")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnresize(val EventHandler) {
	h.Set("onresize", val)
}
func (h HTMLElement) GetOnscroll() EventHandler {
	val := h.Get("onscroll")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnscroll(val EventHandler) {
	h.Set("onscroll", val)
}
func (h HTMLElement) GetOnsecuritypolicyviolation() EventHandler {
	val := h.Get("onsecuritypolicyviolation")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnsecuritypolicyviolation(val EventHandler) {
	h.Set("onsecuritypolicyviolation", val)
}
func (h HTMLElement) GetOnseeked() EventHandler {
	val := h.Get("onseeked")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnseeked(val EventHandler) {
	h.Set("onseeked", val)
}
func (h HTMLElement) GetOnseeking() EventHandler {
	val := h.Get("onseeking")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnseeking(val EventHandler) {
	h.Set("onseeking", val)
}
func (h HTMLElement) GetOnselect() EventHandler {
	val := h.Get("onselect")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnselect(val EventHandler) {
	h.Set("onselect", val)
}
func (h HTMLElement) GetOnstalled() EventHandler {
	val := h.Get("onstalled")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnstalled(val EventHandler) {
	h.Set("onstalled", val)
}
func (h HTMLElement) GetOnsubmit() EventHandler {
	val := h.Get("onsubmit")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnsubmit(val EventHandler) {
	h.Set("onsubmit", val)
}
func (h HTMLElement) GetOnsuspend() EventHandler {
	val := h.Get("onsuspend")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnsuspend(val EventHandler) {
	h.Set("onsuspend", val)
}
func (h HTMLElement) GetOntimeupdate() EventHandler {
	val := h.Get("ontimeupdate")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOntimeupdate(val EventHandler) {
	h.Set("ontimeupdate", val)
}
func (h HTMLElement) GetOntoggle() EventHandler {
	val := h.Get("ontoggle")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOntoggle(val EventHandler) {
	h.Set("ontoggle", val)
}
func (h HTMLElement) GetOnvolumechange() EventHandler {
	val := h.Get("onvolumechange")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnvolumechange(val EventHandler) {
	h.Set("onvolumechange", val)
}
func (h HTMLElement) GetOnwaiting() EventHandler {
	val := h.Get("onwaiting")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnwaiting(val EventHandler) {
	h.Set("onwaiting", val)
}
func (h HTMLElement) GetOnwheel() EventHandler {
	val := h.Get("onwheel")
	return JSValueToEventHandler(val.JSValue())
}
func (h HTMLElement) SetOnwheel(val EventHandler) {
	h.Set("onwheel", val)
}
func (h HTMLElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLElement) GetStyle() CSSStyleDeclaration {
	val := h.Get("style")
	return JSValueToCSSStyleDeclaration(val.JSValue())
}
func (h HTMLElement) GetTabIndex() float64 {
	val := h.Get("tabIndex")
	return val.Float()
}
func (h HTMLElement) SetTabIndex(val float64) {
	h.Set("tabIndex", val)
}
func (h HTMLElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
