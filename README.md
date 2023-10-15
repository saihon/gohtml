## saihon

JavaScript-like HTML parser for Go language. HTML parsers exist for each programming language, but their usage differs. This package can use natural if know the JavaScript.

<br/>

[![GoDoc](https://pkg.go.dev/badge/github.com/saihon/saihon)](https://pkg.go.dev/github.com/saihon/saihon) [![Test](https://github.com/saihon/saihon/actions/workflows/go.yml/badge.svg)](https://github.com/saihon/saihon/actions/workflows/go.yml)

<br>
<br>

## Usage


```go

import (
    "github.com/saihon/saihon"
)

func main() {
    text := "<html><head></head><body></body></html>"

    // Parse text HTML
    document, err := saihon.Parse(strings.NewReader(text))
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
        textcontent := element.TextContent()
        // ...
    }

    // Get HTML collection
    elements := document.GetElementsByClassName("class")
    elements = document.QuerySelectorAll("div > p")
    elements = document.GetElementsByName("name")
    elements = document.GetElementsByTagName("p")

    // Each element
    for i := 0; i < elements.Length(); i++ {
        outerhtml := elements.Get(i).OuterHTML()
        // ...
    }
    // or 
    for element := range elements.Enumerator() {
        outerhtml := element.OuterHTML()
        // ...
    }

    // Set text content
    element.TextContent("hello")
    // Get text content
    textcontent := element.TextContent()
    // Set HTML
    element.InnerHTML("<p>hello</p>")
    // Get
    innerhtml := element.InnerHTML()

    // Get id attribute
    id := element.HasAttribute("id")
    // Get class name attribute
    classname := element.GetAttribute("class")
    // Set attribute
    element.SetAttribute("key", "value")
    // Remove attribute
    element.RemoveAttribute("key")
}

```

[godoc]:https://pkg.go.dev/github.com/saihon/saihon

For more detailed documentation is [here][godoc].

<br>


## License

[MIT License](https://github.com/saihon/saihon/blob/master/LICENSE)

<br>
<br>
