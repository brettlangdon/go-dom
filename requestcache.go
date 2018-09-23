// Code generated DO NOT EDIT
// requestcache.go
package dom

import "syscall/js"

type RequestCache string

const (
	RequestCacheDefault      RequestCache = "default"
	RequestCacheNoStore      RequestCache = "no-store"
	RequestCacheReload       RequestCache = "reload"
	RequestCacheNoCache      RequestCache = "no-cache"
	RequestCacheForceCache   RequestCache = "force-cache"
	RequestCacheOnlyIfCached RequestCache = "only-if-cached"
)

func JSValueToRequestCache(val js.Value) RequestCache {
	return RequestCache(val.String())
}
