// Code generated DO NOT EDIT
// htmlmediaelement.go
package dom

import "syscall/js"

type HTMLMediaElementIFace interface {
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
type HTMLMediaElement struct {
	Value
}

func JSValueToHTMLMediaElement(val js.Value) HTMLMediaElement {
	return HTMLMediaElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLMediaElement() HTMLMediaElement { return HTMLMediaElement{Value: v} }
func NewHTMLMediaElement(args ...interface{}) HTMLMediaElement {
	return HTMLMediaElement{Value: JSValueToValue(js.Global().Get("HTMLMediaElement").New(args...))}
}
func (h HTMLMediaElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLMediaElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLMediaElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLMediaElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLMediaElement) AddTextTrack(args ...interface{}) TextTrack {
	val := h.Call("addTextTrack", args...)
	return JSValueToTextTrack(val.JSValue())
}
func (h HTMLMediaElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLMediaElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLMediaElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLMediaElement) GetAudioTracks() AudioTrackList {
	val := h.Get("audioTracks")
	return JSValueToAudioTrackList(val.JSValue())
}
func (h HTMLMediaElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLMediaElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLMediaElement) GetAutoplay() bool {
	val := h.Get("autoplay")
	return val.Bool()
}
func (h HTMLMediaElement) SetAutoplay(val bool) {
	h.Set("autoplay", val)
}
func (h HTMLMediaElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLMediaElement) GetBuffered() TimeRanges {
	val := h.Get("buffered")
	return JSValueToTimeRanges(val.JSValue())
}
func (h HTMLMediaElement) CanPlayType(args ...interface{}) CanPlayTypeResult {
	val := h.Call("canPlayType", args...)
	return JSValueToCanPlayTypeResult(val.JSValue())
}
func (h HTMLMediaElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLMediaElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLMediaElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLMediaElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLMediaElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLMediaElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLMediaElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLMediaElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLMediaElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLMediaElement) GetControls() bool {
	val := h.Get("controls")
	return val.Bool()
}
func (h HTMLMediaElement) SetControls(val bool) {
	h.Set("controls", val)
}
func (h HTMLMediaElement) GetCrossOrigin() string {
	val := h.Get("crossOrigin")
	return val.String()
}
func (h HTMLMediaElement) SetCrossOrigin(val string) {
	h.Set("crossOrigin", val)
}
func (h HTMLMediaElement) GetCurrentSrc() string {
	val := h.Get("currentSrc")
	return val.String()
}
func (h HTMLMediaElement) GetCurrentTime() float64 {
	val := h.Get("currentTime")
	return val.Float()
}
func (h HTMLMediaElement) SetCurrentTime(val float64) {
	h.Set("currentTime", val)
}
func (h HTMLMediaElement) GetDefaultMuted() bool {
	val := h.Get("defaultMuted")
	return val.Bool()
}
func (h HTMLMediaElement) SetDefaultMuted(val bool) {
	h.Set("defaultMuted", val)
}
func (h HTMLMediaElement) GetDefaultPlaybackRate() float64 {
	val := h.Get("defaultPlaybackRate")
	return val.Float()
}
func (h HTMLMediaElement) SetDefaultPlaybackRate(val float64) {
	h.Set("defaultPlaybackRate", val)
}
func (h HTMLMediaElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLMediaElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLMediaElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLMediaElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLMediaElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLMediaElement) GetDuration() float64 {
	val := h.Get("duration")
	return val.Float()
}
func (h HTMLMediaElement) GetEnded() bool {
	val := h.Get("ended")
	return val.Bool()
}
func (h HTMLMediaElement) GetError() MediaError {
	val := h.Get("error")
	return JSValueToMediaError(val.JSValue())
}
func (h HTMLMediaElement) FastSeek(args ...interface{}) {
	h.Call("fastSeek", args...)
}
func (h HTMLMediaElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLMediaElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLMediaElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLMediaElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLMediaElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLMediaElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLMediaElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLMediaElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLMediaElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLMediaElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLMediaElement) GetStartDate(args ...interface{}) Value {
	val := h.Call("getStartDate", args...)
	return val
}
func (h HTMLMediaElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLMediaElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLMediaElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLMediaElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLMediaElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLMediaElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLMediaElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLMediaElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLMediaElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLMediaElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLMediaElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLMediaElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLMediaElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLMediaElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLMediaElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLMediaElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLMediaElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLMediaElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLMediaElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLMediaElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLMediaElement) Load(args ...interface{}) {
	h.Call("load", args...)
}
func (h HTMLMediaElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLMediaElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLMediaElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLMediaElement) GetLoop() bool {
	val := h.Get("loop")
	return val.Bool()
}
func (h HTMLMediaElement) SetLoop(val bool) {
	h.Set("loop", val)
}
func (h HTMLMediaElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLMediaElement) GetMuted() bool {
	val := h.Get("muted")
	return val.Bool()
}
func (h HTMLMediaElement) SetMuted(val bool) {
	h.Set("muted", val)
}
func (h HTMLMediaElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLMediaElement) GetNetworkState() int {
	val := h.Get("networkState")
	return val.Int()
}
func (h HTMLMediaElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLMediaElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLMediaElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLMediaElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLMediaElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLMediaElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLMediaElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLMediaElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLMediaElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLMediaElement) Pause(args ...interface{}) {
	h.Call("pause", args...)
}
func (h HTMLMediaElement) GetPaused() bool {
	val := h.Get("paused")
	return val.Bool()
}
func (h HTMLMediaElement) Play(args ...interface{}) {
	h.Call("play", args...)
}
func (h HTMLMediaElement) GetPlaybackRate() float64 {
	val := h.Get("playbackRate")
	return val.Float()
}
func (h HTMLMediaElement) SetPlaybackRate(val float64) {
	h.Set("playbackRate", val)
}
func (h HTMLMediaElement) GetPlayed() TimeRanges {
	val := h.Get("played")
	return JSValueToTimeRanges(val.JSValue())
}
func (h HTMLMediaElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLMediaElement) GetPreload() string {
	val := h.Get("preload")
	return val.String()
}
func (h HTMLMediaElement) SetPreload(val string) {
	h.Set("preload", val)
}
func (h HTMLMediaElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLMediaElement) GetReadyState() int {
	val := h.Get("readyState")
	return val.Int()
}
func (h HTMLMediaElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLMediaElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLMediaElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLMediaElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLMediaElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLMediaElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLMediaElement) GetSeekable() TimeRanges {
	val := h.Get("seekable")
	return JSValueToTimeRanges(val.JSValue())
}
func (h HTMLMediaElement) GetSeeking() bool {
	val := h.Get("seeking")
	return val.Bool()
}
func (h HTMLMediaElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLMediaElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLMediaElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLMediaElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLMediaElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLMediaElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLMediaElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLMediaElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLMediaElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLMediaElement) GetSrc() string {
	val := h.Get("src")
	return val.String()
}
func (h HTMLMediaElement) SetSrc(val string) {
	h.Set("src", val)
}
func (h HTMLMediaElement) GetSrcObject() MediaProvider {
	val := h.Get("srcObject")
	return JSValueToMediaProvider(val.JSValue())
}
func (h HTMLMediaElement) SetSrcObject(val MediaProvider) {
	h.Set("srcObject", val)
}
func (h HTMLMediaElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLMediaElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLMediaElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLMediaElement) GetTextTracks() TextTrackList {
	val := h.Get("textTracks")
	return JSValueToTextTrackList(val.JSValue())
}
func (h HTMLMediaElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLMediaElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLMediaElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLMediaElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLMediaElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLMediaElement) GetVideoTracks() VideoTrackList {
	val := h.Get("videoTracks")
	return JSValueToVideoTrackList(val.JSValue())
}
func (h HTMLMediaElement) GetVolume() float64 {
	val := h.Get("volume")
	return val.Float()
}
func (h HTMLMediaElement) SetVolume(val float64) {
	h.Set("volume", val)
}
func (h HTMLMediaElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
