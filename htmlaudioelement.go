// Code generated DO NOT EDIT
// htmlaudioelement.go
package dom

import "syscall/js"

type HTMLAudioElementIFace interface {
	GetAccessKey() string
	SetAccessKey(string)
	GetAccessKeyLabel() string
	AddEventListener(args ...interface{})
	AddTextTrack(args ...interface{}) TextTrack
	AppendChild(args ...interface{}) Node
	AttachShadow(args ...interface{}) ShadowRoot
	GetAttributes() NamedNodeMap
	GetAudioTracks() AudioTrackList
	GetAutocapitalize() string
	SetAutocapitalize(string)
	GetAutoplay() bool
	SetAutoplay(bool)
	GetBaseURI() string
	GetBuffered() TimeRanges
	CanPlayType(args ...interface{}) CanPlayTypeResult
	GetChildNodes() NodeList
	GetClassList() DOMTokenList
	GetClassName() string
	SetClassName(string)
	Click(args ...interface{})
	CloneNode(args ...interface{}) Node
	Closest(args ...interface{}) Element
	CompareDocumentPosition(args ...interface{}) int
	Contains(args ...interface{}) bool
	GetControls() bool
	SetControls(bool)
	GetCrossOrigin() string
	SetCrossOrigin(string)
	GetCurrentSrc() string
	GetCurrentTime() float64
	SetCurrentTime(float64)
	GetDefaultMuted() bool
	SetDefaultMuted(bool)
	GetDefaultPlaybackRate() float64
	SetDefaultPlaybackRate(float64)
	GetDir() string
	SetDir(string)
	DispatchEvent(args ...interface{}) bool
	GetDraggable() bool
	SetDraggable(bool)
	GetDuration() float64
	GetEnded() bool
	GetError() MediaError
	FastSeek(args ...interface{})
	GetFirstChild() Node
	GetAttribute(args ...interface{}) string
	GetAttributeNS(args ...interface{}) string
	GetAttributeNames(args ...interface{})
	GetAttributeNode(args ...interface{}) Attr
	GetAttributeNodeNS(args ...interface{}) Attr
	GetElementsByClassName(args ...interface{}) HTMLCollection
	GetElementsByTagName(args ...interface{}) HTMLCollection
	GetElementsByTagNameNS(args ...interface{}) HTMLCollection
	GetRootNode(args ...interface{}) Node
	GetStartDate(args ...interface{}) Value
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
	InsertAdjacentElement(args ...interface{}) Element
	InsertAdjacentText(args ...interface{})
	InsertBefore(args ...interface{}) Node
	GetIsConnected() bool
	IsDefaultNamespace(args ...interface{}) bool
	IsEqualNode(args ...interface{}) bool
	IsSameNode(args ...interface{}) bool
	GetLang() string
	SetLang(string)
	GetLastChild() Node
	Load(args ...interface{})
	GetLocalName() string
	LookupNamespaceURI(args ...interface{}) string
	LookupPrefix(args ...interface{}) string
	GetLoop() bool
	SetLoop(bool)
	Matches(args ...interface{}) bool
	GetMuted() bool
	SetMuted(bool)
	GetNamespaceURI() string
	GetNetworkState() int
	GetNextSibling() Node
	GetNodeName() string
	GetNodeType() int
	GetNodeValue() string
	SetNodeValue(string)
	Normalize(args ...interface{})
	GetOwnerDocument() Document
	GetParentElement() Element
	GetParentNode() Node
	Pause(args ...interface{})
	GetPaused() bool
	Play(args ...interface{})
	GetPlaybackRate() float64
	SetPlaybackRate(float64)
	GetPlayed() TimeRanges
	GetPrefix() string
	GetPreload() string
	SetPreload(string)
	GetPreviousSibling() Node
	GetReadyState() int
	RemoveAttribute(args ...interface{})
	RemoveAttributeNS(args ...interface{})
	RemoveAttributeNode(args ...interface{}) Attr
	RemoveChild(args ...interface{}) Node
	RemoveEventListener(args ...interface{})
	ReplaceChild(args ...interface{}) Node
	GetSeekable() TimeRanges
	GetSeeking() bool
	SetAttribute(args ...interface{})
	SetAttributeNS(args ...interface{})
	SetAttributeNode(args ...interface{}) Attr
	SetAttributeNodeNS(args ...interface{}) Attr
	GetShadowRoot() ShadowRoot
	GetSlot() string
	SetSlot(string)
	GetSpellcheck() bool
	SetSpellcheck(bool)
	GetSrc() string
	SetSrc(string)
	GetSrcObject() MediaProvider
	SetSrcObject(MediaProvider)
	GetTagName() string
	GetTextContent() string
	SetTextContent(string)
	GetTextTracks() TextTrackList
	GetTitle() string
	SetTitle(string)
	ToggleAttribute(args ...interface{}) bool
	GetTranslate() bool
	SetTranslate(bool)
	GetVideoTracks() VideoTrackList
	GetVolume() float64
	SetVolume(float64)
	WebkitMatchesSelector(args ...interface{}) bool
}
type HTMLAudioElement struct {
	Value
}

