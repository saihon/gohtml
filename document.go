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

// DocumentElement returns <html> element
func (d Document) DocumentElement() *Element {
	var m find.Matcher = func(n *html.Node) bool {
		return utils.IsElement(n) && n.DataAtom == atom.Html
	}
	if n := find.First(d.Node, m); n != nil {
		return &Element{n}
	}
	return nil
}

// All returns all elements of node type ElementNode
func (d Document) All() Collection {
	var m find.Matcher = func(n *html.Node) bool {
		return utils.IsElement(n)
	}
	var c Collection
	c.Nodes = find.All(d.Node, m)
	return c
}

// Body returns <body> element
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

// Head returns <head> element
func (d Document) Head() *Element {
	var m find.Matcher = func(n *html.Node) bool {
		return utils.IsElement(n) && n.DataAtom == atom.Head
	}
	if n := find.First(d.Node, m); n != nil {
		return &Element{n}
	}
	return nil
}

// Form returns all <form> element
func (d Document) Form() Collection {
	var m find.Matcher = func(n *html.Node) bool {
		return utils.IsElement(n) && n.DataAtom == atom.Form
	}
	var c Collection
	c.Nodes = find.All(d.Node, m)
	return c
}

// Images returns all <img> element
func (d Document) Images() Collection {
	var m find.Matcher = func(n *html.Node) bool {
		// atom.Image ??
		return utils.IsElement(n) && n.DataAtom == atom.Img
	}
	var c Collection
	c.Nodes = find.All(d.Node, m)
	return c
}

// Links returns all of <a> and <area> element these have "href" attribute
func (d Document) Links() Collection {
	var m find.Matcher = func(n *html.Node) bool {
		return utils.IsElement(n) && (n.DataAtom == atom.A || n.DataAtom == atom.Area) && attr.Has(n, "href")
	}
	var c Collection
	c.Nodes = find.All(d.Node, m)
	return c
}

// Anchors returns all of <a> element these have "name" attribute
func (d Document) Anchors() Collection {
	var m find.Matcher = func(n *html.Node) bool {
		return utils.IsElement(n) && n.DataAtom == atom.A && attr.Has(n, "name")
	}
	var c Collection
	c.Nodes = find.All(d.Node, m)
	return c
}

// CreateElement create the html.ElementNode
// with specified tag name and then return as the "*Element"
func CreateElement(tagname string) *Element {
	tagname = strings.ToLower(tagname)
	n := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Lookup([]byte(tagname)),
		Data:     tagname,
	}
	return &Element{n}
}

// CreateElement can be called from the "Document"
// has no meaning and is the same the above
func (_ Document) CreateElement(tagname string) *Element {
	return CreateElement(tagname)
}

// CreateTextNode create the html.TextNode
// with specified text and returns the "*Element"
func CreateTextNode(text string) *Element {
	n := &html.Node{
		Type: html.TextNode,
		Data: text,
	}
	return &Element{n}
}

// CreateTextNode same the above
func (_ Document) CreateTextNode(text string) *Element {
	return CreateTextNode(text)
}

// GetElementsByTagName find the all elements have specified tagname
func (d Document) GetElementsByTagName(tagname string) Collection {
	return Collection{find.ByTag(d.Node, tagname)}
}

// GetElementsByName find the all elements have specified name
func (d Document) GetElementsByName(name string) Collection {
	return Collection{find.ByName(d.Node, name)}
}

// GetElementsByClassName find the all elements have specified classname
func (d Document) GetElementsByClassName(classname string) Collection {
	return Collection{find.ByClass(d.Node, classname)}
}

// QuerySelectorAll find the all elements have specified css selector
func (d Document) QuerySelectorAll(s string) Collection {
	return Collection{find.QueryAll(d.Node, s)}
}

// GetElementById find the element have specified id
func (d Document) GetElementById(id string) *Element {
	if n := find.ById(d.Node, id); n != nil {
		return &Element{n}
	}
	return nil
}

// QuerySelector find the first element have specified css selector
func (d Document) QuerySelector(s string) *Element {
	if n := find.Query(d.Node, s); n != nil {
		return &Element{Node: n}
	}
	return nil
}

// FirstChild returns first child node
func (d Document) FirstChild() *html.Node {
	return d.Node.FirstChild
}

// LastChild returns last child node
func (d Document) LastChild() *html.Node {
	return d.Node.LastChild
}

// NextSibling - returns nil!!
func (d Document) NextSibling() *html.Node {
	return d.Node.NextSibling
}

// PreviousSibling - returns nil!!
func (d Document) PreviousSibling() *html.Node {
	return d.Node.PrevSibling
}

// ParentNode - returns nil!!
func (d Document) ParentNode() *html.Node {
	return d.Node.Parent
}

// ChildNodes returns all of child nodes
func (d Document) ChildNodes() []*html.Node {
	return utils.ChildNodes(d.Node)
}

// HasChildNodes returns true if "Document" has node
func (d Document) HasChildNodes() bool {
	return d.Node.FirstChild != nil
}

// ParentElement - returns nil!!
func (d Document) ParentElement() *Element {
	if n := utils.Parent(d.Node); n != nil {
		return &Element{n}
	}
	return nil
}

// Children returns all of the child html.ElementNode as the "Collection"
func (d Document) Children() Collection {
	return Collection{utils.Children(d.Node)}
}

// FirstElementChild returns first html.ElementNode as the "*Element"
func (d Document) FirstElementChild() *Element {
	if n := utils.First(d.Node); n != nil {
		return &Element{n}
	}
	return nil
}

// ChildElementCount returns the number of html.ElementNode
func (d Document) ChildElementCount() int {
	return utils.Count(d.Node)
}

// NextElementSibling - returns nil!!
func (d Document) NextElementSibling() *Element {
	if n := utils.Next(d.Node); n != nil {
		return &Element{n}
	}
	return nil
}

// PreviousElementSibling - returns nil!!
func (d Document) PreviousElementSibling() *Element {
	if n := utils.Prev(d.Node); n != nil {
		return &Element{n}
	}
	return nil
}

// LastElementChild returns the last child html.ElementNode as the "*Element"
func (d Document) LastElementChild() *Element {
	if n := utils.Last(d.Node); n != nil {
		return &Element{n}
	}
	return nil
}

// RemoveChild remove a given the "*Element"
// specified "*Element" is must be the child of "Document"
func (d Document) RemoveChild(c *Element) {
	d.Node.RemoveChild(c.Node)
}

// ReplaceChild replace oldElement to newElement
// given "*Element" is both the must be "Document" child, and same node type
func (d Document) ReplaceChild(newElement, oldElement *Element) *Element {
	n := utils.Replace(d.Node, newElement.Node, oldElement.Node)
	return &Element{n}
}

// AppendChild append "*Element" as a last child
func (d Document) AppendChild(c *Element) {
	d.Node.AppendChild(c.Node)
}

// InsertBefore inserts a newElement before the oldElement as a child of a "Document".
func (d Document) InsertBefore(newChild, oldChild *Element) {
	d.Node.InsertBefore(newChild.Node, oldChild.Node)
}

// CloneNode clone "Document"
func (d Document) CloneNode() *Document {
	if n := utils.Clone(d.Node); n != nil {
		return &Document{n}
	}
	return nil
}

// TextContent - returns nil!!
func (d Document) TextContent(text ...string) string {
	return utils.Text(d.Node, text...)
}
