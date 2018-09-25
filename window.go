// Code generated DO NOT EDIT
// window.go
package dom

import "syscall/js"

type WindowIFace interface {
	AddEventListener(args ...interface{})
	Alert(args ...interface{})
	AlertWithArgs(args ...interface{})
	GetApplicationCache() ApplicationCache
	Atob(args ...interface{}) []byte
	Blur(args ...interface{})
	Btoa(args ...interface{}) string
	CancelAnimationFrame(args ...interface{})
	CaptureEvents(args ...interface{})
	ClearInterval(args ...interface{})
	ClearTimeout(args ...interface{})
	Close(args ...interface{})
	GetClosed() bool
	Confirm(args ...interface{}) bool
	CreateImageBitmap(args ...interface{})
	CreateImageBitmapWithArgs(args ...interface{})
	GetCustomElements() CustomElementRegistry
	DispatchEvent(args ...interface{}) bool
	GetDocument() Document
	GetEvent() Value
	GetExternal() External
	Focus(args ...interface{})
	GetFrameElement() Element
	GetFrames() WindowProxy
	GetComputedStyle(args ...interface{}) CSSStyleDeclaration
	GetHistory() History
	GetLength() int
	GetLocalStorage() Storage
	GetLocation() Location
	GetLocationbar() BarProp
	GetMenubar() BarProp
	GetName() string
	SetName(string)
	GetNavigator() Navigator
	GetOnabort() EventHandler
	SetOnabort(EventHandler)
	GetOnafterprint() EventHandler
	SetOnafterprint(EventHandler)
	GetOnauxclick() EventHandler
	SetOnauxclick(EventHandler)
	GetOnbeforeprint() EventHandler
	SetOnbeforeprint(EventHandler)
	GetOnbeforeunload() OnBeforeUnloadEventHandler
	SetOnbeforeunload(OnBeforeUnloadEventHandler)
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
	GetOncuechange() EventHandler
	SetOncuechange(EventHandler)
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
	GetOnhashchange() EventHandler
	SetOnhashchange(EventHandler)
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
	GetOnlanguagechange() EventHandler
	SetOnlanguagechange(EventHandler)
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
	GetOnmessage() EventHandler
	SetOnmessage(EventHandler)
	GetOnmessageerror() EventHandler
	SetOnmessageerror(EventHandler)
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
	GetOnoffline() EventHandler
	SetOnoffline(EventHandler)
	GetOnonline() EventHandler
	SetOnonline(EventHandler)
	GetOnpagehide() EventHandler
	SetOnpagehide(EventHandler)
	GetOnpageshow() EventHandler
	SetOnpageshow(EventHandler)
	GetOnpause() EventHandler
	SetOnpause(EventHandler)
	GetOnplay() EventHandler
	SetOnplay(EventHandler)
	GetOnplaying() EventHandler
	SetOnplaying(EventHandler)
	GetOnpopstate() EventHandler
	SetOnpopstate(EventHandler)
	GetOnprogress() EventHandler
	SetOnprogress(EventHandler)
	GetOnratechange() EventHandler
	SetOnratechange(EventHandler)
	GetOnrejectionhandled() EventHandler
	SetOnrejectionhandled(EventHandler)
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
	GetOnstorage() EventHandler
	SetOnstorage(EventHandler)
	GetOnsubmit() EventHandler
	SetOnsubmit(EventHandler)
	GetOnsuspend() EventHandler
	SetOnsuspend(EventHandler)
	GetOntimeupdate() EventHandler
	SetOntimeupdate(EventHandler)
	GetOntoggle() EventHandler
	SetOntoggle(EventHandler)
	GetOnunhandledrejection() EventHandler
	SetOnunhandledrejection(EventHandler)
	GetOnunload() EventHandler
	SetOnunload(EventHandler)
	GetOnvolumechange() EventHandler
	SetOnvolumechange(EventHandler)
	GetOnwaiting() EventHandler
	SetOnwaiting(EventHandler)
	GetOnwheel() EventHandler
	SetOnwheel(EventHandler)
	Open(args ...interface{}) WindowProxy
	GetOpener() Value
	SetOpener(Value)
	GetOrigin() string
	GetParent() WindowProxy
	GetPersonalbar() BarProp
	PostMessage(args ...interface{})
	Print(args ...interface{})
	Prompt(args ...interface{}) string
	QueueMicrotask(args ...interface{})
	ReleaseEvents(args ...interface{})
	RemoveEventListener(args ...interface{})
	RequestAnimationFrame(args ...interface{}) int
	GetScrollbars() BarProp
	GetSelf() WindowProxy
	GetSessionStorage() Storage
	SetInterval(args ...interface{}) int
	SetTimeout(args ...interface{}) int
	GetStatus() string
	SetStatus(string)
	GetStatusbar() BarProp
	Stop(args ...interface{})
	GetToolbar() BarProp
	GetTop() WindowProxy
	GetWindow() WindowProxy
}
type Window struct {
	Value
}

