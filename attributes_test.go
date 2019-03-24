package saihon

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestClassList(t *testing.T) {
	var expect DOMTokenList
	expect.List = []string{"foo", "bar", "baz"}
	expect.Value = strings.Join(expect.List, " ")

	s := `<!DOCTYPE html><html><head></head><body><div class="` + expect.Value + `"></div></body></html>`

	doc, _ := html.Parse(strings.NewReader(s))
	body := Document{doc}.Body()
	div := body.FirstElementChild()
	list := div.ClassList()

	if list.Length() != expect.Length() ||
		!reflect.DeepEqual(expect.List, list.List) ||
		list.Value != expect.Value {

		s1 := fmt.Sprintf("Value: %#v, List: %v, Length: %d", list.Value, list.List, list.Length())
		s2 := fmt.Sprintf("Value: %#v, List: %v, Length: %d", expect.Value, expect.List, expect.Length())
		t.Errorf("\nerror: DOMTokenList:\n  actual[%s]\n  expect[%s]\n", s1, s2)
	}
}
