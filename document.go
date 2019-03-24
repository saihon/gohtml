package saihon

import (
	"io"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"github.com/saihon/saihon/attr"
	"github.com/saihon/saihon/find"
	"github.com/saihon/saihon/utils"
)

type Document Element

// Parse form io.Reader
func Parse(r io.Reader) (*Document, error) {
	n, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	return &Document{n}, nil
}

// DocumentElement
func (d Document) DocumentElement() *Element {
	var m find.Matcher = func(n *html.Node) bool {
		return utils.IsElement(n) && n.DataAtom == atom.Html
	}
	if n := find.First(d.Node, m); n != nil {
		return &Element{n}
	}
	return nil
}

// All
func (d Document) All() Collection {
	var m find.Matcher = func(n *html.Node) bool {
		return utils.IsElement(n)
	}
	var c Collection
	c.Nodes = find.All(d.Node, m)
	return c
}

// Body
func (d Document) Body() *Element {
	var m find.Matcher = func(n *html.Node) bool {
		return utils.IsElement(n) && n.DataAtom == atom.Body
	}
	if n := find.First(d.Node, m); n != nil {
		return &Element{n}
	}
	return nil
}

// Title returns a title text
func (d Document) Title() string {
	var m find.Matcher = func(n *html.Node) bool {
		return utils.IsElement(n) && n.DataAtom == atom.Title
	}
	if n := find.First(d.Node, m); n != nil {
		return n.Data
	}
	return ""
}

// Head
func (d Document) Head() *Element {
	var m find.Matcher = func(n *html.Node) bool {
		return utils.IsElement(n) && n.DataAtom == atom.Head
	}
	if n := find.First(d.Node, m); n != nil {
		return &Element{n}
	}
	return nil
}

// Form
func (d Document) Form() Collection {
	var m find.Matcher = func(n *html.Node) bool {
		return utils.IsElement(n) && n.DataAtom == atom.Form
	}
	var c Collection
	c.Nodes = find.All(d.Node, m)
	return c
}

// Images
func (d Document) Images() Collection {
	var m find.Matcher = func(n *html.Node) bool {
		// atom.Image ??
		return utils.IsElement(n) && n.DataAtom == atom.Img
	}
	var c Collection
	c.Nodes = find.All(d.Node, m)
	return c
}

// Links get all of "A" and "AREA" element who has "href" attribute
func (d Document) Links() Collection {
	var m find.Matcher = func(n *html.Node) bool {
		return utils.IsElement(n) && (n.DataAtom == atom.A || n.DataAtom == atom.Area) && attr.Has(n, "href")
	}
	var c Collection
	c.Nodes = find.All(d.Node, m)
	return c
}

// Anchors get all of "A" element who has "name" attribute
func (d Document) Anchors() Collection {
	var m find.Matcher = func(n *html.Node) bool {
		return utils.IsElement(n) && n.DataAtom == atom.A && attr.Has(n, "name")
	}
	var c Collection
	c.Nodes = find.All(d.Node, m)
	return c
}

// CreateElement
func CreateElement(tagname string) *Element {
	tagname = strings.ToLower(tagname)
	n := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Lookup([]byte(tagname)),
		Data:     tagname,
	}
	return &Element{n}
}

// CreateElement
func (_ Document) CreateElement(tagname string) *Element {
	return CreateElement(tagname)
}

// CreateTextNode text data is escaped
func CreateTextNode(text string) *Element {
	n := &html.Node{
		Type: html.TextNode,
		Data: html.EscapeString(text),
	}
	return &Element{n}
}

// CreateTextNode
func (_ Document) CreateTextNode(text string) *Element {
	return CreateTextNode(text)
}

// GetElementsByTagName
func (d Document) GetElementsByTagName(tagname string) Collection {
	return Collection{find.ByTag(d.Node, tagname)}
}

// GetElementsByName
func (d Document) GetElementsByName(name string) Collection {
	return Collection{find.ByName(d.Node, name)}
}

// GetElementsByClassName
func (d Document) GetElementsByClassName(classname string) Collection {
	return Collection{find.ByClass(d.Node, classname)}
}

// QuerySelectorAll
func (d Document) QuerySelectorAll(s string) Collection {
	var c Collection
	c.Nodes = find.QueryAll(d.Node, s)
	return c
}

// GetElementById
func (d Document) GetElementById(id string) *Element {
	if n := find.ById(d.Node, id); n != nil {
		return &Element{n}
	}
	return nil
}

// QuerySelector
func (d Document) QuerySelector(s string) *Element {
	if n := find.Query(d.Node, s); n != nil {
		return &Element{Node: n}
	}
	return nil
}

// FirstChild
func (d Document) FirstChild() *html.Node {
	return d.Node.FirstChild
}

// LastChild
func (d Document) LastChild() *html.Node {
	return d.Node.LastChild
}

// NextSibling - Always returns nil!!
func (d Document) NextSibling() *html.Node {
	return d.Node.NextSibling
}

// PreviousSibling - Always returns nil!!
func (d Document) PreviousSibling() *html.Node {
	return d.Node.PrevSibling
}

// ParentNode - Always returns nil!!
func (d Document) ParentNode() *html.Node {
	return d.Node.Parent
}

// ChildNodes
func (d Document) ChildNodes() []*html.Node {
	return utils.ChildNodes(d.Node)
}

// HasChildNodes
func (d Document) HasChildNodes() bool {
	return d.Node.FirstChild != nil
}

// ParentElement - Always returns nil!!
func (d Document) ParentElement() *Element {
	if n := utils.Parent(d.Node); n != nil {
		return &Element{n}
	}
	return nil
}

// Children
func (d Document) Children() Collection {
	return Collection{utils.Children(d.Node)}
}

// FirstElementChild
func (d Document) FirstElementChild() *Element {
	if n := utils.First(d.Node); n != nil {
		return &Element{n}
	}
	return nil
}

// ChildElementCount
func (d Document) ChildElementCount() int {
	return utils.Count(d.Node)
}

// NextElementSibling - Always returns nil!!
func (d Document) NextElementSibling() *Element {
	if n := utils.Next(d.Node); n != nil {
		return &Element{n}
	}
	return nil
}

// PreviousElementSibling - Always returns nil!!
func (d Document) PreviousElementSibling() *Element {
	if n := utils.Prev(d.Node); n != nil {
		return &Element{n}
	}
	return nil
}

// LastElementChild
func (d Document) LastElementChild() *Element {
	if n := utils.Last(d.Node); n != nil {
		return &Element{n}
	}
	return nil
}

// RemoveChild
func (d Document) RemoveChild(c *Element) {
	d.Node.RemoveChild(c.Node)
}

// ReplaceChild
func (d Document) ReplaceChild(newElement, oldElement *Element) *Element {
	n := utils.Replace(d.Node, newElement.Node, oldElement.Node)
	return &Element{n}
}

// AppendChild
func (d Document) AppendChild(c *Element) {
	d.Node.AppendChild(c.Node)
}

// InsertBefore
func (d Document) InsertBefore(newChild, oldChild *Element) {
	d.Node.InsertBefore(newChild.Node, oldChild.Node)
}

// CloneNode
func (d Document) CloneNode() *Document {
	if n := utils.Clone(d.Node); n != nil {
		return &Document{n}
	}
	return nil
}

// TextContent - Always returns nil!!
func (d Document) TextContent(text ...string) string {
	return utils.Text(d.Node, text...)
}
