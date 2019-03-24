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

// Prev previousElementSibling. returns node type is html.ElementNode
func Prev(n *html.Node) *html.Node {
	for c := n.PrevSibling; c != nil; c = c.PrevSibling {
		if IsElement(c) {
			return c
		}
	}
	return nil
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

// Remove
func Remove(n *html.Node) {
	if p := Parent(n); p != nil {
		p.RemoveChild(n)
	}
}

// Empty
func Empty(n *html.Node) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
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

// InsertBB insert before begin
func InsertBB(pivot, n *html.Node) error {
	return Insert(Beforebegin, pivot, n)
}

// InsertAB insert after begin
func InsertAB(pivot, n *html.Node) error {
	return Insert(Afterbegin, pivot, n)
}

// InsertBE insert before end
func InsertBE(pivot, n *html.Node) error {
	return Insert(Beforeend, pivot, n)
}

// InsertAE insert after end
func InsertAE(pivot, n *html.Node) error {
	return Insert(Afterend, pivot, n)
}
