package utils

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func IsElement(n *html.Node) bool {
	return n.Type == html.ElementNode
}

func IsText(n *html.Node) bool {
	return n.Type == html.TextNode
}

func IsDoctype(n *html.Node) bool {
	return n.Type == html.DoctypeNode
}

func IsComment(n *html.Node) bool {
	return n.Type == html.CommentNode
}

func IsDocument(n *html.Node) bool {
	return n.Type == html.DocumentNode
}

func IsError(n *html.Node) bool {
	return n.Type == html.ErrorNode
}

// ChildNodes
func ChildNodes(n *html.Node) []*html.Node {
	var nodes []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, c)
	}
	return nodes
}

// Html innerHTML
func Html(n *html.Node, text ...string) string {
	if text == nil {
		var buf bytes.Buffer
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if err := html.Render(&buf, c); err != nil {
				break
			}
		}
		return buf.String()
	}

	t := strings.Join(text, "")
	nodes, err := html.ParseFragment(strings.NewReader(t), &html.Node{Type: html.ElementNode})
	if err != nil {
		return ""
	}
	n.RemoveChild(n.FirstChild)
	for _, node := range nodes {
		n.AppendChild(node)
	}
	return t
}

// HTML outerHTML
func HTML(n *html.Node) string {
	var buf bytes.Buffer
	html.Render(&buf, n)
	return buf.String()
}

// Text
func Text(n *html.Node, text ...string) string {
	if text == nil {
		var buf bytes.Buffer
		collectText(n, &buf)
		return buf.String()
	}

	Empty(n)

	s := strings.Join(text, " ")
	n.AppendChild(&html.Node{
		Type: html.TextNode,
		Data: s,
	})
	return s
}

func collectText(n *html.Node, buf *bytes.Buffer) {
	if n.Type == html.TextNode {
		buf.WriteString(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		collectText(c, buf)
	}
}

// Children returns node type is html.ElementNode
func Children(n *html.Node) []*html.Node {
	var nodes []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if IsElement(c) {
			nodes = append(nodes, c)
		}
	}
	return nodes
}

// Count childElementCount count child element
func Count(n *html.Node) int {
	x := 0
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if IsElement(c) {
			x++
		}
	}
	return x
}

// Parent parentElement. returns node type is html.ElementNode
func Parent(n *html.Node) *html.Node {
	for p := n.Parent; p != nil; p = p.Parent {
		if IsElement(p) {
			return p
		}
	}
	return nil
}

// First firstElementChild. returns node type is html.ElementNode
func First(n *html.Node) *html.Node {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if IsElement(c) {
			return c
		}
	}
	return nil
}

// Next nextElementSibling. returns node type is html.ElementNode
func Next(n *html.Node) *html.Node {
	for c := n.NextSibling; c != nil; c = c.NextSibling {
		if IsElement(c) {
			return c
		}
	}
	return nil
}

// NextAll returns all next sibling hhtml.ElementNode
func NextAll(n *html.Node) []*html.Node {
	var nodes []*html.Node
	for c := n.NextSibling; c != nil; c = c.NextSibling {
		if IsElement(c) {
			nodes = append(nodes, c)
		}
	}
	return nodes
}

// Prev previousElementSibling. returns node type is html.ElementNode
func Prev(n *html.Node) *html.Node {
	for c := n.PrevSibling; c != nil; c = c.PrevSibling {
		if IsElement(c) {
			return c
		}
	}
	return nil
}

// PrevAll returns all previous sibling html.ElementNode
func PrevAll(n *html.Node) []*html.Node {
	var nodes []*html.Node
	for c := n.PrevSibling; c != nil; c = c.PrevSibling {
		if IsElement(c) {
			nodes = append(nodes, c)
		}
	}
	return nodes
}

// Last lastElementChild. returns node type is html.ElementNode
func Last(n *html.Node) *html.Node {
	for c := n.LastChild; c != nil; c = c.PrevSibling {
		if IsElement(c) {
			return c
		}
	}
	return nil
}

// Sibling returns all sibling html.ElementNode
func Sibling(n *html.Node) []*html.Node {
	var nodes []*html.Node
	p := n.Parent
	for c := p.FirstChild; c != nil; c = c.NextSibling {
		if IsElement(c) {
			nodes = append(nodes, c)
		}
	}
	return nodes
}

// Clone cloneNode
func Clone(n *html.Node) *html.Node {
	node := &html.Node{
		Type:     n.Type,
		DataAtom: n.DataAtom,
		Data:     n.Data,
		Attr:     make([]html.Attribute, len(n.Attr)),
	}
	copy(node.Attr, n.Attr)
	return node
}

