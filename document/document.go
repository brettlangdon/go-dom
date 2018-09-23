// Code generated DO NOT EDIT
// document/document.go
package document

import dom "github.com/brettlangdon/go-dom"
import "syscall/js"

var value dom.Document

func init() { value = dom.JSValueToDocument(js.Global().Get("document")) }
func AddEventListener(args ...interface{}) {
	value.Call("addEventListener", args...)
}
func AdoptNode(args ...interface{}) dom.Node {
	val := value.Call("adoptNode", args...)
	return dom.JSValueToNode(val.JSValue())
}
func Append(args ...interface{}) {
	value.Call("append", args...)
}
func AppendChild(args ...interface{}) dom.Node {
	val := value.Call("appendChild", args...)
	return dom.JSValueToNode(val.JSValue())
}
func GetBaseURI() string {
	return value.GetBaseURI()
}
func GetCharacterSet() string {
	return value.GetCharacterSet()
}
func GetCharset() string {
	return value.GetCharset()
}
func GetChildElementCount() int {
	return value.GetChildElementCount()
}
func GetChildNodes() dom.NodeList {
	return value.GetChildNodes()
}
func GetChildren() dom.HTMLCollection {
	return value.GetChildren()
}
func CloneNode(args ...interface{}) dom.Node {
	val := value.Call("cloneNode", args...)
	return dom.JSValueToNode(val.JSValue())
}
func CompareDocumentPosition(args ...interface{}) int {
	val := value.Call("compareDocumentPosition", args...)
	return val.Int()
}
func GetCompatMode() string {
	return value.GetCompatMode()
}
func Contains(args ...interface{}) bool {
	val := value.Call("contains", args...)
	return val.Bool()
}
func GetContentType() string {
	return value.GetContentType()
}
func CreateAttribute(args ...interface{}) dom.Attr {
	val := value.Call("createAttribute", args...)
	return dom.JSValueToAttr(val.JSValue())
}
func CreateAttributeNS(args ...interface{}) dom.Attr {
	val := value.Call("createAttributeNS", args...)
	return dom.JSValueToAttr(val.JSValue())
}
func CreateCDATASection(args ...interface{}) dom.CDATASection {
	val := value.Call("createCDATASection", args...)
	return dom.JSValueToCDATASection(val.JSValue())
}
func CreateComment(args ...interface{}) dom.Comment {
	val := value.Call("createComment", args...)
	return dom.JSValueToComment(val.JSValue())
}
func CreateDocumentFragment(args ...interface{}) dom.DocumentFragment {
	val := value.Call("createDocumentFragment", args...)
	return dom.JSValueToDocumentFragment(val.JSValue())
}
func CreateElement(args ...interface{}) dom.Element {
	val := value.Call("createElement", args...)
	return dom.JSValueToElement(val.JSValue())
}
func CreateElementNS(args ...interface{}) dom.Element {
	val := value.Call("createElementNS", args...)
	return dom.JSValueToElement(val.JSValue())
}
func CreateEvent(args ...interface{}) dom.Event {
	val := value.Call("createEvent", args...)
	return dom.JSValueToEvent(val.JSValue())
}
func CreateNodeIterator(args ...interface{}) dom.NodeIterator {
	val := value.Call("createNodeIterator", args...)
	return dom.JSValueToNodeIterator(val.JSValue())
}
func CreateProcessingInstruction(args ...interface{}) dom.ProcessingInstruction {
	val := value.Call("createProcessingInstruction", args...)
	return dom.JSValueToProcessingInstruction(val.JSValue())
}
func CreateRange(args ...interface{}) dom.Range {
	val := value.Call("createRange", args...)
	return dom.JSValueToRange(val.JSValue())
}
func CreateTextNode(args ...interface{}) dom.Text {
	val := value.Call("createTextNode", args...)
	return dom.JSValueToText(val.JSValue())
}
func CreateTreeWalker(args ...interface{}) dom.TreeWalker {
	val := value.Call("createTreeWalker", args...)
	return dom.JSValueToTreeWalker(val.JSValue())
}
func DispatchEvent(args ...interface{}) bool {
	val := value.Call("dispatchEvent", args...)
	return val.Bool()
}
func GetDoctype() dom.DocumentType {
	return value.GetDoctype()
}
func GetDocumentElement() dom.Element {
	return value.GetDocumentElement()
}
func GetDocumentURI() string {
	return value.GetDocumentURI()
}
func GetFirstChild() dom.Node {
	return value.GetFirstChild()
}
func GetFirstElementChild() dom.Element {
	return value.GetFirstElementChild()
}
func GetElementById(args ...interface{}) dom.Element {
	val := value.Call("getElementById", args...)
	return dom.JSValueToElement(val.JSValue())
}
func GetElementsByClassName(args ...interface{}) dom.HTMLCollection {
	val := value.Call("getElementsByClassName", args...)
	return dom.JSValueToHTMLCollection(val.JSValue())
}
func GetElementsByTagName(args ...interface{}) dom.HTMLCollection {
	val := value.Call("getElementsByTagName", args...)
	return dom.JSValueToHTMLCollection(val.JSValue())
}
func GetElementsByTagNameNS(args ...interface{}) dom.HTMLCollection {
	val := value.Call("getElementsByTagNameNS", args...)
	return dom.JSValueToHTMLCollection(val.JSValue())
}
func GetRootNode(args ...interface{}) dom.Node {
	val := value.Call("getRootNode", args...)
	return dom.JSValueToNode(val.JSValue())
}
func HasChildNodes(args ...interface{}) bool {
	val := value.Call("hasChildNodes", args...)
	return val.Bool()
}
func GetImplementation() dom.DOMImplementation {
	return value.GetImplementation()
}
func ImportNode(args ...interface{}) dom.Node {
	val := value.Call("importNode", args...)
	return dom.JSValueToNode(val.JSValue())
}
func GetInputEncoding() string {
	return value.GetInputEncoding()
}
func InsertBefore(args ...interface{}) dom.Node {
	val := value.Call("insertBefore", args...)
	return dom.JSValueToNode(val.JSValue())
}
func GetIsConnected() bool {
	return value.GetIsConnected()
}
func IsDefaultNamespace(args ...interface{}) bool {
	val := value.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func IsEqualNode(args ...interface{}) bool {
	val := value.Call("isEqualNode", args...)
	return val.Bool()
}
func IsSameNode(args ...interface{}) bool {
	val := value.Call("isSameNode", args...)
	return val.Bool()
}
func GetLastChild() dom.Node {
	return value.GetLastChild()
}
func GetLastElementChild() dom.Element {
	return value.GetLastElementChild()
}
func LookupNamespaceURI(args ...interface{}) string {
	val := value.Call("lookupNamespaceURI", args...)
	return val.String()
}
func LookupPrefix(args ...interface{}) string {
	val := value.Call("lookupPrefix", args...)
	return val.String()
}
func GetNextSibling() dom.Node {
	return value.GetNextSibling()
}
func GetNodeName() string {
	return value.GetNodeName()
}
func GetNodeType() int {
	return value.GetNodeType()
}
func GetNodeValue() string {
	return value.GetNodeValue()
}
func SetNodeValue(val string) { value.SetNodeValue(val) }
func Normalize(args ...interface{}) {
	value.Call("normalize", args...)
}
func GetOnabort() dom.EventHandler {
	return value.GetOnabort()
}
func SetOnabort(val dom.EventHandler) { value.SetOnabort(val) }
func GetOnauxclick() dom.EventHandler {
	return value.GetOnauxclick()
}
func SetOnauxclick(val dom.EventHandler) { value.SetOnauxclick(val) }
func GetOnblur() dom.EventHandler {
	return value.GetOnblur()
}
func SetOnblur(val dom.EventHandler) { value.SetOnblur(val) }
func GetOncancel() dom.EventHandler {
	return value.GetOncancel()
}
func SetOncancel(val dom.EventHandler) { value.SetOncancel(val) }
func GetOncanplay() dom.EventHandler {
	return value.GetOncanplay()
}
func SetOncanplay(val dom.EventHandler) { value.SetOncanplay(val) }
func GetOncanplaythrough() dom.EventHandler {
	return value.GetOncanplaythrough()
}
func SetOncanplaythrough(val dom.EventHandler) { value.SetOncanplaythrough(val) }
func GetOnchange() dom.EventHandler {
	return value.GetOnchange()
}
func SetOnchange(val dom.EventHandler) { value.SetOnchange(val) }
func GetOnclick() dom.EventHandler {
	return value.GetOnclick()
}
func SetOnclick(val dom.EventHandler) { value.SetOnclick(val) }
func GetOnclose() dom.EventHandler {
	return value.GetOnclose()
}
func SetOnclose(val dom.EventHandler) { value.SetOnclose(val) }
func GetOncontextmenu() dom.EventHandler {
	return value.GetOncontextmenu()
}
func SetOncontextmenu(val dom.EventHandler) { value.SetOncontextmenu(val) }
func GetOncopy() dom.EventHandler {
	return value.GetOncopy()
}
func SetOncopy(val dom.EventHandler) { value.SetOncopy(val) }
func GetOncuechange() dom.EventHandler {
	return value.GetOncuechange()
}
func SetOncuechange(val dom.EventHandler) { value.SetOncuechange(val) }
func GetOncut() dom.EventHandler {
	return value.GetOncut()
}
func SetOncut(val dom.EventHandler) { value.SetOncut(val) }
func GetOndblclick() dom.EventHandler {
	return value.GetOndblclick()
}
func SetOndblclick(val dom.EventHandler) { value.SetOndblclick(val) }
func GetOndrag() dom.EventHandler {
	return value.GetOndrag()
}
func SetOndrag(val dom.EventHandler) { value.SetOndrag(val) }
func GetOndragend() dom.EventHandler {
	return value.GetOndragend()
}
func SetOndragend(val dom.EventHandler) { value.SetOndragend(val) }
func GetOndragenter() dom.EventHandler {
	return value.GetOndragenter()
}
func SetOndragenter(val dom.EventHandler) { value.SetOndragenter(val) }
func GetOndragexit() dom.EventHandler {
	return value.GetOndragexit()
}
func SetOndragexit(val dom.EventHandler) { value.SetOndragexit(val) }
func GetOndragleave() dom.EventHandler {
	return value.GetOndragleave()
}
func SetOndragleave(val dom.EventHandler) { value.SetOndragleave(val) }
func GetOndragover() dom.EventHandler {
	return value.GetOndragover()
}
func SetOndragover(val dom.EventHandler) { value.SetOndragover(val) }
func GetOndragstart() dom.EventHandler {
	return value.GetOndragstart()
}
func SetOndragstart(val dom.EventHandler) { value.SetOndragstart(val) }
func GetOndrop() dom.EventHandler {
	return value.GetOndrop()
}
func SetOndrop(val dom.EventHandler) { value.SetOndrop(val) }
func GetOndurationchange() dom.EventHandler {
	return value.GetOndurationchange()
}
func SetOndurationchange(val dom.EventHandler) { value.SetOndurationchange(val) }
func GetOnemptied() dom.EventHandler {
	return value.GetOnemptied()
}
func SetOnemptied(val dom.EventHandler) { value.SetOnemptied(val) }
func GetOnended() dom.EventHandler {
	return value.GetOnended()
}
func SetOnended(val dom.EventHandler) { value.SetOnended(val) }
func GetOnerror() dom.OnErrorEventHandler {
	return value.GetOnerror()
}
func SetOnerror(val dom.OnErrorEventHandler) { value.SetOnerror(val) }
func GetOnfocus() dom.EventHandler {
	return value.GetOnfocus()
}
func SetOnfocus(val dom.EventHandler) { value.SetOnfocus(val) }
func GetOninput() dom.EventHandler {
	return value.GetOninput()
}
func SetOninput(val dom.EventHandler) { value.SetOninput(val) }
func GetOninvalid() dom.EventHandler {
	return value.GetOninvalid()
}
func SetOninvalid(val dom.EventHandler) { value.SetOninvalid(val) }
func GetOnkeydown() dom.EventHandler {
	return value.GetOnkeydown()
}
func SetOnkeydown(val dom.EventHandler) { value.SetOnkeydown(val) }
func GetOnkeypress() dom.EventHandler {
	return value.GetOnkeypress()
}
func SetOnkeypress(val dom.EventHandler) { value.SetOnkeypress(val) }
func GetOnkeyup() dom.EventHandler {
	return value.GetOnkeyup()
}
func SetOnkeyup(val dom.EventHandler) { value.SetOnkeyup(val) }
func GetOnload() dom.EventHandler {
	return value.GetOnload()
}
func SetOnload(val dom.EventHandler) { value.SetOnload(val) }
func GetOnloadeddata() dom.EventHandler {
	return value.GetOnloadeddata()
}
func SetOnloadeddata(val dom.EventHandler) { value.SetOnloadeddata(val) }
func GetOnloadedmetadata() dom.EventHandler {
	return value.GetOnloadedmetadata()
}
func SetOnloadedmetadata(val dom.EventHandler) { value.SetOnloadedmetadata(val) }
func GetOnloadend() dom.EventHandler {
	return value.GetOnloadend()
}
func SetOnloadend(val dom.EventHandler) { value.SetOnloadend(val) }
func GetOnloadstart() dom.EventHandler {
	return value.GetOnloadstart()
}
func SetOnloadstart(val dom.EventHandler) { value.SetOnloadstart(val) }
func GetOnmousedown() dom.EventHandler {
	return value.GetOnmousedown()
}
func SetOnmousedown(val dom.EventHandler) { value.SetOnmousedown(val) }
func GetOnmouseenter() dom.EventHandler {
	return value.GetOnmouseenter()
}
func SetOnmouseenter(val dom.EventHandler) { value.SetOnmouseenter(val) }
func GetOnmouseleave() dom.EventHandler {
	return value.GetOnmouseleave()
}
func SetOnmouseleave(val dom.EventHandler) { value.SetOnmouseleave(val) }
func GetOnmousemove() dom.EventHandler {
	return value.GetOnmousemove()
}
func SetOnmousemove(val dom.EventHandler) { value.SetOnmousemove(val) }
func GetOnmouseout() dom.EventHandler {
	return value.GetOnmouseout()
}
func SetOnmouseout(val dom.EventHandler) { value.SetOnmouseout(val) }
func GetOnmouseover() dom.EventHandler {
	return value.GetOnmouseover()
}
func SetOnmouseover(val dom.EventHandler) { value.SetOnmouseover(val) }
func GetOnmouseup() dom.EventHandler {
	return value.GetOnmouseup()
}
func SetOnmouseup(val dom.EventHandler) { value.SetOnmouseup(val) }
func GetOnpaste() dom.EventHandler {
	return value.GetOnpaste()
}
func SetOnpaste(val dom.EventHandler) { value.SetOnpaste(val) }
func GetOnpause() dom.EventHandler {
	return value.GetOnpause()
}
func SetOnpause(val dom.EventHandler) { value.SetOnpause(val) }
func GetOnplay() dom.EventHandler {
	return value.GetOnplay()
}
func SetOnplay(val dom.EventHandler) { value.SetOnplay(val) }
func GetOnplaying() dom.EventHandler {
	return value.GetOnplaying()
}
func SetOnplaying(val dom.EventHandler) { value.SetOnplaying(val) }
func GetOnprogress() dom.EventHandler {
	return value.GetOnprogress()
}
func SetOnprogress(val dom.EventHandler) { value.SetOnprogress(val) }
func GetOnratechange() dom.EventHandler {
	return value.GetOnratechange()
}
func SetOnratechange(val dom.EventHandler) { value.SetOnratechange(val) }
func GetOnreset() dom.EventHandler {
	return value.GetOnreset()
}
func SetOnreset(val dom.EventHandler) { value.SetOnreset(val) }
func GetOnresize() dom.EventHandler {
	return value.GetOnresize()
}
func SetOnresize(val dom.EventHandler) { value.SetOnresize(val) }
func GetOnscroll() dom.EventHandler {
	return value.GetOnscroll()
}
func SetOnscroll(val dom.EventHandler) { value.SetOnscroll(val) }
func GetOnsecuritypolicyviolation() dom.EventHandler {
	return value.GetOnsecuritypolicyviolation()
}
func SetOnsecuritypolicyviolation(val dom.EventHandler) { value.SetOnsecuritypolicyviolation(val) }
func GetOnseeked() dom.EventHandler {
	return value.GetOnseeked()
}
func SetOnseeked(val dom.EventHandler) { value.SetOnseeked(val) }
func GetOnseeking() dom.EventHandler {
	return value.GetOnseeking()
}
func SetOnseeking(val dom.EventHandler) { value.SetOnseeking(val) }
func GetOnselect() dom.EventHandler {
	return value.GetOnselect()
}
func SetOnselect(val dom.EventHandler) { value.SetOnselect(val) }
func GetOnstalled() dom.EventHandler {
	return value.GetOnstalled()
}
func SetOnstalled(val dom.EventHandler) { value.SetOnstalled(val) }
func GetOnsubmit() dom.EventHandler {
	return value.GetOnsubmit()
}
func SetOnsubmit(val dom.EventHandler) { value.SetOnsubmit(val) }
func GetOnsuspend() dom.EventHandler {
	return value.GetOnsuspend()
}
func SetOnsuspend(val dom.EventHandler) { value.SetOnsuspend(val) }
func GetOntimeupdate() dom.EventHandler {
	return value.GetOntimeupdate()
}
func SetOntimeupdate(val dom.EventHandler) { value.SetOntimeupdate(val) }
func GetOntoggle() dom.EventHandler {
	return value.GetOntoggle()
}
func SetOntoggle(val dom.EventHandler) { value.SetOntoggle(val) }
func GetOnvolumechange() dom.EventHandler {
	return value.GetOnvolumechange()
}
func SetOnvolumechange(val dom.EventHandler) { value.SetOnvolumechange(val) }
func GetOnwaiting() dom.EventHandler {
	return value.GetOnwaiting()
}
func SetOnwaiting(val dom.EventHandler) { value.SetOnwaiting(val) }
func GetOnwheel() dom.EventHandler {
	return value.GetOnwheel()
}
func SetOnwheel(val dom.EventHandler) { value.SetOnwheel(val) }
func GetOrigin() string {
	return value.GetOrigin()
}
func GetOwnerDocument() dom.Document {
	return value.GetOwnerDocument()
}
func GetParentElement() dom.Element {
	return value.GetParentElement()
}
func GetParentNode() dom.Node {
	return value.GetParentNode()
}
func Prepend(args ...interface{}) {
	value.Call("prepend", args...)
}
func GetPreviousSibling() dom.Node {
	return value.GetPreviousSibling()
}
func QuerySelector(args ...interface{}) dom.Element {
	val := value.Call("querySelector", args...)
	return dom.JSValueToElement(val.JSValue())
}
func QuerySelectorAll(args ...interface{}) dom.NodeList {
	val := value.Call("querySelectorAll", args...)
	return dom.JSValueToNodeList(val.JSValue())
}
func RemoveChild(args ...interface{}) dom.Node {
	val := value.Call("removeChild", args...)
	return dom.JSValueToNode(val.JSValue())
}
func RemoveEventListener(args ...interface{}) {
	value.Call("removeEventListener", args...)
}
func ReplaceChild(args ...interface{}) dom.Node {
	val := value.Call("replaceChild", args...)
	return dom.JSValueToNode(val.JSValue())
}
func GetTextContent() string {
	return value.GetTextContent()
}
func SetTextContent(val string) { value.SetTextContent(val) }
func GetURL() string {
	return value.GetURL()
}
