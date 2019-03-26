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

// GetElementsByTagName returns find all elements have given tagname
func (e Element) GetElementsByTagName(tagname string) Collection {
	var c Collection
	c.Nodes = find.ByTag(e.Node, tagname)
	return c
}

// GetElementsByName returns find all elements has given name
func (e Element) GetElementsByName(name string) Collection {
	var c Collection
	c.Nodes = find.ByName(e.Node, name)
	return c
}

// GetElementsByClassName returns find all elements has given classname
func (e Element) GetElementsByClassName(classname string) Collection {
	var c Collection
	c.Nodes = find.ByClass(e.Node, classname)
	return c
}

// QuerySelectorAll returns find all elements has given css selector
func (e Element) QuerySelectorAll(s string) Collection {
	var c Collection
	c.Nodes = find.QueryAll(e.Node, s)
	return c
}

// GetElementById returns find an element has given id
func (e Element) GetElementById(id string) *Element {
	if n := find.ById(e.Node, id); n != nil {
		return &Element{Node: n}
	}
	return nil
}

// QuerySelector returns find an element have given css selector
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

// Children returns all of child html.ElementNode as "Collection"
func (e Element) Children() Collection {
	var nodes []*html.Node
	for c := e.Node.FirstChild; c != nil; c = c.NextSibling {
		if utils.IsElement(c) {
			nodes = append(nodes, c)
		}
	}
	return Collection{Nodes: nodes}
}

// ParentElement returns parent node as "*Element"
func (e Element) ParentElement() *Element {
	if n := utils.Parent(e.Node); n != nil {
		return &Element{n}
	}
	return nil
}

// FirstElementChild returns first html.ElementNode as "*Element"
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

// NextElementSibling returns next html.ElementNode as "*Element"
func (e Element) NextElementSibling() *Element {
	if n := utils.Next(e.Node); n != nil {
		return &Element{Node: n}
	}
	return nil
}

// PreviousElementSibling returns previous html.ElementNode as "*Element"
func (e Element) PreviousElementSibling() *Element {
	if n := utils.Prev(e.Node); n != nil {
		return &Element{Node: n}
	}
	return nil
}

// LastElementChild returns last html.ElementNode as "*Element"
func (e Element) LastElementChild() *Element {
	if n := utils.Last(e.Node); n != nil {
		return &Element{Node: n}
	}
	return nil
}

// CloneNode returns clone "*Element"
func (e Element) CloneNode() *Element {
	if n := utils.Clone(e.Node); n != nil {
		return &Element{n}
	}
	return nil
}

// Remove delete Element itself
func (e Element) Remove() {
	utils.Remove(e.Node)
}

// RemoveChild remove a given "*Element"
// specified *Element is must be child
func (e Element) RemoveChild(c *Element) {
	e.Node.RemoveChild(c.Node)
}

// ReplaceChild returns old element. panic if an error
func (e Element) ReplaceChild(newElement, oldElement *Element) *Element {
	n := utils.Replace(e.Node, newElement.Node, oldElement.Node)
	return &Element{n}
}

// AppendChild append "*Element" as last child
func (e Element) AppendChild(c *Element) {
	e.Node.AppendChild(c.Node)
}

// InsertBefore inserts a newChild before the oldChild as child
func (e Element) InsertBefore(newChild, oldChild *Element) {
	e.Node.InsertBefore(newChild.Node, oldChild.Node)
}

// Position
// <!-- beforebegin -->
// <p>
//   <!-- afterbegin -->
//   childnodes
//   <!-- beforeend -->
// </p>
// <!-- afterend -->
type Position int

const (
	Beforebegin = Position(utils.Beforebegin)
	Afterbegin  = Position(utils.Afterbegin)
	Beforeend   = Position(utils.Beforeend)
	Afterend    = Position(utils.Afterend)
)

// InsertAdjacentHTML inserts text HTML as the html.ElementNode to specified position
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

// InsertAdjacentText inserts text as the html.TextNode to specified position
func (e Element) InsertAdjacentText(p Position, text string) error {
	n := &html.Node{
		Type: html.TextNode,
		Data: html.EscapeString(text),
	}
	return utils.Insert(utils.Position(p), e.Node, n)
}

// InsertAdjacentElement inserts element to specified position
func (e Element) InsertAdjacentElement(p Position, newElement *Element) error {
	return utils.Insert(utils.Position(p), e.Node, newElement.Node)
}

// TextContent set or get text to an element
func (e Element) TextContent(text ...string) string {
	return utils.Text(e.Node, text...)
}

// InnerText set or get text to an element
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

// Id returns id value of attribute or if element not has id empty string
func (e Element) Id() string {
	return attr.Get(e.Node, "id")
}

// ClassName returns class value of attribute or if element not has class empty string
func (e Element) ClassName() string {
	return attr.Get(e.Node, "class")
}
