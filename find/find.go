package find

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"github.com/andybalholm/cascadia"
	"github.com/saihon/saihon/attr"
	"github.com/saihon/saihon/utils"
)

var (
	// Whether cache use in Query and QueryAll
	// when using in multi thread maybe better set a false
	CacheEnabled = true
	cache        map[string]cascadia.Selector
)

func getSelector(key string) (cascadia.Selector, bool) {
	if CacheEnabled {
		if cache == nil {
			cache = make(map[string]cascadia.Selector)
		}
		s, ok := cache[key]
		if ok {
			return s, true
		}
	}

	s, err := cascadia.Compile(key)
	if err != nil {
		return s, false
	}

	if CacheEnabled {
		cache[key] = s
	}
	return s, true
}

// QueryAll
func QueryAll(n *html.Node, selector string) []*html.Node {
	s, ok := getSelector(selector)
	if !ok {
		return nil
	}
	v := s.MatchAll(n)
	if len(v) > 0 && v[0] == n {
		return v[1:]
	}
	return v
}

// Query
func Query(n *html.Node, selector string) *html.Node {
	s, ok := getSelector(selector)
	if !ok {
		return nil
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if utils.IsElement(c) {
			if v := s.MatchFirst(c); v != nil {
				return v
			}
		}
	}
	return nil
}

// Matcher
type Matcher func(*html.Node) bool

// All
func All(node *html.Node, m Matcher) []*html.Node {
	var nodes []*html.Node
	m.all(node, &nodes)
	return nodes
}

func (m Matcher) all(node *html.Node, nodes *[]*html.Node) {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if m(c) {
			*nodes = append(*nodes, c)
		}
		m.all(c, nodes)
	}
}

// First
func First(node *html.Node, m Matcher) *html.Node {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if m(c) {
			return c
		}
		if n := First(c, m); n != nil {
			return n
		}
	}
	return nil
}

// ById
func ById(n *html.Node, id string) *html.Node {
	var m Matcher = func(n *html.Node) bool {
		return utils.IsElement(n) && attr.HasKV(n, "id", id)
	}
	return First(n, m)
}

// ByTag
func ByTag(n *html.Node, tag string) []*html.Node {
	a := atom.Lookup([]byte(tag))
	if a == 0 {
		return nil
	}
	var m Matcher = func(n *html.Node) bool {
		return n.Type == html.ElementNode && n.DataAtom == a
	}
	return All(n, m)
}

// ByClass
func ByClass(n *html.Node, className string) []*html.Node {
	var m Matcher = func(n *html.Node) bool {
		return utils.IsElement(n) && attr.HasKV(n, "class", className)
	}
	return All(n, m)
}

// ByName
func ByName(n *html.Node, name string) []*html.Node {
	var m Matcher = func(n *html.Node) bool {
		return utils.IsElement(n) && attr.HasKV(n, "name", name)
	}
	return All(n, m)
}
