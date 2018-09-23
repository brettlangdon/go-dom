// Code generated DO NOT EDIT
// referrerpolicy.go
package dom

import "syscall/js"

type ReferrerPolicy string

const (
	ReferrerPolicyEmpty                       ReferrerPolicy = "Empty"
	ReferrerPolicyNoReferrer                  ReferrerPolicy = "no-referrer"
	ReferrerPolicyNoReferrerWhenDowngrade     ReferrerPolicy = "no-referrer-when-downgrade"
	ReferrerPolicySameOrigin                  ReferrerPolicy = "same-origin"
	ReferrerPolicyOrigin                      ReferrerPolicy = "origin"
	ReferrerPolicyStrictOrigin                ReferrerPolicy = "strict-origin"
	ReferrerPolicyOriginWhenCrossOrigin       ReferrerPolicy = "origin-when-cross-origin"
	ReferrerPolicyStrictOriginWhenCrossOrigin ReferrerPolicy = "strict-origin-when-cross-origin"
	ReferrerPolicyUnsafeUrl                   ReferrerPolicy = "unsafe-url"
)

func JSValueToReferrerPolicy(val js.Value) ReferrerPolicy {
	return ReferrerPolicy(val.String())
}
