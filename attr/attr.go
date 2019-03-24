package attr

import (
	"reflect"

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

// HasKV
func HasKV(n *html.Node, key, value string) bool {
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
		// inherit a namespace if already there
		n.Attr[i] = html.Attribute{
			Key:       key,
			Val:       value,
			Namespace: n.Attr[i].Namespace,
		}
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
