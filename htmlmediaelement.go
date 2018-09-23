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
	HTMLElement
	Element
	Node
	EventTarget
}

func JSValueToHTMLMediaElement(val js.Value) HTMLMediaElement {
	return HTMLMediaElement{Value: Value{Value: val}}
}
func (v Value) AsHTMLMediaElement() HTMLMediaElement { return HTMLMediaElement{Value: v} }
func (h HTMLMediaElement) AddTextTrack(args ...interface{}) TextTrack {
	val := h.Call("addTextTrack", args...)
	return JSValueToTextTrack(val.JSValue())
}
func (h HTMLMediaElement) GetAudioTracks() AudioTrackList {
	val := h.Get("audioTracks")
	return JSValueToAudioTrackList(val.JSValue())
}
func (h HTMLMediaElement) GetAutoplay() bool {
	val := h.Get("autoplay")
	return val.Bool()
}
func (h HTMLMediaElement) SetAutoplay(val bool) {
	h.Set("autoplay", val)
}
func (h HTMLMediaElement) GetBuffered() TimeRanges {
	val := h.Get("buffered")
	return JSValueToTimeRanges(val.JSValue())
}
func (h HTMLMediaElement) CanPlayType(args ...interface{}) CanPlayTypeResult {
	val := h.Call("canPlayType", args...)
	return JSValueToCanPlayTypeResult(val.JSValue())
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
func (h HTMLMediaElement) GetStartDate(args ...interface{}) Value {
	val := h.Call("getStartDate", args...)
	return val
}
func (h HTMLMediaElement) Load(args ...interface{}) {
	h.Call("load", args...)
}
func (h HTMLMediaElement) GetLoop() bool {
	val := h.Get("loop")
	return val.Bool()
}
func (h HTMLMediaElement) SetLoop(val bool) {
	h.Set("loop", val)
}
func (h HTMLMediaElement) GetMuted() bool {
	val := h.Get("muted")
	return val.Bool()
}
func (h HTMLMediaElement) SetMuted(val bool) {
	h.Set("muted", val)
}
func (h HTMLMediaElement) GetNetworkState() int {
	val := h.Get("networkState")
	return val.Int()
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
func (h HTMLMediaElement) GetPreload() string {
	val := h.Get("preload")
	return val.String()
}
func (h HTMLMediaElement) SetPreload(val string) {
	h.Set("preload", val)
}
func (h HTMLMediaElement) GetReadyState() int {
	val := h.Get("readyState")
	return val.Int()
}
func (h HTMLMediaElement) GetSeekable() TimeRanges {
	val := h.Get("seekable")
	return JSValueToTimeRanges(val.JSValue())
}
func (h HTMLMediaElement) GetSeeking() bool {
	val := h.Get("seeking")
	return val.Bool()
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
func (h HTMLMediaElement) GetTextTracks() TextTrackList {
	val := h.Get("textTracks")
	return JSValueToTextTrackList(val.JSValue())
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
