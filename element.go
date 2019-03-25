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

// GetElementsByTagName find the all elements have specified tagname
func (e Element) GetElementsByTagName(tagname string) Collection {
	var c Collection
	c.Nodes = find.ByTag(e.Node, tagname)
	return c
}

// GetElementsByName find the all elements have specified name
func (e Element) GetElementsByName(name string) Collection {
	var c Collection
	c.Nodes = find.ByName(e.Node, name)
	return c
}

// GetElementsByClassName find the all elements have specified classname
func (e Element) GetElementsByClassName(classname string) Collection {
	var c Collection
	c.Nodes = find.ByClass(e.Node, classname)
	return c
}

// QuerySelectorAll find the all elements have specified css selector
func (e Element) QuerySelectorAll(s string) Collection {
	var c Collection
	c.Nodes = find.QueryAll(e.Node, s)
	return c
}

// GetElementById find a element have specified id
func (e Element) GetElementById(id string) *Element {
	if n := find.ById(e.Node, id); n != nil {
		return &Element{Node: n}
	}
	return nil
}

// QuerySelector find a element have specified css selector
func (e Element) QuerySelector(s string) *Element {
	if n := find.Query(e.Node, s); n != nil {
		return &Element{Node: n}
	}
	return nil
}

// FirstChild returns first child node
func (e Element) FirstChild() *html.Node {
	return e.Node.FirstChild
}

// LastChild returns last child node
func (e Element) LastChild() *html.Node {
	return e.Node.LastChild
}

// NextSibling returns next sibling node
func (e Element) NextSibling() *html.Node {
	return e.Node.NextSibling
}

// PreviousSibling returns previous sibling node
func (e Element) PreviousSibling() *html.Node {
	return e.Node.PrevSibling
}

// ParentNode returns parent node
func (e Element) ParentNode() *html.Node {
	return e.Node.Parent
}

// ChildNodes returns all of child nodes
func (e Element) ChildNodes() []*html.Node {
	return utils.ChildNodes(e.Node)
}

// HasChildNodes  returns true if "Document" has node
func (e Element) HasChildNodes() bool {
	return e.Node.FirstChild != nil
}

// InnerHTML set or get inner html to an element
func (e Element) InnerHTML(text ...string) string {
	return utils.Html(e.Node, text...)
}

// OuterHTML include element itself
func (e Element) OuterHTML() string {
	return utils.HTML(e.Node)
}

// Children returns all of the child html.ElementNode as the "Collection"
func (e Element) Children() Collection {
	var nodes []*html.Node
	for c := e.Node.FirstChild; c != nil; c = c.NextSibling {
		if utils.IsElement(c) {
			nodes = append(nodes, c)
		}
	}
	return Collection{Nodes: nodes}
}

// ParentElement returns parent node as the "*Element"
func (e Element) ParentElement() *Element {
	if n := utils.Parent(e.Node); n != nil {
		return &Element{n}
	}
	return nil
}

// FirstElementChild returns first html.ElementNode as the "*Element"
func (e Element) FirstElementChild() *Element {
	if n := utils.First(e.Node); n != nil {
		return &Element{Node: n}
	}
	return nil
}

// ChildElementCount returns the number of html.ElementNode
func (e Element) ChildElementCount() int {
	return utils.Count(e.Node)
}

// NextElementSibling returns next html.ElementNode as the "*Element"
func (e Element) NextElementSibling() *Element {
	if n := utils.Next(e.Node); n != nil {
		return &Element{Node: n}
	}
	return nil
}

// PreviousElementSibling returns previous html.ElementNode as the "*Element"
func (e Element) PreviousElementSibling() *Element {
	if n := utils.Prev(e.Node); n != nil {
		return &Element{Node: n}
	}
	return nil
}

// LastElementChild returns last html.ElementNode as the "*Element"
func (e Element) LastElementChild() *Element {
	if n := utils.Last(e.Node); n != nil {
		return &Element{Node: n}
	}
	return nil
}

// CloneNode clone "*Element"
func (e Element) CloneNode() *Element {
	if n := utils.Clone(e.Node); n != nil {
		return &Element{n}
	}
	return nil
}

// Remove delete the Element itself
func (e Element) Remove() {
	utils.Remove(e.Node)
}

// RemoveChild remove a given the "*Element"
// specified c is must be the child of "Element"
func (e Element) RemoveChild(c *Element) {
	e.Node.RemoveChild(c.Node)
}

// ReplaceChild returns a old element. panic if an error
func (e Element) ReplaceChild(newElement, oldElement *Element) *Element {
	n := utils.Replace(e.Node, newElement.Node, oldElement.Node)
	return &Element{n}
}

// AppendChild append "*Element" as a last child
func (e Element) AppendChild(c *Element) {
	e.Node.AppendChild(c.Node)
}

// InsertBefore inserts a newElement before the oldElement as a child of a "Element".
func (e Element) InsertBefore(newChild, oldChild *Element) {
	e.Node.InsertBefore(newChild.Node, oldChild.Node)
}

type Position int

const (
	Beforebegin = Position(utils.Beforebegin)
	Afterbegin  = Position(utils.Afterbegin)
	Beforeend   = Position(utils.Beforeend)
	Afterend    = Position(utils.Afterend)
	// <!-- beforebegin -->
	// <p>
	//   <!-- afterbegin -->
	//   childnodes
	//   <!-- beforeend -->
	// </p>
	// <!-- afterend -->
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

// InnerText same as above
func (e Element) InnerText(text ...string) string {
	return utils.Text(e.Node, text...)
}

// TagName returns string as uppercase
func (e Element) TagName() string {
	if e.Node.Type == html.ElementNode {
		return strings.ToUpper(e.Node.Data)
	}
	return ""
}

// LocalName returns string as lowercase
func (e Element) LocalName() string {
	if e.Node.Type == html.ElementNode {
		return strings.ToLower(e.Node.Data)
	}
	return ""
}

// Id get id
func (e Element) Id() string {
	return attr.Get(e.Node, "id")
}

// ClassName get class name
func (e Element) ClassName() string {
	return attr.Get(e.Node, "class")
}