func JSValueToHTMLAudioElement(val js.Value) HTMLAudioElement {
	return HTMLAudioElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLAudioElement() HTMLAudioElement { return HTMLAudioElement{Value: v} }
func NewHTMLAudioElement(args ...interface{}) HTMLAudioElement {
	return HTMLAudioElement{Value: JSValueToValue(js.Global().Get("HTMLAudioElement").New(args...))}
}
func (h HTMLAudioElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLAudioElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLAudioElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLAudioElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLAudioElement) AddTextTrack(args ...interface{}) TextTrack {
	val := h.Call("addTextTrack", args...)
	return JSValueToTextTrack(val.JSValue())
}
func (h HTMLAudioElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLAudioElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLAudioElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLAudioElement) GetAudioTracks() AudioTrackList {
	val := h.Get("audioTracks")
	return JSValueToAudioTrackList(val.JSValue())
}
func (h HTMLAudioElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLAudioElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLAudioElement) GetAutoplay() bool {
	val := h.Get("autoplay")
	return val.Bool()
}
func (h HTMLAudioElement) SetAutoplay(val bool) {
	h.Set("autoplay", val)
}
func (h HTMLAudioElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLAudioElement) GetBuffered() TimeRanges {
	val := h.Get("buffered")
	return JSValueToTimeRanges(val.JSValue())
}
func (h HTMLAudioElement) CanPlayType(args ...interface{}) CanPlayTypeResult {
	val := h.Call("canPlayType", args...)
	return JSValueToCanPlayTypeResult(val.JSValue())
}
func (h HTMLAudioElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLAudioElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLAudioElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLAudioElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLAudioElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLAudioElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLAudioElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLAudioElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLAudioElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLAudioElement) GetControls() bool {
	val := h.Get("controls")
	return val.Bool()
}
func (h HTMLAudioElement) SetControls(val bool) {
	h.Set("controls", val)
}
func (h HTMLAudioElement) GetCrossOrigin() string {
	val := h.Get("crossOrigin")
	return val.String()
}
func (h HTMLAudioElement) SetCrossOrigin(val string) {
	h.Set("crossOrigin", val)
}
func (h HTMLAudioElement) GetCurrentSrc() string {
	val := h.Get("currentSrc")
	return val.String()
}
func (h HTMLAudioElement) GetCurrentTime() float64 {
	val := h.Get("currentTime")
	return val.Float()
}
func (h HTMLAudioElement) SetCurrentTime(val float64) {
	h.Set("currentTime", val)
}
func (h HTMLAudioElement) GetDefaultMuted() bool {
	val := h.Get("defaultMuted")
	return val.Bool()
}
func (h HTMLAudioElement) SetDefaultMuted(val bool) {
	h.Set("defaultMuted", val)
}
func (h HTMLAudioElement) GetDefaultPlaybackRate() float64 {
	val := h.Get("defaultPlaybackRate")
	return val.Float()
}
func (h HTMLAudioElement) SetDefaultPlaybackRate(val float64) {
	h.Set("defaultPlaybackRate", val)
}
func (h HTMLAudioElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLAudioElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLAudioElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLAudioElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLAudioElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLAudioElement) GetDuration() float64 {
	val := h.Get("duration")
	return val.Float()
}
func (h HTMLAudioElement) GetEnded() bool {
	val := h.Get("ended")
	return val.Bool()
}
func (h HTMLAudioElement) GetError() MediaError {
	val := h.Get("error")
	return JSValueToMediaError(val.JSValue())
}
func (h HTMLAudioElement) FastSeek(args ...interface{}) {
	h.Call("fastSeek", args...)
}
func (h HTMLAudioElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLAudioElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLAudioElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLAudioElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLAudioElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLAudioElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLAudioElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLAudioElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLAudioElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLAudioElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLAudioElement) GetStartDate(args ...interface{}) Value {
	val := h.Call("getStartDate", args...)
	return val
}
func (h HTMLAudioElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLAudioElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLAudioElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLAudioElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLAudioElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLAudioElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLAudioElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLAudioElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLAudioElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLAudioElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLAudioElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLAudioElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLAudioElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLAudioElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLAudioElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLAudioElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLAudioElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLAudioElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLAudioElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLAudioElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLAudioElement) Load(args ...interface{}) {
	h.Call("load", args...)
}
func (h HTMLAudioElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLAudioElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLAudioElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLAudioElement) GetLoop() bool {
	val := h.Get("loop")
	return val.Bool()
}
func (h HTMLAudioElement) SetLoop(val bool) {
	h.Set("loop", val)
}
func (h HTMLAudioElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLAudioElement) GetMuted() bool {
	val := h.Get("muted")
	return val.Bool()
}
func (h HTMLAudioElement) SetMuted(val bool) {
	h.Set("muted", val)
}
func (h HTMLAudioElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLAudioElement) GetNetworkState() int {
	val := h.Get("networkState")
	return val.Int()
}
func (h HTMLAudioElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLAudioElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLAudioElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLAudioElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLAudioElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLAudioElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLAudioElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLAudioElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLAudioElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLAudioElement) Pause(args ...interface{}) {
	h.Call("pause", args...)
}
func (h HTMLAudioElement) GetPaused() bool {
	val := h.Get("paused")
	return val.Bool()
}
func (h HTMLAudioElement) Play(args ...interface{}) {
	h.Call("play", args...)
}
func (h HTMLAudioElement) GetPlaybackRate() float64 {
	val := h.Get("playbackRate")
	return val.Float()
}
func (h HTMLAudioElement) SetPlaybackRate(val float64) {
	h.Set("playbackRate", val)
}
func (h HTMLAudioElement) GetPlayed() TimeRanges {
	val := h.Get("played")
	return JSValueToTimeRanges(val.JSValue())
}
func (h HTMLAudioElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLAudioElement) GetPreload() string {
	val := h.Get("preload")
	return val.String()
}
func (h HTMLAudioElement) SetPreload(val string) {
	h.Set("preload", val)
}
func (h HTMLAudioElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLAudioElement) GetReadyState() int {
	val := h.Get("readyState")
	return val.Int()
}
func (h HTMLAudioElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLAudioElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLAudioElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLAudioElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLAudioElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLAudioElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLAudioElement) GetSeekable() TimeRanges {
	val := h.Get("seekable")
	return JSValueToTimeRanges(val.JSValue())
}
func (h HTMLAudioElement) GetSeeking() bool {
	val := h.Get("seeking")
	return val.Bool()
}
func (h HTMLAudioElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLAudioElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLAudioElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLAudioElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLAudioElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLAudioElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLAudioElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLAudioElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLAudioElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLAudioElement) GetSrc() string {
	val := h.Get("src")
	return val.String()
}
func (h HTMLAudioElement) SetSrc(val string) {
	h.Set("src", val)
}
func (h HTMLAudioElement) GetSrcObject() MediaProvider {
	val := h.Get("srcObject")
	return JSValueToMediaProvider(val.JSValue())
}
func (h HTMLAudioElement) SetSrcObject(val MediaProvider) {
	h.Set("srcObject", val)
}
func (h HTMLAudioElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLAudioElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLAudioElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLAudioElement) GetTextTracks() TextTrackList {
	val := h.Get("textTracks")
	return JSValueToTextTrackList(val.JSValue())
}
func (h HTMLAudioElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLAudioElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLAudioElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLAudioElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLAudioElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLAudioElement) GetVideoTracks() VideoTrackList {
	val := h.Get("videoTracks")
	return JSValueToVideoTrackList(val.JSValue())
}
func (h HTMLAudioElement) GetVolume() float64 {
	val := h.Get("volume")
	return val.Float()
}
func (h HTMLAudioElement) SetVolume(val float64) {
	h.Set("volume", val)
}
func (h HTMLAudioElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
