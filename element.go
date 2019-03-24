package saihon

import (
	"strings"

	"golang.org/x/net/html"

	"github.com/saihon/saihon/attr"
	"github.com/saihon/saihon/find"
	"github.com/saihon/saihon/utils"
)

// Element
type Element struct {
	Node *html.Node
}

// GetElementsByTagName
func (e Element) GetElementsByTagName(tagname string) Collection {
	var c Collection
	c.Nodes = find.ByTag(e.Node, tagname)
	return c
}

// GetElementsByName
func (e Element) GetElementsByName(name string) Collection {
	var c Collection
	c.Nodes = find.ByName(e.Node, name)
	return c
}

// GetElementsByClassName
func (e Element) GetElementsByClassName(classname string) Collection {
	var c Collection
	c.Nodes = find.ByClass(e.Node, classname)
	return c
}

// QuerySelectorAll
func (e Element) QuerySelectorAll(s string) Collection {
	var c Collection
	c.Nodes = find.QueryAll(e.Node, s)
	return c
}

// GetElementById
func (e Element) GetElementById(id string) *Element {
	if n := find.ById(e.Node, id); n != nil {
		return &Element{Node: n}
	}
	return nil
}

// QuerySelector
func (e Element) QuerySelector(s string) *Element {
	if n := find.Query(e.Node, s); n != nil {
		return &Element{Node: n}
	}
	return nil
}

// FirstChild
func (e Element) FirstChild() *html.Node {
	return e.Node.FirstChild
}

// LastChild
func (e Element) LastChild() *html.Node {
	return e.Node.LastChild
}

// NextSibling
func (e Element) NextSibling() *html.Node {
	return e.Node.NextSibling
}

// PreviousSibling
func (e Element) PreviousSibling() *html.Node {
	return e.Node.PrevSibling
}

// ParentNode
func (e Element) ParentNode() *html.Node {
	return e.Node.Parent
}

// ChildNodes
func (e Element) ChildNodes() []*html.Node {
	return utils.ChildNodes(e.Node)
}

// HasChildNodes
func (e Element) HasChildNodes() bool {
	return e.Node.FirstChild != nil
}

// InnerHTML set or get inner html to an element
func (e Element) InnerHTML(text ...string) string {
	return utils.Html(e.Node, text...)
}

// OuterHTML including HTML of self
func (e Element) OuterHTML() string {
	return utils.HTML(e.Node)
}

// Children
func (e Element) Children() Collection {
	var nodes []*html.Node
	for c := e.Node.FirstChild; c != nil; c = c.NextSibling {
		if utils.IsElement(c) {
			nodes = append(nodes, c)
		}
	}
	return Collection{Nodes: nodes}
}

// ParentElement
func (e Element) ParentElement() *Element {
	if n := utils.Parent(e.Node); n != nil {
		return &Element{n}
	}
	return nil
}

// FirstElementChild
func (e Element) FirstElementChild() *Element {
	if n := utils.First(e.Node); n != nil {
		return &Element{Node: n}
	}
	return nil
}

// ChildElementCount
func (e Element) ChildElementCount() int {
	return utils.Count(e.Node)
}

// NextElementSibling
func (e Element) NextElementSibling() *Element {
	if n := utils.Next(e.Node); n != nil {
		return &Element{Node: n}
	}
	return nil
}

// PreviousElementSibling
func (e Element) PreviousElementSibling() *Element {
	if n := utils.Prev(e.Node); n != nil {
		return &Element{Node: n}
	}
	return nil
}

// LastElementChild
func (e Element) LastElementChild() *Element {
	if n := utils.Last(e.Node); n != nil {
		return &Element{Node: n}
	}
	return nil
}

// CloneNode
func (e Element) CloneNode() *Element {
	if n := utils.Clone(e.Node); n != nil {
		return &Element{n}
	}
	return nil
}

// Remove remove a element self
func (e Element) Remove() {
	utils.Remove(e.Node)
}

// RemoveChild
func (e Element) RemoveChild(c *Element) {
	e.Node.RemoveChild(c.Node)
}

// ReplaceChild returns a old element. an error internally if it is nil.
func (e Element) ReplaceChild(newElement, oldElement *Element) *Element {
	n := utils.Replace(e.Node, newElement.Node, oldElement.Node)
	return &Element{n}
}

// AppendChild
func (e Element) AppendChild(c *Element) {
	e.Node.AppendChild(c.Node)
}

// InsertBefore
func (e Element) InsertBefore(newChild, oldChild *Element) {
	e.Node.InsertBefore(newChild.Node, oldChild.Node)
}

type Position int

const (
	Beforebegin = Position(utils.Beforebegin)
	Afterbegin  = Position(utils.Afterbegin)
	Beforeend   = Position(utils.Beforeend)
	Afterend    = Position(utils.Afterend)
)

// InsertAdjacentHTML
func (e Element) InsertAdjacentHTML(p Position, texthtml string) error {
	nodes, err := html.ParseFragment(strings.NewReader(texthtml), &html.Node{Type: html.ElementNode})
	if err != nil {
		return err
	}
	for _, n := range nodes {
		if err := utils.Insert(utils.Position(p), e.Node, n); err != nil {
			return err
		}
	}
	return nil
}

// InsertAdjacentText
func (e Element) InsertAdjacentText(p Position, text string) error {
	n := &html.Node{
		Type: html.TextNode,
		Data: html.EscapeString(text),
	}
	return utils.Insert(utils.Position(p), e.Node, n)
}

// InsertAdjacentElement
func (e Element) InsertAdjacentElement(p Position, newElement *Element) error {
	return utils.Insert(utils.Position(p), e.Node, newElement.Node)
}

// TextContent set or get text to an element
func (e Element) TextContent(text ...string) string {
	return utils.Text(e.Node, text...)
}

// InnerText
func (e Element) InnerText(text ...string) string {
	return utils.Text(e.Node, text...)
}

// TagName returns value is uppercase
func (e Element) TagName() string {
	if e.Node.Type == html.ElementNode {
		return strings.ToUpper(e.Node.Data)
	}
	return ""
}

// LocalName returns value is lowercase
func (e Element) LocalName() string {
	if e.Node.Type == html.ElementNode {
		return strings.ToLower(e.Node.Data)
	}
	return ""
}

// Id
func (e Element) Id() string {
	return attr.Get(e.Node, "id")
}

// ClassName
func (e Element) ClassName() string {
	return attr.Get(e.Node, "class")
}
