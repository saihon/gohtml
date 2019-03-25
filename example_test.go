package saihon_test

import (
	"fmt"
	"strings"

	"github.com/saihon/saihon"
	"github.com/saihon/saihon/utils"
)

func Example() {

	text := `
<html>
<head></head>
<body>
	<div id="id">hello</div>
</body>
</html>`

	doc, err := saihon.Parse(strings.NewReader(text))
	if err != nil {
		return
	}

	v := doc.GetElementById("id")
	if v == nil {
		return
	}
	fmt.Println(v.TextContent()) // hello

	// Attribute
	// set
	v.SetAttribute("class", "class-1")
	// get
	classname := v.GetAttribute("class")
	fmt.Println(classname) // class-1

	// get body
	body := doc.Body()

	// remove
	body.RemoveChild(v)

	// create element
	div := saihon.CreateElement("div")
	// create text node
	textnode := saihon.CreateTextNode("hello world")
	div.AppendChild(textnode)
	body.AppendChild(div)

	// remove itself
	div.Remove()

	for _, v := range []string{"foo", "bar", "baz"} {
		p := saihon.CreateElement("p")
		t := saihon.CreateTextNode(v)
		p.AppendChild(p)
		body.AppendChild(t)
	}

	collection := body.GetElementsByTagName("p")
	for i := 0; i < collection.Length(); i++ {
		element := collection.Get(i)
		fmt.Println(element.InnerText())
	}
	// or
	for element := range collection.Enumerator() {
		fmt.Println(element.OuterHTML())
	}

	utils.Empty(body.Node)
}

func ExampleParse() {

	// Text HTML
	text := `<html><head></head><body></body></html>`

	document, err := saihon.Parse(strings.NewReader(text))
	if err != nil {
		return
	}

	fmt.Println(document.Title())

	//
	//
	// []byte:
	// saihon.Parse(bytes.NewReader(b))
	//

	//
	// File:
	//
	// fp, _ := os.Open("index.html")
	// saihon.Parse(fp)
	// fp.Close()
	//

	//
	// HTTP Response:
	//
	// resp, _ := http.Get("example.com")
	// saihon.Parse(resp.Body)
	// resp.Body.Close()
	//
}
