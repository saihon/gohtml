package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getbody(n *html.Node) (*html.Node, error) {
	b := n.FirstChild
	if b == nil {
		goto ERROR
	}
	if b = b.FirstChild; b == nil {
		goto ERROR
	}
	if b = b.NextSibling; b == nil {
		goto ERROR
	}
	if b.DataAtom != atom.Body {
		return nil, fmt.Errorf("not a body %s", b.Data)
	}
	return b, nil
ERROR:
	return nil, errors.New("nil value not a body")
}

// innerHtml
func TestHtml(t *testing.T) {
	expect := "<div>hello world</div>"
	s := `<html><head></head><body>` + expect + `</body></html>`

	doc, _ := html.Parse(strings.NewReader(s))
	body, err := getbody(doc)
	if err != nil {
		t.Errorf("\n%v\n", err)
		return
	}

	// get
	actual := Html(body)
	if actual != expect {
		t.Errorf("\n1: got : %s\nwant: %s\n", actual, expect)
	}

	// set
	expect = `<p>hello world</p>`
	Html(body, expect)
	actual = Html(body)
	if actual != expect {
		t.Errorf("\n2 - got : %s\nwant: %s\n", actual, expect)
	}
}

// outerHtml
func TestHTML(t *testing.T) {
	expect := "<body><div>hello world</div></body>"
	s := `<html><head></head>` + expect + `</html>`

	doc, _ := html.Parse(strings.NewReader(s))
	body, err := getbody(doc)
	if err != nil {
		t.Errorf("\n%v\n", err)
		return
	}

	actual := HTML(body)
	if actual != expect {
		t.Errorf("\ngot : %s\nwant: %s\n", actual, expect)
	}
}

func TestText(t *testing.T) {
	expect := "hello world"
	s := `<html><head></head><body>` + expect + `</body></html>`

	doc, _ := html.Parse(strings.NewReader(s))
	body, err := getbody(doc)
	if err != nil {
		t.Errorf("\n%v\n", err)
		return
	}

	// get
	actual := Text(body)
	if actual != expect {
		t.Errorf("\ngot : %s\nwant: %s\n", actual, expect)
	}

	// set
	expect = "<p>foo bar baz</p>"
	Text(body, expect)
	actual = body.FirstChild.Data
	if actual != expect {
		t.Errorf("\ngot : %s\nwant: %s\n", actual, expect)
	}
}

func TestChildren(t *testing.T) {
	data := []struct {
		length int
		text   string
	}{
		{length: 0, text: ""},
		{length: 1, text: `<p></p>`},
		{length: 2, text: `<p></p><p></p>`},
		{length: 3, text: `<p></p><p></p><p></p>`},
		{length: 3, text: `<p>foo</p><p>bar</p><p>baz</p>`},
	}

	for i, v := range data {
		s := `<html><head></head><body>` + v.text + `</body></html>`
		doc, _ := html.Parse(strings.NewReader(s))
		body, err := getbody(doc)
		if err != nil {
			t.Errorf("\n%v\n", err)
			break
		}
		expect := v.length
		actual := len(Children(body))
		if actual != expect {
			t.Errorf("\n%d - got : %d\nwant: %d\n", i, actual, expect)
		}
	}
}

func TestParent(t *testing.T) {
	s := `<html><head></head><body></body></html>`
	doc, _ := html.Parse(strings.NewReader(s))
	body, err := getbody(doc)
	if err != nil {
		t.Errorf("\n%v\n", err)
		return
	}

	n := Parent(body)
	if n.DataAtom != atom.Html {
		t.Errorf("\nshould be parent is <html>\n")
	}
}

func TestFirst(t *testing.T) {
	s := `<html><head></head><body><p>first element child</p><span></span></body></html>`
	doc, _ := html.Parse(strings.NewReader(s))
	body, err := getbody(doc)
	if err != nil {
		t.Errorf("\n%v\n", err)
		return
	}

	n := First(body)
	if n.DataAtom != atom.P {
		t.Errorf("\nshould be parent is <p>\n")
	}
}

