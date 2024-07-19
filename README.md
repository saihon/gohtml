## gohtml

JavaScript-like HTML parser for Go language. If function names, etc. are the same as in JavaScript, it will be easier for JavaScript users to use.

<br/>

[![GoDoc](https://pkg.go.dev/badge/github.com/saihon/gohtml)](https://pkg.go.dev/github.com/saihon/gohtml) [![Test](https://github.com/saihon/gohtml/actions/workflows/go.yml/badge.svg)](https://github.com/saihon/gohtml/actions/workflows/go.yml)

<br>
<br>

## Usage


```go

import (
    "github.com/saihon/gohtml"
)

func main() {
    text := "<html><head></head><body></body></html>"

    // Parse text HTML
    document, err := gohtml.Parse(strings.NewReader(text))
    if err != nil {
       return
    }

    documentElement := document.DocumentElement()
    all     := document.All()
    body    := document.Body()
    title   := document.Title() // title string
    head    := document.Head()
    form    := document.Form()
    images  := document.Images()
    links   := document.Links()
    anchors := document.Anchors()


    element := document.GetElementById("id")
    element = document.QuerySelector("div > p")
    // Must be check if element is nil.
    if element != nil {
        textContent := element.TextContent()
        // ...
    }

    // Get HTML collection
    elements := document.GetElementsByClassName("class")
    elements = document.QuerySelectorAll("div > p")
    elements = document.GetElementsByName("name")
    elements = document.GetElementsByTagName("p")

    // for loop (Recommended because fast)
    for i := 0; i < elements.Length(); i++ {
        outerHtml := elements.Get(i).OuterHTML()
        // ...
    }
    // or 
    for element := range elements.Enumerator() {
        outerHtml := element.OuterHTML()
        // ...
    }
    // or 
    elements.ForEach(func(element *Element, index int, collection Collection) {
        outerHtml := element.OuterHTML()
        // ...
    })


    // Set text content
    element.TextContent("hello")
    // Get text content
    textContent := element.TextContent()
    // Set HTML
    element.InnerHTML("<p>hello</p>")
    // Get
    innerHtml := element.InnerHTML()

    // Get id attribute
    id := element.HasAttribute("id")
    // Get class name attribute
    className := element.GetAttribute("class")
    // Set attribute
    element.SetAttribute("key", "value")
    // Remove attribute
    element.RemoveAttribute("key")
}

```

[godoc]:https://pkg.go.dev/github.com/saihon/gohtml

For more detailed documentation is [here][godoc].

<br>


## License

[MIT License](https://github.com/saihon/gohtml/blob/master/LICENSE)

<br>
<br>
