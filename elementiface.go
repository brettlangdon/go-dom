// DO NOT EDIT - generated file
package dom

type ElementIFace interface {
	QuerySelector(selector string) *Element
	QuerySelectorAll(selector string) []*Element
	AttachShadow(shadowRootInit ShadowRootInit) *ShadowRoot
}
