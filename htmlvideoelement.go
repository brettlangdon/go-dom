// Code generated DO NOT EDIT
// htmlvideoelement.go
package dom

import "syscall/js"

type HTMLVideoElementIFace interface {
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
	GetHeight() int
	SetHeight(int)
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
	GetPlaysInline() bool
	SetPlaysInline(bool)
	GetPoster() string
	SetPoster(string)
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
	GetVideoHeight() int
	GetVideoTracks() VideoTrackList
	GetVideoWidth() int
	GetVolume() float64
	SetVolume(float64)
	WebkitMatchesSelector(args ...interface{}) bool
	GetWidth() int
	SetWidth(int)
}
type HTMLVideoElement struct {
	Value
}

func JSValueToHTMLVideoElement(val js.Value) HTMLVideoElement {
	return HTMLVideoElement{Value: JSValueToValue(val)}
}
func (v Value) AsHTMLVideoElement() HTMLVideoElement { return HTMLVideoElement{Value: v} }
func NewHTMLVideoElement(args ...interface{}) HTMLVideoElement {
	return HTMLVideoElement{Value: JSValueToValue(js.Global().Get("HTMLVideoElement").New(args...))}
}
func (h HTMLVideoElement) GetAccessKey() string {
	val := h.Get("accessKey")
	return val.String()
}
func (h HTMLVideoElement) SetAccessKey(val string) {
	h.Set("accessKey", val)
}
func (h HTMLVideoElement) GetAccessKeyLabel() string {
	val := h.Get("accessKeyLabel")
	return val.String()
}
func (h HTMLVideoElement) AddEventListener(args ...interface{}) {
	h.Call("addEventListener", args...)
}
func (h HTMLVideoElement) AddTextTrack(args ...interface{}) TextTrack {
	val := h.Call("addTextTrack", args...)
	return JSValueToTextTrack(val.JSValue())
}
func (h HTMLVideoElement) AppendChild(args ...interface{}) Node {
	val := h.Call("appendChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLVideoElement) AttachShadow(args ...interface{}) ShadowRoot {
	val := h.Call("attachShadow", args...)
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLVideoElement) GetAttributes() NamedNodeMap {
	val := h.Get("attributes")
	return JSValueToNamedNodeMap(val.JSValue())
}
func (h HTMLVideoElement) GetAudioTracks() AudioTrackList {
	val := h.Get("audioTracks")
	return JSValueToAudioTrackList(val.JSValue())
}
func (h HTMLVideoElement) GetAutocapitalize() string {
	val := h.Get("autocapitalize")
	return val.String()
}
func (h HTMLVideoElement) SetAutocapitalize(val string) {
	h.Set("autocapitalize", val)
}
func (h HTMLVideoElement) GetAutoplay() bool {
	val := h.Get("autoplay")
	return val.Bool()
}
func (h HTMLVideoElement) SetAutoplay(val bool) {
	h.Set("autoplay", val)
}
func (h HTMLVideoElement) GetBaseURI() string {
	val := h.Get("baseURI")
	return val.String()
}
func (h HTMLVideoElement) GetBuffered() TimeRanges {
	val := h.Get("buffered")
	return JSValueToTimeRanges(val.JSValue())
}
func (h HTMLVideoElement) CanPlayType(args ...interface{}) CanPlayTypeResult {
	val := h.Call("canPlayType", args...)
	return JSValueToCanPlayTypeResult(val.JSValue())
}
func (h HTMLVideoElement) GetChildNodes() NodeList {
	val := h.Get("childNodes")
	return JSValueToNodeList(val.JSValue())
}
func (h HTMLVideoElement) GetClassList() DOMTokenList {
	val := h.Get("classList")
	return JSValueToDOMTokenList(val.JSValue())
}
func (h HTMLVideoElement) GetClassName() string {
	val := h.Get("className")
	return val.String()
}
func (h HTMLVideoElement) SetClassName(val string) {
	h.Set("className", val)
}
func (h HTMLVideoElement) Click(args ...interface{}) {
	h.Call("click", args...)
}
func (h HTMLVideoElement) CloneNode(args ...interface{}) Node {
	val := h.Call("cloneNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLVideoElement) Closest(args ...interface{}) Element {
	val := h.Call("closest", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLVideoElement) CompareDocumentPosition(args ...interface{}) int {
	val := h.Call("compareDocumentPosition", args...)
	return val.Int()
}
func (h HTMLVideoElement) Contains(args ...interface{}) bool {
	val := h.Call("contains", args...)
	return val.Bool()
}
func (h HTMLVideoElement) GetControls() bool {
	val := h.Get("controls")
	return val.Bool()
}
func (h HTMLVideoElement) SetControls(val bool) {
	h.Set("controls", val)
}
func (h HTMLVideoElement) GetCrossOrigin() string {
	val := h.Get("crossOrigin")
	return val.String()
}
func (h HTMLVideoElement) SetCrossOrigin(val string) {
	h.Set("crossOrigin", val)
}
func (h HTMLVideoElement) GetCurrentSrc() string {
	val := h.Get("currentSrc")
	return val.String()
}
func (h HTMLVideoElement) GetCurrentTime() float64 {
	val := h.Get("currentTime")
	return val.Float()
}
func (h HTMLVideoElement) SetCurrentTime(val float64) {
	h.Set("currentTime", val)
}
func (h HTMLVideoElement) GetDefaultMuted() bool {
	val := h.Get("defaultMuted")
	return val.Bool()
}
func (h HTMLVideoElement) SetDefaultMuted(val bool) {
	h.Set("defaultMuted", val)
}
func (h HTMLVideoElement) GetDefaultPlaybackRate() float64 {
	val := h.Get("defaultPlaybackRate")
	return val.Float()
}
func (h HTMLVideoElement) SetDefaultPlaybackRate(val float64) {
	h.Set("defaultPlaybackRate", val)
}
func (h HTMLVideoElement) GetDir() string {
	val := h.Get("dir")
	return val.String()
}
func (h HTMLVideoElement) SetDir(val string) {
	h.Set("dir", val)
}
func (h HTMLVideoElement) DispatchEvent(args ...interface{}) bool {
	val := h.Call("dispatchEvent", args...)
	return val.Bool()
}
func (h HTMLVideoElement) GetDraggable() bool {
	val := h.Get("draggable")
	return val.Bool()
}
func (h HTMLVideoElement) SetDraggable(val bool) {
	h.Set("draggable", val)
}
func (h HTMLVideoElement) GetDuration() float64 {
	val := h.Get("duration")
	return val.Float()
}
func (h HTMLVideoElement) GetEnded() bool {
	val := h.Get("ended")
	return val.Bool()
}
func (h HTMLVideoElement) GetError() MediaError {
	val := h.Get("error")
	return JSValueToMediaError(val.JSValue())
}
func (h HTMLVideoElement) FastSeek(args ...interface{}) {
	h.Call("fastSeek", args...)
}
func (h HTMLVideoElement) GetFirstChild() Node {
	val := h.Get("firstChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLVideoElement) GetAttribute(args ...interface{}) string {
	val := h.Call("getAttribute", args...)
	return val.String()
}
func (h HTMLVideoElement) GetAttributeNS(args ...interface{}) string {
	val := h.Call("getAttributeNS", args...)
	return val.String()
}
func (h HTMLVideoElement) GetAttributeNames(args ...interface{}) {
	h.Call("getAttributeNames", args...)
}
func (h HTMLVideoElement) GetAttributeNode(args ...interface{}) Attr {
	val := h.Call("getAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLVideoElement) GetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("getAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLVideoElement) GetElementsByClassName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByClassName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLVideoElement) GetElementsByTagName(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagName", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLVideoElement) GetElementsByTagNameNS(args ...interface{}) HTMLCollection {
	val := h.Call("getElementsByTagNameNS", args...)
	return JSValueToHTMLCollection(val.JSValue())
}
func (h HTMLVideoElement) GetRootNode(args ...interface{}) Node {
	val := h.Call("getRootNode", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLVideoElement) GetStartDate(args ...interface{}) Value {
	val := h.Call("getStartDate", args...)
	return val
}
func (h HTMLVideoElement) HasAttribute(args ...interface{}) bool {
	val := h.Call("hasAttribute", args...)
	return val.Bool()
}
func (h HTMLVideoElement) HasAttributeNS(args ...interface{}) bool {
	val := h.Call("hasAttributeNS", args...)
	return val.Bool()
}
func (h HTMLVideoElement) HasAttributes(args ...interface{}) bool {
	val := h.Call("hasAttributes", args...)
	return val.Bool()
}
func (h HTMLVideoElement) HasChildNodes(args ...interface{}) bool {
	val := h.Call("hasChildNodes", args...)
	return val.Bool()
}
func (h HTMLVideoElement) GetHeight() int {
	val := h.Get("height")
	return val.Int()
}
func (h HTMLVideoElement) SetHeight(val int) {
	h.Set("height", val)
}
func (h HTMLVideoElement) GetHidden() bool {
	val := h.Get("hidden")
	return val.Bool()
}
func (h HTMLVideoElement) SetHidden(val bool) {
	h.Set("hidden", val)
}
func (h HTMLVideoElement) GetId() string {
	val := h.Get("id")
	return val.String()
}
func (h HTMLVideoElement) SetId(val string) {
	h.Set("id", val)
}
func (h HTMLVideoElement) GetInnerText() string {
	val := h.Get("innerText")
	return val.String()
}
func (h HTMLVideoElement) SetInnerText(val string) {
	h.Set("innerText", val)
}
func (h HTMLVideoElement) InsertAdjacentElement(args ...interface{}) Element {
	val := h.Call("insertAdjacentElement", args...)
	return JSValueToElement(val.JSValue())
}
func (h HTMLVideoElement) InsertAdjacentText(args ...interface{}) {
	h.Call("insertAdjacentText", args...)
}
func (h HTMLVideoElement) InsertBefore(args ...interface{}) Node {
	val := h.Call("insertBefore", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLVideoElement) GetIsConnected() bool {
	val := h.Get("isConnected")
	return val.Bool()
}
func (h HTMLVideoElement) IsDefaultNamespace(args ...interface{}) bool {
	val := h.Call("isDefaultNamespace", args...)
	return val.Bool()
}
func (h HTMLVideoElement) IsEqualNode(args ...interface{}) bool {
	val := h.Call("isEqualNode", args...)
	return val.Bool()
}
func (h HTMLVideoElement) IsSameNode(args ...interface{}) bool {
	val := h.Call("isSameNode", args...)
	return val.Bool()
}
func (h HTMLVideoElement) GetLang() string {
	val := h.Get("lang")
	return val.String()
}
func (h HTMLVideoElement) SetLang(val string) {
	h.Set("lang", val)
}
func (h HTMLVideoElement) GetLastChild() Node {
	val := h.Get("lastChild")
	return JSValueToNode(val.JSValue())
}
func (h HTMLVideoElement) Load(args ...interface{}) {
	h.Call("load", args...)
}
func (h HTMLVideoElement) GetLocalName() string {
	val := h.Get("localName")
	return val.String()
}
func (h HTMLVideoElement) LookupNamespaceURI(args ...interface{}) string {
	val := h.Call("lookupNamespaceURI", args...)
	return val.String()
}
func (h HTMLVideoElement) LookupPrefix(args ...interface{}) string {
	val := h.Call("lookupPrefix", args...)
	return val.String()
}
func (h HTMLVideoElement) GetLoop() bool {
	val := h.Get("loop")
	return val.Bool()
}
func (h HTMLVideoElement) SetLoop(val bool) {
	h.Set("loop", val)
}
func (h HTMLVideoElement) Matches(args ...interface{}) bool {
	val := h.Call("matches", args...)
	return val.Bool()
}
func (h HTMLVideoElement) GetMuted() bool {
	val := h.Get("muted")
	return val.Bool()
}
func (h HTMLVideoElement) SetMuted(val bool) {
	h.Set("muted", val)
}
func (h HTMLVideoElement) GetNamespaceURI() string {
	val := h.Get("namespaceURI")
	return val.String()
}
func (h HTMLVideoElement) GetNetworkState() int {
	val := h.Get("networkState")
	return val.Int()
}
func (h HTMLVideoElement) GetNextSibling() Node {
	val := h.Get("nextSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLVideoElement) GetNodeName() string {
	val := h.Get("nodeName")
	return val.String()
}
func (h HTMLVideoElement) GetNodeType() int {
	val := h.Get("nodeType")
	return val.Int()
}
func (h HTMLVideoElement) GetNodeValue() string {
	val := h.Get("nodeValue")
	return val.String()
}
func (h HTMLVideoElement) SetNodeValue(val string) {
	h.Set("nodeValue", val)
}
func (h HTMLVideoElement) Normalize(args ...interface{}) {
	h.Call("normalize", args...)
}
func (h HTMLVideoElement) GetOwnerDocument() Document {
	val := h.Get("ownerDocument")
	return JSValueToDocument(val.JSValue())
}
func (h HTMLVideoElement) GetParentElement() Element {
	val := h.Get("parentElement")
	return JSValueToElement(val.JSValue())
}
func (h HTMLVideoElement) GetParentNode() Node {
	val := h.Get("parentNode")
	return JSValueToNode(val.JSValue())
}
func (h HTMLVideoElement) Pause(args ...interface{}) {
	h.Call("pause", args...)
}
func (h HTMLVideoElement) GetPaused() bool {
	val := h.Get("paused")
	return val.Bool()
}
func (h HTMLVideoElement) Play(args ...interface{}) {
	h.Call("play", args...)
}
func (h HTMLVideoElement) GetPlaybackRate() float64 {
	val := h.Get("playbackRate")
	return val.Float()
}
func (h HTMLVideoElement) SetPlaybackRate(val float64) {
	h.Set("playbackRate", val)
}
func (h HTMLVideoElement) GetPlayed() TimeRanges {
	val := h.Get("played")
	return JSValueToTimeRanges(val.JSValue())
}
func (h HTMLVideoElement) GetPlaysInline() bool {
	val := h.Get("playsInline")
	return val.Bool()
}
func (h HTMLVideoElement) SetPlaysInline(val bool) {
	h.Set("playsInline", val)
}
func (h HTMLVideoElement) GetPoster() string {
	val := h.Get("poster")
	return val.String()
}
func (h HTMLVideoElement) SetPoster(val string) {
	h.Set("poster", val)
}
func (h HTMLVideoElement) GetPrefix() string {
	val := h.Get("prefix")
	return val.String()
}
func (h HTMLVideoElement) GetPreload() string {
	val := h.Get("preload")
	return val.String()
}
func (h HTMLVideoElement) SetPreload(val string) {
	h.Set("preload", val)
}
func (h HTMLVideoElement) GetPreviousSibling() Node {
	val := h.Get("previousSibling")
	return JSValueToNode(val.JSValue())
}
func (h HTMLVideoElement) GetReadyState() int {
	val := h.Get("readyState")
	return val.Int()
}
func (h HTMLVideoElement) RemoveAttribute(args ...interface{}) {
	h.Call("removeAttribute", args...)
}
func (h HTMLVideoElement) RemoveAttributeNS(args ...interface{}) {
	h.Call("removeAttributeNS", args...)
}
func (h HTMLVideoElement) RemoveAttributeNode(args ...interface{}) Attr {
	val := h.Call("removeAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLVideoElement) RemoveChild(args ...interface{}) Node {
	val := h.Call("removeChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLVideoElement) RemoveEventListener(args ...interface{}) {
	h.Call("removeEventListener", args...)
}
func (h HTMLVideoElement) ReplaceChild(args ...interface{}) Node {
	val := h.Call("replaceChild", args...)
	return JSValueToNode(val.JSValue())
}
func (h HTMLVideoElement) GetSeekable() TimeRanges {
	val := h.Get("seekable")
	return JSValueToTimeRanges(val.JSValue())
}
func (h HTMLVideoElement) GetSeeking() bool {
	val := h.Get("seeking")
	return val.Bool()
}
func (h HTMLVideoElement) SetAttribute(args ...interface{}) {
	h.Call("setAttribute", args...)
}
func (h HTMLVideoElement) SetAttributeNS(args ...interface{}) {
	h.Call("setAttributeNS", args...)
}
func (h HTMLVideoElement) SetAttributeNode(args ...interface{}) Attr {
	val := h.Call("setAttributeNode", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLVideoElement) SetAttributeNodeNS(args ...interface{}) Attr {
	val := h.Call("setAttributeNodeNS", args...)
	return JSValueToAttr(val.JSValue())
}
func (h HTMLVideoElement) GetShadowRoot() ShadowRoot {
	val := h.Get("shadowRoot")
	return JSValueToShadowRoot(val.JSValue())
}
func (h HTMLVideoElement) GetSlot() string {
	val := h.Get("slot")
	return val.String()
}
func (h HTMLVideoElement) SetSlot(val string) {
	h.Set("slot", val)
}
func (h HTMLVideoElement) GetSpellcheck() bool {
	val := h.Get("spellcheck")
	return val.Bool()
}
func (h HTMLVideoElement) SetSpellcheck(val bool) {
	h.Set("spellcheck", val)
}
func (h HTMLVideoElement) GetSrc() string {
	val := h.Get("src")
	return val.String()
}
func (h HTMLVideoElement) SetSrc(val string) {
	h.Set("src", val)
}
func (h HTMLVideoElement) GetSrcObject() MediaProvider {
	val := h.Get("srcObject")
	return JSValueToMediaProvider(val.JSValue())
}
func (h HTMLVideoElement) SetSrcObject(val MediaProvider) {
	h.Set("srcObject", val)
}
func (h HTMLVideoElement) GetTagName() string {
	val := h.Get("tagName")
	return val.String()
}
func (h HTMLVideoElement) GetTextContent() string {
	val := h.Get("textContent")
	return val.String()
}
func (h HTMLVideoElement) SetTextContent(val string) {
	h.Set("textContent", val)
}
func (h HTMLVideoElement) GetTextTracks() TextTrackList {
	val := h.Get("textTracks")
	return JSValueToTextTrackList(val.JSValue())
}
func (h HTMLVideoElement) GetTitle() string {
	val := h.Get("title")
	return val.String()
}
func (h HTMLVideoElement) SetTitle(val string) {
	h.Set("title", val)
}
func (h HTMLVideoElement) ToggleAttribute(args ...interface{}) bool {
	val := h.Call("toggleAttribute", args...)
	return val.Bool()
}
func (h HTMLVideoElement) GetTranslate() bool {
	val := h.Get("translate")
	return val.Bool()
}
func (h HTMLVideoElement) SetTranslate(val bool) {
	h.Set("translate", val)
}
func (h HTMLVideoElement) GetVideoHeight() int {
	val := h.Get("videoHeight")
	return val.Int()
}
func (h HTMLVideoElement) GetVideoTracks() VideoTrackList {
	val := h.Get("videoTracks")
	return JSValueToVideoTrackList(val.JSValue())
}
func (h HTMLVideoElement) GetVideoWidth() int {
	val := h.Get("videoWidth")
	return val.Int()
}
func (h HTMLVideoElement) GetVolume() float64 {
	val := h.Get("volume")
	return val.Float()
}
func (h HTMLVideoElement) SetVolume(val float64) {
	h.Set("volume", val)
}
func (h HTMLVideoElement) WebkitMatchesSelector(args ...interface{}) bool {
	val := h.Call("webkitMatchesSelector", args...)
	return val.Bool()
}
func (h HTMLVideoElement) GetWidth() int {
	val := h.Get("width")
	return val.Int()
}
func (h HTMLVideoElement) SetWidth(val int) {
	h.Set("width", val)
}
