package saihon

import (
	"regexp"
	"strings"

	"golang.org/x/net/html"

	"github.com/saihon/saihon/attr"
)

// Attributes
func (e Element) Attributes() []html.Attribute {
	return e.Node.Attr
}

// GetAttribute
func (e Element) GetAttribute(key string) string {
	return attr.Get(e.Node, key)
}

// GetAttributeNS
func (e Element) GetAttributeNS(namespace, key string) string {
	return attr.GetNS(e.Node, namespace, key)
}

// GetAttributeNode
func (e Element) GetAttributeNode(key string) (html.Attribute, bool) {
	return attr.GetNode(e.Node, key)
}

// GetAttributeNodeNS
func (e Element) GetAttributeNodeNS(namespace, key string) (html.Attribute, bool) {
	return attr.GetNodeNS(e.Node, namespace, key)
}

// SetAttribute
func (e Element) SetAttribute(key string, value string) {
	attr.Set(e.Node, key, value)
}

// SetAttributeNode
func (e Element) SetAttributeNode(a html.Attribute) {
	attr.SetNode(e.Node, a)
}

// SetAttributeNS
func (e Element) SetAttributeNS(namespace, key, value string) {
	attr.SetNS(e.Node, namespace, key, value)
}

// SetAttributeNodeNS
func (e Element) SetAttributeNodeNS(a html.Attribute) {
	attr.SetNodeNS(e.Node, a)
}

// HasAttributes
func (e Element) HasAttributes() bool {
	return len(e.Node.Attr) > 0
}

// HasAttribute
func (e Element) HasAttribute(key string) bool {
	return attr.Has(e.Node, key)
}

// HasAttributeNS
func (e Element) HasAttributeNS(namespace, key string) bool {
	return attr.HasNS(e.Node, namespace, key)
}

// RemoveAttribute
func (e Element) RemoveAttribute(key string) {
	attr.Remove(e.Node, key)
}

// RemoveAttributeNS
func (e Element) RemoveAttributeNS(namespace, key string) {
	attr.RemoveNS(e.Node, namespace, key)
}

// RemoveAttributeNode
func (e Element) RemoveAttributeNode(a html.Attribute) {
	attr.RemoveNode(e.Node, a)
}

// DOMTokenList
type DOMTokenList struct {
	List  []string
	Value string
}

var (
	respace = regexp.MustCompile(`\s+`)
)

// ClassList
func (e Element) ClassList() DOMTokenList {
	var tl DOMTokenList
	tl.Value = strings.TrimSpace(attr.Get(e.Node, "class"))
	if tl.Value == "" {
		return tl
	}

	tl.List = respace.Split(tl.Value, -1)
	return tl
}

// Length
func (t DOMTokenList) Length() int {
	return len(t.List)
}
