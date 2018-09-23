// Code generated DO NOT EDIT
// window/window.go
package window

import dom "github.com/brettlangdon/go-dom"
import "syscall/js"

var value dom.Window

func init() { value = dom.JSValueToWindow(js.Global().Get("window")) }
func AddEventListener(args ...interface{}) {
	value.Call("addEventListener", args...)
}
func Alert(args ...interface{}) {
	value.Call("alert", args...)
}
func AlertWithArgs(args ...interface{}) {
	value.Call("alertWithArgs", args...)
}
func GetApplicationCache() dom.ApplicationCache {
	return value.GetApplicationCache()
}
func Atob(args ...interface{}) []byte {
	val := value.Call("atob", args...)
	return []byte(val.String())
}
func Blur(args ...interface{}) {
	value.Call("blur", args...)
}
func Btoa(args ...interface{}) string {
	val := value.Call("btoa", args...)
	return val.String()
}
func CancelAnimationFrame(args ...interface{}) {
	value.Call("cancelAnimationFrame", args...)
}
func ClearInterval(args ...interface{}) {
	value.Call("clearInterval", args...)
}
func ClearTimeout(args ...interface{}) {
	value.Call("clearTimeout", args...)
}
func Close(args ...interface{}) {
	value.Call("close", args...)
}
func GetClosed() bool {
	return value.GetClosed()
}
func Confirm(args ...interface{}) bool {
	val := value.Call("confirm", args...)
	return val.Bool()
}
func CreateImageBitmap(args ...interface{}) {
	value.Call("createImageBitmap", args...)
}
func CreateImageBitmapWithArgs(args ...interface{}) {
	value.Call("createImageBitmapWithArgs", args...)
}
func GetCustomElements() dom.CustomElementRegistry {
	return value.GetCustomElements()
}
func DispatchEvent(args ...interface{}) bool {
	val := value.Call("dispatchEvent", args...)
	return val.Bool()
}
func GetDocument() dom.Document {
	return value.GetDocument()
}
func Focus(args ...interface{}) {
	value.Call("focus", args...)
}
func GetFrameElement() dom.Element {
	return value.GetFrameElement()
}
func GetFrames() dom.WindowProxy {
	return value.GetFrames()
}
func GetHistory() dom.History {
	return value.GetHistory()
}
func GetLength() int {
	return value.GetLength()
}
func GetLocalStorage() dom.Storage {
	return value.GetLocalStorage()
}
func GetLocation() dom.Location {
	return value.GetLocation()
}
func GetLocationbar() dom.BarProp {
	return value.GetLocationbar()
}
func GetMenubar() dom.BarProp {
	return value.GetMenubar()
}
func GetName() string {
	return value.GetName()
}
func SetName(val string) { value.SetName(val) }
func GetNavigator() dom.Navigator {
	return value.GetNavigator()
}
func GetOnabort() dom.EventHandler {
	return value.GetOnabort()
}
func SetOnabort(val dom.EventHandler) { value.SetOnabort(val) }
func GetOnafterprint() dom.EventHandler {
	return value.GetOnafterprint()
}
func SetOnafterprint(val dom.EventHandler) { value.SetOnafterprint(val) }
func GetOnauxclick() dom.EventHandler {
	return value.GetOnauxclick()
}
func SetOnauxclick(val dom.EventHandler) { value.SetOnauxclick(val) }
func GetOnbeforeprint() dom.EventHandler {
	return value.GetOnbeforeprint()
}
func SetOnbeforeprint(val dom.EventHandler) { value.SetOnbeforeprint(val) }
func GetOnbeforeunload() dom.OnBeforeUnloadEventHandler {
	return value.GetOnbeforeunload()
}
func SetOnbeforeunload(val dom.OnBeforeUnloadEventHandler) { value.SetOnbeforeunload(val) }
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
func GetOncuechange() dom.EventHandler {
	return value.GetOncuechange()
}
func SetOncuechange(val dom.EventHandler) { value.SetOncuechange(val) }
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
func GetOnhashchange() dom.EventHandler {
	return value.GetOnhashchange()
}
func SetOnhashchange(val dom.EventHandler) { value.SetOnhashchange(val) }
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
func GetOnlanguagechange() dom.EventHandler {
	return value.GetOnlanguagechange()
}
func SetOnlanguagechange(val dom.EventHandler) { value.SetOnlanguagechange(val) }
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
func GetOnmessage() dom.EventHandler {
	return value.GetOnmessage()
}
func SetOnmessage(val dom.EventHandler) { value.SetOnmessage(val) }
func GetOnmessageerror() dom.EventHandler {
	return value.GetOnmessageerror()
}
func SetOnmessageerror(val dom.EventHandler) { value.SetOnmessageerror(val) }
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
func GetOnoffline() dom.EventHandler {
	return value.GetOnoffline()
}
func SetOnoffline(val dom.EventHandler) { value.SetOnoffline(val) }
func GetOnonline() dom.EventHandler {
	return value.GetOnonline()
}
func SetOnonline(val dom.EventHandler) { value.SetOnonline(val) }
func GetOnpagehide() dom.EventHandler {
	return value.GetOnpagehide()
}
func SetOnpagehide(val dom.EventHandler) { value.SetOnpagehide(val) }
func GetOnpageshow() dom.EventHandler {
	return value.GetOnpageshow()
}
func SetOnpageshow(val dom.EventHandler) { value.SetOnpageshow(val) }
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
func GetOnpopstate() dom.EventHandler {
	return value.GetOnpopstate()
}
func SetOnpopstate(val dom.EventHandler) { value.SetOnpopstate(val) }
func GetOnprogress() dom.EventHandler {
	return value.GetOnprogress()
}
func SetOnprogress(val dom.EventHandler) { value.SetOnprogress(val) }
func GetOnratechange() dom.EventHandler {
	return value.GetOnratechange()
}
func SetOnratechange(val dom.EventHandler) { value.SetOnratechange(val) }
func GetOnrejectionhandled() dom.EventHandler {
	return value.GetOnrejectionhandled()
}
func SetOnrejectionhandled(val dom.EventHandler) { value.SetOnrejectionhandled(val) }
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
func GetOnstorage() dom.EventHandler {
	return value.GetOnstorage()
}
func SetOnstorage(val dom.EventHandler) { value.SetOnstorage(val) }
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
func GetOnunhandledrejection() dom.EventHandler {
	return value.GetOnunhandledrejection()
}
func SetOnunhandledrejection(val dom.EventHandler) { value.SetOnunhandledrejection(val) }
func GetOnunload() dom.EventHandler {
	return value.GetOnunload()
}
func SetOnunload(val dom.EventHandler) { value.SetOnunload(val) }
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
func Open(args ...interface{}) dom.WindowProxy {
	val := value.Call("open", args...)
	return dom.JSValueToWindowProxy(val.JSValue())
}
func GetOpener() dom.Value {
	return value.GetOpener()
}
func SetOpener(val dom.Value) { value.SetOpener(val) }
func GetOrigin() string {
	return value.GetOrigin()
}
func GetParent() dom.WindowProxy {
	return value.GetParent()
}
func GetPersonalbar() dom.BarProp {
	return value.GetPersonalbar()
}
func PostMessage(args ...interface{}) {
	value.Call("postMessage", args...)
}
func Print(args ...interface{}) {
	value.Call("print", args...)
}
func Prompt(args ...interface{}) string {
	val := value.Call("prompt", args...)
	return val.String()
}
func QueueMicrotask(args ...interface{}) {
	value.Call("queueMicrotask", args...)
}
func RemoveEventListener(args ...interface{}) {
	value.Call("removeEventListener", args...)
}
func RequestAnimationFrame(args ...interface{}) int {
	val := value.Call("requestAnimationFrame", args...)
	return val.Int()
}
func GetScrollbars() dom.BarProp {
	return value.GetScrollbars()
}
func GetSelf() dom.WindowProxy {
	return value.GetSelf()
}
func GetSessionStorage() dom.Storage {
	return value.GetSessionStorage()
}
func SetInterval(args ...interface{}) int {
	val := value.Call("setInterval", args...)
	return val.Int()
}
func SetTimeout(args ...interface{}) int {
	val := value.Call("setTimeout", args...)
	return val.Int()
}
func GetStatus() string {
	return value.GetStatus()
}
func SetStatus(val string) { value.SetStatus(val) }
func GetStatusbar() dom.BarProp {
	return value.GetStatusbar()
}
func Stop(args ...interface{}) {
	value.Call("stop", args...)
}
func GetToolbar() dom.BarProp {
	return value.GetToolbar()
}
func GetTop() dom.WindowProxy {
	return value.GetTop()
}
func GetWindow() dom.WindowProxy {
	return value.GetWindow()
}
