## saihon

JavaScript-like HTML parser for Go language.

<br/>

[![Build Status](https://travis-ci.org/saihon/saihon.svg?branch=master)](https://travis-ci.org/saihon/saihon) [![GoDoc](https://godoc.org/github.com/saihon/saihon?status.svg)](https://godoc.org/github.com/saihon/saihon)

<br>
<br>

## Usage


```go

import (
    "github.com/saihon/saihon"
)

func main() {
    text := "<html><head></head><body></body></html>"
    
    // parse from text HTML
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
    // should be verified
    if element != nil {
        textcontent := element.TextContent()
        // ...
    }


    // returns collection
    elements := document.GetElementsByClassName("class")
    elements = document.QuerySelectorAll("div > p")
    elements = document.GetElementsByName("name")
    elements = document.GetElementsByTagName("p")

    // each element
    for i := 0; i < elements.Length(); i++ {
        outerhtml := elements.Get(i).OuterHTML()
        // ...
    }
    // or 
    for element := range elements.Enumerator() {
        outerhtml := element.OuterHTML()
        // ...
    }


    // set
    element.TextContent("hello")
    // get
    textcontent := element.TextContent()
    // set
    element.InnerHTML("<p>hello</p>")
    // get
    innerhtml := element.InnerHTML()

    // get id
    id := element.HasAttribute("id")
    // get class name
    classname := element.GetAttribute("class")
    // set attribute
    element.SetAttribute("key", "value")
    // remove
    element.RemoveAttribute("key")
}


```

<br>

#### alias

```go

    GetByTag  : GetElementsByTagName
    GetByName : GetElementsByName
    GetByClass: GetElementsByClassName
    GetById   : GetElementById
    QueryAll  : QuerySelectorAll
    Query     : QuerySelector

```

<br>


## License

[MIT License](https://github.com/saihon/saihon/blob/master/LICENSE)

<br>
<br>