func JSValueToWindow(val js.Value) Window { return Window{Value: JSValueToValue(val)} }
func (v Value) AsWindow() Window          { return Window{Value: v} }
func NewWindow(args ...interface{}) Window {
	return Window{Value: JSValueToValue(js.Global().Get("Window").New(args...))}
}
func (w Window) AddEventListener(args ...interface{}) {
	w.Call("addEventListener", args...)
}
func (w Window) Alert(args ...interface{}) {
	w.Call("alert", args...)
}
func (w Window) AlertWithArgs(args ...interface{}) {
	w.Call("alertWithArgs", args...)
}
func (w Window) GetApplicationCache() ApplicationCache {
	val := w.Get("applicationCache")
	return JSValueToApplicationCache(val.JSValue())
}
func (w Window) Atob(args ...interface{}) []byte {
	val := w.Call("atob", args...)
	return []byte(val.String())
}
func (w Window) Blur(args ...interface{}) {
	w.Call("blur", args...)
}
func (w Window) Btoa(args ...interface{}) string {
	val := w.Call("btoa", args...)
	return val.String()
}
func (w Window) CancelAnimationFrame(args ...interface{}) {
	w.Call("cancelAnimationFrame", args...)
}
func (w Window) CaptureEvents(args ...interface{}) {
	w.Call("captureEvents", args...)
}
func (w Window) ClearInterval(args ...interface{}) {
	w.Call("clearInterval", args...)
}
func (w Window) ClearTimeout(args ...interface{}) {
	w.Call("clearTimeout", args...)
}
func (w Window) Close(args ...interface{}) {
	w.Call("close", args...)
}
func (w Window) GetClosed() bool {
	val := w.Get("closed")
	return val.Bool()
}
func (w Window) Confirm(args ...interface{}) bool {
	val := w.Call("confirm", args...)
	return val.Bool()
}
func (w Window) CreateImageBitmap(args ...interface{}) {
	w.Call("createImageBitmap", args...)
}
func (w Window) CreateImageBitmapWithArgs(args ...interface{}) {
	w.Call("createImageBitmapWithArgs", args...)
}
func (w Window) GetCustomElements() CustomElementRegistry {
	val := w.Get("customElements")
	return JSValueToCustomElementRegistry(val.JSValue())
}
func (w Window) DispatchEvent(args ...interface{}) bool {
	val := w.Call("dispatchEvent", args...)
	return val.Bool()
}
func (w Window) GetDocument() Document {
	val := w.Get("document")
	return JSValueToDocument(val.JSValue())
}
func (w Window) GetEvent() Value {
	val := w.Get("event")
	return val
}
func (w Window) GetExternal() External {
	val := w.Get("external")
	return JSValueToExternal(val.JSValue())
}
func (w Window) Focus(args ...interface{}) {
	w.Call("focus", args...)
}
func (w Window) GetFrameElement() Element {
	val := w.Get("frameElement")
	return JSValueToElement(val.JSValue())
}
func (w Window) GetFrames() WindowProxy {
	val := w.Get("frames")
	return JSValueToWindowProxy(val.JSValue())
}
func (w Window) GetComputedStyle(args ...interface{}) CSSStyleDeclaration {
	val := w.Call("getComputedStyle", args...)
	return JSValueToCSSStyleDeclaration(val.JSValue())
}
func (w Window) GetHistory() History {
	val := w.Get("history")
	return JSValueToHistory(val.JSValue())
}
func (w Window) GetLength() int {
	val := w.Get("length")
	return val.Int()
}
func (w Window) GetLocalStorage() Storage {
	val := w.Get("localStorage")
	return JSValueToStorage(val.JSValue())
}
func (w Window) GetLocation() Location {
	val := w.Get("location")
	return JSValueToLocation(val.JSValue())
}
func (w Window) GetLocationbar() BarProp {
	val := w.Get("locationbar")
	return JSValueToBarProp(val.JSValue())
}
func (w Window) GetMenubar() BarProp {
	val := w.Get("menubar")
	return JSValueToBarProp(val.JSValue())
}
func (w Window) GetName() string {
	val := w.Get("name")
	return val.String()
}
func (w Window) SetName(val string) {
	w.Set("name", val)
}
func (w Window) GetNavigator() Navigator {
	val := w.Get("navigator")
	return JSValueToNavigator(val.JSValue())
}
func (w Window) GetOnabort() EventHandler {
	val := w.Get("onabort")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnabort(val EventHandler) {
	w.Set("onabort", val)
}
func (w Window) GetOnafterprint() EventHandler {
	val := w.Get("onafterprint")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnafterprint(val EventHandler) {
	w.Set("onafterprint", val)
}
func (w Window) GetOnauxclick() EventHandler {
	val := w.Get("onauxclick")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnauxclick(val EventHandler) {
	w.Set("onauxclick", val)
}
func (w Window) GetOnbeforeprint() EventHandler {
	val := w.Get("onbeforeprint")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnbeforeprint(val EventHandler) {
	w.Set("onbeforeprint", val)
}
func (w Window) GetOnbeforeunload() OnBeforeUnloadEventHandler {
	val := w.Get("onbeforeunload")
	return JSValueToOnBeforeUnloadEventHandler(val.JSValue())
}
func (w Window) SetOnbeforeunload(val OnBeforeUnloadEventHandler) {
	w.Set("onbeforeunload", val)
}
func (w Window) GetOnblur() EventHandler {
	val := w.Get("onblur")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnblur(val EventHandler) {
	w.Set("onblur", val)
}
func (w Window) GetOncancel() EventHandler {
	val := w.Get("oncancel")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOncancel(val EventHandler) {
	w.Set("oncancel", val)
}
func (w Window) GetOncanplay() EventHandler {
	val := w.Get("oncanplay")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOncanplay(val EventHandler) {
	w.Set("oncanplay", val)
}
func (w Window) GetOncanplaythrough() EventHandler {
	val := w.Get("oncanplaythrough")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOncanplaythrough(val EventHandler) {
	w.Set("oncanplaythrough", val)
}
func (w Window) GetOnchange() EventHandler {
	val := w.Get("onchange")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnchange(val EventHandler) {
	w.Set("onchange", val)
}
func (w Window) GetOnclick() EventHandler {
	val := w.Get("onclick")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnclick(val EventHandler) {
	w.Set("onclick", val)
}
func (w Window) GetOnclose() EventHandler {
	val := w.Get("onclose")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnclose(val EventHandler) {
	w.Set("onclose", val)
}
func (w Window) GetOncontextmenu() EventHandler {
	val := w.Get("oncontextmenu")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOncontextmenu(val EventHandler) {
	w.Set("oncontextmenu", val)
}
func (w Window) GetOncuechange() EventHandler {
	val := w.Get("oncuechange")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOncuechange(val EventHandler) {
	w.Set("oncuechange", val)
}
func (w Window) GetOndblclick() EventHandler {
	val := w.Get("ondblclick")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOndblclick(val EventHandler) {
	w.Set("ondblclick", val)
}
func (w Window) GetOndrag() EventHandler {
	val := w.Get("ondrag")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOndrag(val EventHandler) {
	w.Set("ondrag", val)
}
func (w Window) GetOndragend() EventHandler {
	val := w.Get("ondragend")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOndragend(val EventHandler) {
	w.Set("ondragend", val)
}
func (w Window) GetOndragenter() EventHandler {
	val := w.Get("ondragenter")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOndragenter(val EventHandler) {
	w.Set("ondragenter", val)
}
func (w Window) GetOndragexit() EventHandler {
	val := w.Get("ondragexit")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOndragexit(val EventHandler) {
	w.Set("ondragexit", val)
}
func (w Window) GetOndragleave() EventHandler {
	val := w.Get("ondragleave")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOndragleave(val EventHandler) {
	w.Set("ondragleave", val)
}
func (w Window) GetOndragover() EventHandler {
	val := w.Get("ondragover")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOndragover(val EventHandler) {
	w.Set("ondragover", val)
}
func (w Window) GetOndragstart() EventHandler {
	val := w.Get("ondragstart")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOndragstart(val EventHandler) {
	w.Set("ondragstart", val)
}
func (w Window) GetOndrop() EventHandler {
	val := w.Get("ondrop")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOndrop(val EventHandler) {
	w.Set("ondrop", val)
}
func (w Window) GetOndurationchange() EventHandler {
	val := w.Get("ondurationchange")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOndurationchange(val EventHandler) {
	w.Set("ondurationchange", val)
}
func (w Window) GetOnemptied() EventHandler {
	val := w.Get("onemptied")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnemptied(val EventHandler) {
	w.Set("onemptied", val)
}
func (w Window) GetOnended() EventHandler {
	val := w.Get("onended")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnended(val EventHandler) {
	w.Set("onended", val)
}
func (w Window) GetOnerror() OnErrorEventHandler {
	val := w.Get("onerror")
	return JSValueToOnErrorEventHandler(val.JSValue())
}
func (w Window) SetOnerror(val OnErrorEventHandler) {
	w.Set("onerror", val)
}
func (w Window) GetOnfocus() EventHandler {
	val := w.Get("onfocus")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnfocus(val EventHandler) {
	w.Set("onfocus", val)
}
func (w Window) GetOnhashchange() EventHandler {
	val := w.Get("onhashchange")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnhashchange(val EventHandler) {
	w.Set("onhashchange", val)
}
func (w Window) GetOninput() EventHandler {
	val := w.Get("oninput")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOninput(val EventHandler) {
	w.Set("oninput", val)
}
func (w Window) GetOninvalid() EventHandler {
	val := w.Get("oninvalid")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOninvalid(val EventHandler) {
	w.Set("oninvalid", val)
}
func (w Window) GetOnkeydown() EventHandler {
	val := w.Get("onkeydown")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnkeydown(val EventHandler) {
	w.Set("onkeydown", val)
}
func (w Window) GetOnkeypress() EventHandler {
	val := w.Get("onkeypress")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnkeypress(val EventHandler) {
	w.Set("onkeypress", val)
}
func (w Window) GetOnkeyup() EventHandler {
	val := w.Get("onkeyup")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnkeyup(val EventHandler) {
	w.Set("onkeyup", val)
}
func (w Window) GetOnlanguagechange() EventHandler {
	val := w.Get("onlanguagechange")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnlanguagechange(val EventHandler) {
	w.Set("onlanguagechange", val)
}
func (w Window) GetOnload() EventHandler {
	val := w.Get("onload")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnload(val EventHandler) {
	w.Set("onload", val)
}
func (w Window) GetOnloadeddata() EventHandler {
	val := w.Get("onloadeddata")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnloadeddata(val EventHandler) {
	w.Set("onloadeddata", val)
}
func (w Window) GetOnloadedmetadata() EventHandler {
	val := w.Get("onloadedmetadata")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnloadedmetadata(val EventHandler) {
	w.Set("onloadedmetadata", val)
}
func (w Window) GetOnloadend() EventHandler {
	val := w.Get("onloadend")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnloadend(val EventHandler) {
	w.Set("onloadend", val)
}
func (w Window) GetOnloadstart() EventHandler {
	val := w.Get("onloadstart")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnloadstart(val EventHandler) {
	w.Set("onloadstart", val)
}
func (w Window) GetOnmessage() EventHandler {
	val := w.Get("onmessage")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnmessage(val EventHandler) {
	w.Set("onmessage", val)
}
func (w Window) GetOnmessageerror() EventHandler {
	val := w.Get("onmessageerror")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnmessageerror(val EventHandler) {
	w.Set("onmessageerror", val)
}
func (w Window) GetOnmousedown() EventHandler {
	val := w.Get("onmousedown")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnmousedown(val EventHandler) {
	w.Set("onmousedown", val)
}
func (w Window) GetOnmouseenter() EventHandler {
	val := w.Get("onmouseenter")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnmouseenter(val EventHandler) {
	w.Set("onmouseenter", val)
}
func (w Window) GetOnmouseleave() EventHandler {
	val := w.Get("onmouseleave")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnmouseleave(val EventHandler) {
	w.Set("onmouseleave", val)
}
func (w Window) GetOnmousemove() EventHandler {
	val := w.Get("onmousemove")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnmousemove(val EventHandler) {
	w.Set("onmousemove", val)
}
func (w Window) GetOnmouseout() EventHandler {
	val := w.Get("onmouseout")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnmouseout(val EventHandler) {
	w.Set("onmouseout", val)
}
func (w Window) GetOnmouseover() EventHandler {
	val := w.Get("onmouseover")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnmouseover(val EventHandler) {
	w.Set("onmouseover", val)
}
func (w Window) GetOnmouseup() EventHandler {
	val := w.Get("onmouseup")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnmouseup(val EventHandler) {
	w.Set("onmouseup", val)
}
func (w Window) GetOnoffline() EventHandler {
	val := w.Get("onoffline")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnoffline(val EventHandler) {
	w.Set("onoffline", val)
}
func (w Window) GetOnonline() EventHandler {
	val := w.Get("ononline")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnonline(val EventHandler) {
	w.Set("ononline", val)
}
func (w Window) GetOnpagehide() EventHandler {
	val := w.Get("onpagehide")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnpagehide(val EventHandler) {
	w.Set("onpagehide", val)
}
func (w Window) GetOnpageshow() EventHandler {
	val := w.Get("onpageshow")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnpageshow(val EventHandler) {
	w.Set("onpageshow", val)
}
func (w Window) GetOnpause() EventHandler {
	val := w.Get("onpause")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnpause(val EventHandler) {
	w.Set("onpause", val)
}
func (w Window) GetOnplay() EventHandler {
	val := w.Get("onplay")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnplay(val EventHandler) {
	w.Set("onplay", val)
}
func (w Window) GetOnplaying() EventHandler {
	val := w.Get("onplaying")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnplaying(val EventHandler) {
	w.Set("onplaying", val)
}
func (w Window) GetOnpopstate() EventHandler {
	val := w.Get("onpopstate")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnpopstate(val EventHandler) {
	w.Set("onpopstate", val)
}
func (w Window) GetOnprogress() EventHandler {
	val := w.Get("onprogress")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnprogress(val EventHandler) {
	w.Set("onprogress", val)
}
func (w Window) GetOnratechange() EventHandler {
	val := w.Get("onratechange")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnratechange(val EventHandler) {
	w.Set("onratechange", val)
}
func (w Window) GetOnrejectionhandled() EventHandler {
	val := w.Get("onrejectionhandled")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnrejectionhandled(val EventHandler) {
	w.Set("onrejectionhandled", val)
}
func (w Window) GetOnreset() EventHandler {
	val := w.Get("onreset")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnreset(val EventHandler) {
	w.Set("onreset", val)
}
func (w Window) GetOnresize() EventHandler {
	val := w.Get("onresize")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnresize(val EventHandler) {
	w.Set("onresize", val)
}
func (w Window) GetOnscroll() EventHandler {
	val := w.Get("onscroll")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnscroll(val EventHandler) {
	w.Set("onscroll", val)
}
func (w Window) GetOnsecuritypolicyviolation() EventHandler {
	val := w.Get("onsecuritypolicyviolation")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnsecuritypolicyviolation(val EventHandler) {
	w.Set("onsecuritypolicyviolation", val)
}
func (w Window) GetOnseeked() EventHandler {
	val := w.Get("onseeked")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnseeked(val EventHandler) {
	w.Set("onseeked", val)
}
func (w Window) GetOnseeking() EventHandler {
	val := w.Get("onseeking")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnseeking(val EventHandler) {
	w.Set("onseeking", val)
}
func (w Window) GetOnselect() EventHandler {
	val := w.Get("onselect")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnselect(val EventHandler) {
	w.Set("onselect", val)
}
func (w Window) GetOnstalled() EventHandler {
	val := w.Get("onstalled")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnstalled(val EventHandler) {
	w.Set("onstalled", val)
}
func (w Window) GetOnstorage() EventHandler {
	val := w.Get("onstorage")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnstorage(val EventHandler) {
	w.Set("onstorage", val)
}
func (w Window) GetOnsubmit() EventHandler {
	val := w.Get("onsubmit")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnsubmit(val EventHandler) {
	w.Set("onsubmit", val)
}
func (w Window) GetOnsuspend() EventHandler {
	val := w.Get("onsuspend")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnsuspend(val EventHandler) {
	w.Set("onsuspend", val)
}
func (w Window) GetOntimeupdate() EventHandler {
	val := w.Get("ontimeupdate")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOntimeupdate(val EventHandler) {
	w.Set("ontimeupdate", val)
}
func (w Window) GetOntoggle() EventHandler {
	val := w.Get("ontoggle")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOntoggle(val EventHandler) {
	w.Set("ontoggle", val)
}
func (w Window) GetOnunhandledrejection() EventHandler {
	val := w.Get("onunhandledrejection")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnunhandledrejection(val EventHandler) {
	w.Set("onunhandledrejection", val)
}
func (w Window) GetOnunload() EventHandler {
	val := w.Get("onunload")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnunload(val EventHandler) {
	w.Set("onunload", val)
}
func (w Window) GetOnvolumechange() EventHandler {
	val := w.Get("onvolumechange")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnvolumechange(val EventHandler) {
	w.Set("onvolumechange", val)
}
func (w Window) GetOnwaiting() EventHandler {
	val := w.Get("onwaiting")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnwaiting(val EventHandler) {
	w.Set("onwaiting", val)
}
func (w Window) GetOnwheel() EventHandler {
	val := w.Get("onwheel")
	return JSValueToEventHandler(val.JSValue())
}
func (w Window) SetOnwheel(val EventHandler) {
	w.Set("onwheel", val)
}
func (w Window) Open(args ...interface{}) WindowProxy {
	val := w.Call("open", args...)
	return JSValueToWindowProxy(val.JSValue())
}
func (w Window) GetOpener() Value {
	val := w.Get("opener")
	return val
}
func (w Window) SetOpener(val Value) {
	w.Set("opener", val)
}
func (w Window) GetOrigin() string {
	val := w.Get("origin")
	return val.String()
}
func (w Window) GetParent() WindowProxy {
	val := w.Get("parent")
	return JSValueToWindowProxy(val.JSValue())
}
func (w Window) GetPersonalbar() BarProp {
	val := w.Get("personalbar")
	return JSValueToBarProp(val.JSValue())
}
func (w Window) PostMessage(args ...interface{}) {
	w.Call("postMessage", args...)
}
func (w Window) Print(args ...interface{}) {
	w.Call("print", args...)
}
func (w Window) Prompt(args ...interface{}) string {
	val := w.Call("prompt", args...)
	return val.String()
}
func (w Window) QueueMicrotask(args ...interface{}) {
	w.Call("queueMicrotask", args...)
}
func (w Window) ReleaseEvents(args ...interface{}) {
	w.Call("releaseEvents", args...)
}
func (w Window) RemoveEventListener(args ...interface{}) {
	w.Call("removeEventListener", args...)
}
func (w Window) RequestAnimationFrame(args ...interface{}) int {
	val := w.Call("requestAnimationFrame", args...)
	return val.Int()
}
func (w Window) GetScrollbars() BarProp {
	val := w.Get("scrollbars")
	return JSValueToBarProp(val.JSValue())
}
func (w Window) GetSelf() WindowProxy {
	val := w.Get("self")
	return JSValueToWindowProxy(val.JSValue())
}
func (w Window) GetSessionStorage() Storage {
	val := w.Get("sessionStorage")
	return JSValueToStorage(val.JSValue())
}
func (w Window) SetInterval(args ...interface{}) int {
	val := w.Call("setInterval", args...)
	return val.Int()
}
func (w Window) SetTimeout(args ...interface{}) int {
	val := w.Call("setTimeout", args...)
	return val.Int()
}
func (w Window) GetStatus() string {
	val := w.Get("status")
	return val.String()
}
func (w Window) SetStatus(val string) {
	w.Set("status", val)
}
func (w Window) GetStatusbar() BarProp {
	val := w.Get("statusbar")
	return JSValueToBarProp(val.JSValue())
}
func (w Window) Stop(args ...interface{}) {
	w.Call("stop", args...)
}
func (w Window) GetToolbar() BarProp {
	val := w.Get("toolbar")
	return JSValueToBarProp(val.JSValue())
}
func (w Window) GetTop() WindowProxy {
	val := w.Get("top")
	return JSValueToWindowProxy(val.JSValue())
}
func (w Window) GetWindow() WindowProxy {
	val := w.Get("window")
	return JSValueToWindowProxy(val.JSValue())
}