// LastDescendant if FirstChild is nil returns given html.Node
func LastDescendant(n *html.Node) *html.Node {
	c := Last(n)
	if c == nil {
		return n
	}

	for c != nil {
		if _c := Last(c); _c == nil {
			break
		} else {
			c = _c
		}
	}
	return c
}

// CloneAll clone while keeping all descendents
func CloneAll(n *html.Node) *html.Node {
	parent := Clone(n)
	cloneall(parent, n)
	return parent
}

func cloneall(parent, original *html.Node) {
	for c := original.FirstChild; c != nil; c = c.NextSibling {
		_c := Clone(c)
		Append(parent, _c)
		cloneall(_c, c)
	}
}

// Wrap wraps html.Node of the first argument specified
func Wrap(node, wrapper *html.Node) error {
	w := CloneAll(wrapper)
	last := LastDescendant(w)

	n, err := Create(HTML(node))
	if err != nil {
		return err
	}
	for _, nn := range n {
		Append(last, nn)
	}

	Before(node, w)
	Remove(node)
	return nil
}

// WrapAll wraps given nodes of the first argument by node of second argument
func WrapAll(nodes []*html.Node, wrapper *html.Node) error {
	if len(nodes) == 0 {
		return errors.New("given slice of the nodes is empty")
	}

	// if parent is nil inserts to "Afterend" position
	if Parent(wrapper) == nil {
		After(nodes[len(nodes)-1], wrapper)
	}

	last := LastDescendant(wrapper)
	for _, n := range nodes {
		c := CloneAll(n)
		Append(last, c)
		Remove(n)
	}

	return nil
}

// Remove
func Remove(n *html.Node) {
	if p := Parent(n); p != nil {
		p.RemoveChild(n)
	}
}

// Empty remove all of child nodes
func Empty(n *html.Node) {
	for c := n.LastChild; c != nil; c = n.LastChild {
		n.RemoveChild(c)
	}
}

// Replace replaceChild
func Replace(parentNode, newNode, oldNode *html.Node) *html.Node {
	if newNode.Type != oldNode.Type {
		panic("Replace called for a different type Node")
	}
	ok := false
	for c := parentNode.FirstChild; c != nil; c = c.NextSibling {
		if c == oldNode {
			ok = true
			break
		}
	}
	if !ok {
		panic("Replace called for a non-child Node")
	}

	if err := Insert(Beforebegin, oldNode, newNode); err != nil {
		panic(fmt.Sprintf("Replace %v", err))
	}

	_oldNode := Clone(oldNode)
	_oldNode.Parent = nil
	_oldNode.NextSibling = nil
	_oldNode.PrevSibling = nil
	return _oldNode
}

type Position int

const (
	Beforebegin Position = iota
	Afterbegin
	Beforeend
	Afterend
)

// Insert
// <!-- beforebegin -->
// <p>
//   <!-- afterbegin -->
//   childnodes
//   <!-- beforeend -->
// </p>
// <!-- afterend -->
func Insert(position Position, pivot, n *html.Node) error {
	switch position {
	case Beforebegin:
		if parent := Parent(pivot); parent != nil {
			parent.InsertBefore(n, pivot)
		} else {
			return errors.New("parent element not exist")
		}
	case Afterbegin:
		if first := pivot.FirstChild; first != nil {
			pivot.InsertBefore(n, first)
		} else {
			pivot.AppendChild(n)
		}
	case Beforeend:
		pivot.AppendChild(n)
	case Afterend:
		parent := Parent(pivot)
		if parent == nil {
			return errors.New("parent element not exist")
		}
		if next := pivot.NextSibling; next != nil {
			parent.InsertBefore(n, next)
		} else {
			parent.AppendChild(n)
		}
	default:
		return errors.New("invalid position")
	}
	return nil
}

// Before insert before begin
func Before(pivot, n *html.Node) error {
	return Insert(Beforebegin, pivot, n)
}

// Prepend insert after begin
func Prepend(pivot, n *html.Node) error {
	return Insert(Afterbegin, pivot, n)
}

// Append
func Append(p, c *html.Node) {
	p.AppendChild(c)
}

// After insert after end
func After(pivot, n *html.Node) error {
	return Insert(Afterend, pivot, n)
}

// Create create ElementNode from text
func Create(text string) ([]*html.Node, error) {
	return html.ParseFragment(strings.NewReader(text), &html.Node{Type: html.ElementNode})
}
