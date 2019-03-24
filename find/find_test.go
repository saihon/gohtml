package find

import (
	"strings"
	"testing"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func TestAll(t *testing.T) {
	s := `<html><head><meta charset="UTF-8"><title>title</title></head><body><p>hello world</p></body></html>`

	doc, _ := html.Parse(strings.NewReader(s))
	var m Matcher = func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			return true
		}
		return false
	}

	nodes := All(doc, m)
	actual := len(nodes)
	expect := 6
	if actual != expect {
		t.Errorf("\ngot : %d, want: %d\n", actual, expect)
	}
}

func TestFirst(t *testing.T) {
	s := `<html><head></head><body>hello world</body></html>`

	doc, _ := html.Parse(strings.NewReader(s))
	var m Matcher = func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			return true
		}
		return false
	}

	actual := First(doc, m)
	if actual == nil {
		t.Errorf("\n1 - should not be nil\n")
	}

	if actual.DataAtom != atom.Html {
		t.Errorf("\n1 - first element is should be <html>\n")
	}
}

func TestById(t *testing.T) {
	s := `<html><head></head><body><p id="id">hello world</p></body></html>`

	doc, _ := html.Parse(strings.NewReader(s))
	actual := ById(doc, "id")
	if actual == nil {
		t.Errorf("\nshould not be nil\n")
		return
	}
	if actual.DataAtom != atom.P {
		t.Errorf("\nfirst element is should be <p>\n")
	}
	ok := false
	for _, a := range actual.Attr {
		if a.Key == "id" && a.Val == "id" {
			ok = true
			break
		}
	}
	if !ok {
		t.Errorf("\nshould have a id\n")
	}
}

func TestByTag(t *testing.T) {
	s := `<html><head></head><body><p></p><p></p><p></p></body></html>`
	l := 3

	doc, _ := html.Parse(strings.NewReader(s))
	actual := ByTag(doc, "p")

	if len(actual) != l {
		t.Errorf("\ngot : %d, want: %d\n", len(actual), l)
	}

	for _, v := range actual {
		if v.DataAtom != atom.P {
			t.Errorf("\nfirst element is should be <p>\n")
		}
		if v.Type != html.ElementNode {
			t.Errorf("\nnode type is should be an ElementNode\n")
		}
	}
}

func TestByClass(t *testing.T) {
	s := `<html><head></head><body><p></p><p class="class"></p><p class="class"></p></body></html>`
	l := 2

	doc, _ := html.Parse(strings.NewReader(s))
	actual := ByClass(doc, "class")

	if len(actual) != l {
		t.Errorf("\ngot : %d, want: %d\n", len(actual), l)
	}

	for _, v := range actual {
		if v.DataAtom != atom.P {
			t.Errorf("\nfirst element is should be <p>\n")
		}
		if v.Type != html.ElementNode {
			t.Errorf("\nnode type is should be an ElementNode\n")
		}
		ok := false
		for _, a := range v.Attr {
			if a.Key == "class" && a.Val == "class" {
				ok = true
				break
			}
		}
		if !ok {
			t.Errorf("\nshould have a class name\n")
		}
	}
}

func TestByName(t *testing.T) {
	s := `<html><head></head><body><p name="name"></p><p></p><p name="name"></p></body></html>`
	l := 2

	doc, _ := html.Parse(strings.NewReader(s))
	actual := ByName(doc, "name")

	if len(actual) != l {
		t.Errorf("\ngot : %d, want: %d\n", len(actual), l)
	}

	for _, v := range actual {
		if v.DataAtom != atom.P {
			t.Errorf("\nfirst element is should be <p>\n")
		}
		if v.Type != html.ElementNode {
			t.Errorf("\nnode type is should be an ElementNode\n")
		}
		ok := false
		for _, a := range v.Attr {
			if a.Key == "name" && a.Val == "name" {
				ok = true
				break
			}
		}
		if !ok {
			t.Errorf("\nshould have a name\n")
		}
	}
}

var (
	doc *html.Node
)

func init() {
	// fp, _ := os.Open("index.html")
	// doc, _ = html.Parse(fp)
	// fp.Close()
}

func queryAllNoCached(n *html.Node, selector string) []*html.Node {
	s, err := cascadia.Compile(selector)
	if err != nil {
		return nil
	}
	return s.MatchAll(n)
}

func queryNoCached(n *html.Node, selector string) *html.Node {
	s, err := cascadia.Compile(selector)
	if err != nil {
		return nil
	}
	return s.MatchFirst(n)
}

func BenchmarkQueryAllCached(b *testing.B) {
	b.ResetTimer()
	CacheEnabled = true
	for i := 0; i < b.N; i++ {
		QueryAll(doc, "tr td a")
	}
}

func BenchmarkQueryAllNoCached(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queryAllNoCached(doc, "tr td a")
	}
}

func BenchmarkQueryCached(b *testing.B) {
	b.ResetTimer()
	CacheEnabled = true
	for i := 0; i < b.N; i++ {
		Query(doc, "tr td a")
	}
}

func BenchmarkQueryNoCached(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queryNoCached(doc, "tr td a")
	}
}
