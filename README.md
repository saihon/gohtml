## saihon

HTML parser with JavaScript-like method name.

This package is using the following packages:
- [golang.org/x/net/html](https://golang.org/x/net/html)
- [github.com/andybalholm/cascadia](https://github.com/andybalholm/cascadia)




## Usage


```go

import (
    "github.com/saihon/saihon"
)

func main() {
    text := "<html><head></head><body></body></html>"
    document, err := saihon.Parse(strings.NewReader(text))
    if err != nil {
       return
    }

    documentElement := document.DocumentElement()
    all     := document.All()
    body    := document.Body()
    title   := document.Title() // string
    head    := document.Head()
    form    := document.Form()
    images  := document.Images()
    links   := document.Links()
    anchors := document.Anchors()


    // alias
    // GetElementsByTagName  : GetByTag
    // GetElementsByName     : GetByName
    // GetElementsByClassName: GetByClass
    // GetElementById        : GetById
    // QuerySelectorAll      : QueryAll
    // QuerySelector         : Query


    element := document.GetElementById("id")
    element = document.QuerySelector("div > p")
    // Should be verified
    if element != nil {
        textcontent := element.TextContent()
        // ...
    }


    elements := document.GetElementsByClassName("class")
    elements = document.QuerySelectorAll("div > p")
    elements = document.GetElementsByName("name")
    elements = document.GetElementsByTagName("p")
    // loop
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


    element.HasAttribute("key")
    element.GetAttribute("key")
    element.SetAttribute("key", "value")
    element.RemoveAttribute("key")
}


```


## License

[MIT](https://github.com/saihon/saihon/blob/master/LICENSE)