func TestNext(t *testing.T) {
	s := `<html><head></head><body><p></p><span></span></body></html>`
	doc, _ := html.Parse(strings.NewReader(s))
	body, err := getbody(doc)
	if err != nil {
		t.Errorf("\n%v\n", err)
		return
	}

	n := Next(body.FirstChild)
	if n.DataAtom != atom.Span {
		t.Errorf("\nshould be parent is <span>\n")
	}
}

func TestPrev(t *testing.T) {
	s := `<html><head></head><body><p></p><span></span></body></html>`
	doc, _ := html.Parse(strings.NewReader(s))
	body, err := getbody(doc)
	if err != nil {
		t.Errorf("\n%v\n", err)
		return
	}

	n := Prev(body.FirstChild.NextSibling)
	if n.DataAtom != atom.P {
		t.Errorf("\nshould be parent is <p>\n")
	}
}

func TestLast(t *testing.T) {
	s := `<html><head></head><body><p></p><span></span></body></html>`
	doc, _ := html.Parse(strings.NewReader(s))
	body, err := getbody(doc)
	if err != nil {
		t.Errorf("\n%v\n", err)
		return
	}

	n := Last(body)
	if n.DataAtom != atom.Span {
		t.Errorf("\nshould be parent is <span>\n")
	}
}

func TestRemove(t *testing.T) {
	s := `<html><head></head><body><div><p></p><span></span></div></body></html>`

	doc, _ := html.Parse(strings.NewReader(s))
	body, err := getbody(doc)
	if err != nil {
		t.Errorf("\n%v\n", err)
		return
	}

	Remove(body.FirstChild)
	if body.FirstChild != nil {
		t.Errorf("\nbody.FirstChild is should be nil\n")
	}
}

func TestEmpty(t *testing.T) {
	s := `<html><head></head><body><p></p><div></div><span></span></body></html>`

	doc, _ := html.Parse(strings.NewReader(s))
	body, err := getbody(doc)
	if err != nil {
		t.Errorf("\n%v\n", err)
		return
	}

	Empty(body)
	if body.FirstChild != nil || body.LastChild != nil {
		t.Errorf("\nbody is should not have node\n")
	}
}

func TestReplace(t *testing.T) {
	s := `<html><head></head><body><p>old element</p></body></html>`

	doc, _ := html.Parse(strings.NewReader(s))
	body, err := getbody(doc)
	if err != nil {
		t.Errorf("\n%v\n", err)
		return
	}

	expect := &html.Node{
		Type:     html.ElementNode,
		Data:     "div",
		DataAtom: atom.Div,
	}

	Replace(body, expect, body.FirstChild)

	actual := body.FirstChild
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("\ngot : %#v\nwant: %#v\n", *actual, *expect)
	}
}

func TestInsert(t *testing.T) {
	data := map[Position]string{
		Beforebegin: `<body><p>top</p><div>hello</div><p id="middle">middle</p><p>bottom</p></body>`,
		Afterbegin:  `<body><p>top</p><p id="middle"><div>hello</div>middle</p><p>bottom</p></body>`,
		Beforeend:   `<body><p>top</p><p id="middle">middle<div>hello</div></p><p>bottom</p></body>`,
		Afterend:    `<body><p>top</p><p id="middle">middle</p><div>hello</div><p>bottom</p></body>`,
	}

	for position, expect := range data {
		s := `<html><head></head><body><p>top</p><p id="middle">middle</p><p>bottom</p></body></html>`
		doc, _ := html.Parse(strings.NewReader(s))

		body, err := getbody(doc)
		if err != nil {
			t.Errorf("\n%v\n", err)
			break
		}

		middle := body.FirstChild.NextSibling
		if middle.Attr[0].Val != "middle" {
			t.Errorf("\n id is not a `middle' %s\n", middle.Attr[0].Val)
			return
		}

		node := &html.Node{
			Type:     html.ElementNode,
			Data:     `div`,
			DataAtom: atom.Div,
			FirstChild: &html.Node{
				Type: html.TextNode,
				Data: `hello`,
			},
		}
		Insert(position, middle, node)
		actual := HTML(body)
		if actual != expect {
			t.Errorf("\nposition: %d\n  got : \n%s\n  want: \n%s\n", position, actual, expect)
		}
	}
}
