package attr

import (
	"reflect"
	"strings"

	"golang.org/x/net/html"
)

// IndexOf
func IndexOf(n *html.Node, key string) int {
	for i, v := range n.Attr {
		if v.Key == key {
			return i
		}
	}
	return -1
}

func IndexOfNode(n *html.Node, a html.Attribute) int {
	for i, v := range n.Attr {
		if reflect.DeepEqual(v, a) {
			return i
		}
	}
	return -1
}

// IndexOfNS
func IndexOfNS(n *html.Node, namespace, key string) int {
	for i, v := range n.Attr {
		if v.Namespace == namespace && v.Key == key {
			return i
		}
	}
	return -1
}

// Has
func Has(n *html.Node, key string) bool {
	for _, v := range n.Attr {
		if v.Key == key {
			return true
		}
	}
	return false
}

// HasValue returns the bool value whether has value in attribute as assuming no duplicates
func HasValue(n *html.Node, key, value string) bool {
	for _, v := range n.Attr {
		if v.Key == key {
			return v.Val == value
		}
	}
	return false
}

// HasNode
func HasNode(n *html.Node, a html.Attribute) bool {
	for _, v := range n.Attr {
		if reflect.DeepEqual(v, a) {
			return true
		}
	}
	return false
}

// HasNS
func HasNS(n *html.Node, namespace, key string) bool {
	for _, v := range n.Attr {
		if v.Namespace == namespace && v.Key == key {
			return true
		}
	}
	return false
}

// Get get attribute value
func Get(n *html.Node, key string) string {
	for _, v := range n.Attr {
		if v.Key == key {
			return v.Val
		}
	}
	return ""
}

// GetNode
func GetNode(n *html.Node, key string) (html.Attribute, bool) {
	for _, v := range n.Attr {
		if v.Key == key {
			return v, true
		}
	}
	return html.Attribute{}, false
}

// GetNS
func GetNS(n *html.Node, namespace, key string) string {
	for _, v := range n.Attr {
		if v.Namespace == namespace && v.Key == key {
			return v.Val
		}
	}
	return ""
}

// GetNodeNS
func GetNodeNS(n *html.Node, namespace, key string) (html.Attribute, bool) {
	for _, v := range n.Attr {
		if v.Namespace == namespace && v.Key == key {
			return v, true
		}
	}
	return html.Attribute{}, false
}

// Set
func Set(n *html.Node, key, value string) {
	i := IndexOf(n, key)
	if i >= 0 {
		// inherit a namespace if already set
		n.Attr[i].Val = value
		return
	}
	n.Attr = append(n.Attr, html.Attribute{Key: key, Val: value})
}

// SetNode
func SetNode(n *html.Node, a html.Attribute) {
	i := IndexOf(n, a.Key)
	if i >= 0 {
		n.Attr[i] = a
		return
	}
	n.Attr = append(n.Attr, a)
}

// SetNS
func SetNS(n *html.Node, namespace, key, value string) {
	i := IndexOfNS(n, namespace, key)
	if i >= 0 {
		n.Attr[i].Val = value
		return
	}
	n.Attr = append(n.Attr, html.Attribute{
		Namespace: namespace,
		Key:       key,
		Val:       value,
	})
}

// SetNodeNS
func SetNodeNS(n *html.Node, a html.Attribute) {
	i := IndexOfNS(n, a.Namespace, a.Key)
	if i >= 0 {
		n.Attr[i] = a
		return
	}
	n.Attr = append(n.Attr, a)
}

// Remove
func Remove(n *html.Node, key string) {
	// If there is a node with the same key but
	// different namespace, only remove first one
	i := IndexOf(n, key)
	if i >= 0 {
		n.Attr = append(n.Attr[:i], n.Attr[i+1:]...)
	}
}

// RemoveNode
func RemoveNode(n *html.Node, a html.Attribute) {
	i := IndexOfNode(n, a)
	if i >= 0 {
		n.Attr = append(n.Attr[:i], n.Attr[i+1:]...)
	}
}

// RemoveNS
func RemoveNS(n *html.Node, namespace, key string) {
	i := IndexOfNS(n, namespace, key)
	if i >= 0 {
		n.Attr = append(n.Attr[:i], n.Attr[i+1:]...)
	}
}

// Attr
// if given a 1 argument: treat as the "Key" and returns it value. same as Get
// if given a 2 argument: treat as the "Key" and "Value",
//                        sets these as the attribute. same as Set
// if given a 3 argument: treat as the "Namespace", "Key" and "Value",
//                        sets these as the attribute. same as SetNS
func Attr(n *html.Node, a ...string) string {
	switch len(a) {
	case 1:
		return Get(n, a[0])
	case 2:
		Set(n, a[0], a[1])
	case 3:
		SetNS(n, a[0], a[1], a[2])
	}
	return ""
}

// AddClass
func AddClass(n *html.Node, classname string) {
	i := IndexOf(n, "class")
	if i == -1 {
		n.Attr = append(n.Attr, html.Attribute{Key: "class", Val: classname})
		return
	}

	if hasClass(n.Attr[i].Val, classname) {
		return
	}

	n.Attr[i].Val = strings.TrimSpace(n.Attr[i].Val) + " " + classname
}

func hasClass(value, classname string) bool {
	for _, v := range strings.Split(value, " ") {
		if v == classname {
			return true
		}
	}
	return false
}

// HasClass
func HasClass(n *html.Node, classname string) bool {
	i := IndexOf(n, "class")
	if i == -1 {
		return false
	}

	return hasClass(n.Attr[i].Val, classname)
}

func removeClass(value, classname string) string {
	a := strings.Split(value, " ")
	for i := 0; i < len(a); i++ {
		if a[i] == classname {
			a = append(a[:i], a[i+1:]...)
			break
		}
	}
	return strings.Join(a, " ")
}

// RemoveClass
func RemoveClass(n *html.Node, classname string) {
	i := IndexOf(n, "class")
	if i == -1 {
		return
	}

	n.Attr[i].Val = removeClass(n.Attr[i].Val, classname)
}

// ToggleClass
func ToggleClass(n *html.Node, classname string) {
	i := IndexOf(n, "class")
	if i == -1 {
		n.Attr = append(n.Attr, html.Attribute{Key: "class", Val: classname})
		return
	}

	if hasClass(n.Attr[i].Val, classname) {
		n.Attr[i].Val = removeClass(n.Attr[i].Val, classname)
		return
	}

	n.Attr[i].Val = strings.TrimSpace(n.Attr[i].Val) + " " + classname
}
