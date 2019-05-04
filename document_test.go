package saihon

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var (
	test_html = `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>title text</title>
</head>
<body>
	<img>
	<area href="">
	<a href="" name="">
	<form></form>
	<img>
	<area href="">
	<a href="" name="">
	<form></form>
	<img>
	<area href="">
	<a href="" name="">
	<form></form>
</body>
</html>
`
)

func TestAll(t *testing.T) {
	doc, _ := Parse(strings.NewReader(test_html))
	a := doc.All()
	expect := 17
	actual := len(a.Nodes)
	if actual != expect {
		t.Errorf("\ngot : %v, want: %v\n", actual, expect)
	}
}

func TestDocumentElement(t *testing.T) {
	doc, _ := Parse(strings.NewReader(test_html))
	h := doc.DocumentElement()
	if h.Node.DataAtom != atom.Html {
		t.Errorf("\nreturn element is should be <html> element: %s\n", h.Node.Data)
	}
}

func TestBody(t *testing.T) {
	doc, _ := Parse(strings.NewReader(test_html))
	b := doc.Body()
	if b.Node.DataAtom != atom.Body {
		t.Errorf("\nreturn element is should be <body> element: %s\n", b.Node.Data)
	}
}

func TestTitle(t *testing.T) {
	doc, _ := Parse(strings.NewReader(test_html))
	actual := doc.Title()
	expect := "title text"
	if actual != expect {
		t.Errorf("\ngot : %v, want: %v\n", actual, expect)
	}
}

func TestHead(t *testing.T) {
	doc, _ := Parse(strings.NewReader(test_html))
	h := doc.Head()
	if h.Node.DataAtom != atom.Head {
		t.Errorf("\nreturn element is should be <head> element: %s\n", h.Node.Data)
	}
}

func TestForm(t *testing.T) {
	doc, _ := Parse(strings.NewReader(test_html))
	elements := doc.Form()

	expect := 3
	if len(elements.Nodes) != expect {
		t.Errorf("\ngot : %v, want: %v\n", len(elements.Nodes), expect)
		return
	}
	for i := 0; i < elements.Length(); i++ {
		e := elements.Get(i)
		if e.Node.DataAtom != atom.Form {
			t.Errorf("\nelement is should be <form> element: %s\n", e.Node.Data)
		}
	}
}

func TestImages(t *testing.T) {
	doc, _ := Parse(strings.NewReader(test_html))
	elements := doc.Images()

	expect := 3
	if len(elements.Nodes) != expect {
		t.Errorf("\ngot : %v, want: %v\n", len(elements.Nodes), expect)
		return
	}
	for i := 0; i < elements.Length(); i++ {
		e := elements.Get(i)
		if e.Node.DataAtom != atom.Img {
			t.Errorf("\nelement is should be <img> element: %s\n", e.Node.Data)
		}
	}
}

func TestLinks(t *testing.T) {
	doc, _ := Parse(strings.NewReader(test_html))
	elements := doc.Links()

	expect := 6
	if len(elements.Nodes) != expect {
		t.Errorf("\ngot : %v, want: %v\n", len(elements.Nodes), expect)
		return
	}
	for i := 0; i < elements.Length(); i++ {
		e := elements.Get(i)
		if e.Node.DataAtom != atom.A && e.Node.DataAtom != atom.Area {
			t.Errorf("\nelement is should be <a> or <area> element: %s\n", e.Node.Data)
		}
	}
}

func TestAnchors(t *testing.T) {
	doc, _ := Parse(strings.NewReader(test_html))
	elements := doc.Anchors()

	expect := 3
	if len(elements.Nodes) != expect {
		t.Errorf("\ngot : %v, want: %v\n", len(elements.Nodes), expect)
		return
	}
	for i := 0; i < elements.Length(); i++ {
		e := elements.Get(i)
		if e.Node.DataAtom != atom.A {
			t.Errorf("\nelement is should be <a> element: %s\n", e.Node.Data)
		}
	}
}

func TestCreateElement(t *testing.T) {
	tagname := "div"
	div := CreateElement(tagname)

	n := div.Node
	if n.Data != tagname || n.DataAtom != atom.Div || n.Type != html.ElementNode {
		t.Errorf(
			"\nelement is should be <div> and node type is an ElementNode: got : %v, want: %v\n",
			n.Data, tagname)
	}
}

func TestCreateTextNode(t *testing.T) {
	text := "text"
	element := CreateTextNode(text)

	n := element.Node
	if n.Data != text || n.Type != html.TextNode {
		t.Errorf(
			"\nelement node type is should be an TextNode: got : %v, want: %v\n",
			n.Data, text)
	}
}
