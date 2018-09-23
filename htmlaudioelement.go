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
	HTMLMediaElement
	HTMLElement
	Element
	Node
	EventTarget
}

func JSValueToHTMLAudioElement(val js.Value) HTMLAudioElement {
	return HTMLAudioElement{Value: Value{Value: val}}
}
func (v Value) AsHTMLAudioElement() HTMLAudioElement { return HTMLAudioElement{Value: v} }
