package attr

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"golang.org/x/net/html"
)

func TestIndexOf(t *testing.T) {
	n := new(html.Node)

	n.Attr = []html.Attribute{
		{Key: "key-0"},
		{Key: "key-1"},
		{Key: "key-2"},
	}

	expect := 2
	actual := IndexOf(n, "key-2")
	if actual != expect {
		t.Errorf("\ngot : %d, want: %d\n", actual, expect)
	}
}

func TestIndexOfNode(t *testing.T) {
	n := new(html.Node)

	n.Attr = []html.Attribute{
		{Key: "key-0"},
		{Key: "key-1"},
		{Key: "key-2"},
	}

	expect := 1
	actual := IndexOfNode(n, html.Attribute{Key: "key-1"})
	if actual != expect {
		t.Errorf("\ngot : %d, want: %d\n", actual, expect)
	}
}

func TestIndexOfNS(t *testing.T) {
	n := new(html.Node)

	n.Attr = []html.Attribute{
		{Key: "key", Namespace: "foo"},
		{Key: "key", Namespace: "bar"},
		{Key: "key", Namespace: "baz"},
	}

	expect := 2
	actual := IndexOfNS(n, "baz", "key")
	if actual != expect {
		t.Errorf("\ngot : %d, want: %d\n", actual, expect)
	}
}

func TestHas(t *testing.T) {
	n := new(html.Node)

	n.Attr = []html.Attribute{
		{Key: "key-0"},
		{Key: "key-1"},
		{Key: "key-2"},
	}

	actual := Has(n, "key-2")
	if !actual {
		t.Errorf("\nshould have a key\n")
	}
}

func TestHasKV(t *testing.T) {
	n := new(html.Node)

	n.Attr = []html.Attribute{
		{Key: "key-0", Val: "value-0"},
		{Key: "key-1", Val: "value-1"},
		{Key: "key-2", Val: "value-2"},
	}

	N := 1
	actual := HasKV(n, fmt.Sprintf("key-%d", N), fmt.Sprintf("value-%d", N))
	if !actual {
		t.Errorf("\nshould have a key and value\n")
	}
}

func TestHasNS(t *testing.T) {
	n := new(html.Node)

	n.Attr = []html.Attribute{
		{Key: "key-0", Namespace: "namespace-0"},
		{Key: "key-1", Namespace: "namespace-1"},
		{Key: "key-2", Namespace: "namespace-1"},
		{Key: "key-3", Namespace: "namespace-1"},
		{Key: "key-2", Namespace: "namespace-2"},
	}

	actual := HasNS(n, "namespace-1", "key-1")
	if !actual {
		t.Errorf("\nshould have a namespace and key\n")
	}
}

func TestGet(t *testing.T) {
	n := new(html.Node)

	n.Attr = []html.Attribute{
		{Key: "key-0", Val: "value-0"},
		{Key: "key-1", Val: "value-1"},
		{Key: "key-2", Val: "value-2"},
		{Key: "key-3", Val: "value-3"},
	}

	N := 2
	expect := n.Attr[N].Val
	actual := Get(n, fmt.Sprintf("key-%d", N))
	if actual != expect {
		t.Errorf("\ngot : %s, want: %s\n", actual, expect)
	}
}

func TestGetNode(t *testing.T) {
	n := new(html.Node)

	n.Attr = []html.Attribute{
		{Key: "key-0", Val: "value-0"},
		{Key: "key-1", Val: "value-1"},
		{Key: "key-2", Val: "value-2"},
		{Key: "key-3", Val: "value-3"},
	}

	N := 2
	expect := n.Attr[N]
	actual, ok := GetNode(n, fmt.Sprintf("key-%d", N))
	if !ok || !reflect.DeepEqual(actual, expect) {
		t.Errorf("\ngot : %s, want: %s\n", actual, expect)
	}
}

func TestGetNS(t *testing.T) {
	n := new(html.Node)

	n.Attr = []html.Attribute{
		{Key: "key-0", Val: "value-0", Namespace: "namespace-0"},
		{Key: "key-2", Val: "value-2", Namespace: "namespace-2"},
		{Key: "key-1", Val: "value-1", Namespace: "namespace-1"},
		{Key: "key-2", Val: "value-2", Namespace: "namespace-1"},
		{Key: "key-3", Val: "value-3", Namespace: "namespace-1"},
	}

	expect := n.Attr[3].Val
	actual := GetNS(n, "namespace-2", "key-2")
	if actual != expect {
		t.Errorf("\ngot : %s, want: %s\n", actual, expect)
	}
}

func TestGetNodeNS(t *testing.T) {
	n := new(html.Node)

	n.Attr = []html.Attribute{
		{Key: "key-2", Val: "value-2", Namespace: "namespace-2"},
		{Key: "key-0", Val: "value-0", Namespace: "namespace-0"},
		{Key: "key-1", Val: "value-1", Namespace: "namespace-1"},
		{Key: "key-2", Val: "value-2", Namespace: "namespace-1"},
		{Key: "key-3", Val: "value-3", Namespace: "namespace-1"},
	}

	expect := n.Attr[0]
	actual, ok := GetNodeNS(n, "namespace-2", "key-2")
	if !ok || !reflect.DeepEqual(actual, expect) {
		t.Errorf("\ngot : %s, want: %s\n", actual, expect)
	}
}

func TestSet(t *testing.T) {
	n := new(html.Node)

	NS := "namespace-1"
	n.Attr = []html.Attribute{
		{Key: "key-0", Val: "value-0", Namespace: "namespace-0"},
		{Key: "key-1", Val: "value-1", Namespace: NS},
		{Key: "key-2", Val: "value-2", Namespace: "namespace-1"},
		{Key: "key-3", Val: "value-3", Namespace: "namespace-1"},
		{Key: "key-2", Val: "value-2", Namespace: "namespace-2"},
	}

	expect := "value-new"
	Set(n, "key-1", expect)
	actual := n.Attr[1].Val
	ns := n.Attr[1].Namespace
	if actual != expect || ns != NS {
		t.Errorf("\ngot : %s, want: %s\n", actual, expect)
	}

	N := 5
	a := html.Attribute{Key: "key-4", Val: "value-4"}
	Set(n, a.Key, a.Val)
	if !reflect.DeepEqual(n.Attr[N], a) {
		t.Errorf("\ngot : %s, want: %s\n", n.Attr[N], a)
	}
}

func TestSetNode(t *testing.T) {
	n := new(html.Node)

	n.Attr = []html.Attribute{
		{Key: "key-0", Val: "value-0"},
		{Key: "key-1", Val: "value-1"},
		{Key: "key-2", Val: "value-2"},
	}

	expect := html.Attribute{Key: "key-2", Val: "value-new"}
	SetNode(n, expect)
	actual := n.Attr[2]
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("\ngot : %s, want: %s\n", actual, expect)
	}
}

func TestSetNS(t *testing.T) {
	n := new(html.Node)

	n.Attr = []html.Attribute{
		{Key: "key-0", Val: "value-0", Namespace: "namespace-0"},
		{Key: "key-1", Val: "value-1", Namespace: "namespace-0"},
		{Key: "key-2", Val: "value-2", Namespace: "namespace-0"},
	}

	expect := html.Attribute{Key: "key-2", Val: "value-new", Namespace: "namespace-1"}
	SetNS(n, expect.Namespace, expect.Key, expect.Val)
	actual := n.Attr[3]
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("\ngot : %s, want: %s\n", actual, expect)
	}
}

func TestSetNodeNS(t *testing.T) {
	n := new(html.Node)

	n.Attr = []html.Attribute{
		{Key: "key-0", Val: "value-0", Namespace: "namespace-0"},
		{Key: "key-1", Val: "value-1", Namespace: "namespace-0"},
		{Key: "key-2", Val: "value-2", Namespace: "namespace-0"},
	}

	expect := html.Attribute{Key: "key-1", Val: "value-new", Namespace: "namespace-0"}
	SetNodeNS(n, expect)
	actual := n.Attr[1]
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("\ngot : %s, want: %s\n", actual, expect)
	}
}

func TestRemove(t *testing.T) {
	n := new(html.Node)

	n.Attr = []html.Attribute{
		{Key: "key-0", Val: "value-0"},
		{Key: "key-1", Val: "value-1"},
		{Key: "key-2", Val: "value-2"},
		{Key: "key-3", Val: "value-3"},
		{Key: "key-4", Val: "value-4"},
	}

	expect := len(n.Attr)
	for len(n.Attr) > 0 {
		index := rand.Intn(len(n.Attr))
		key := n.Attr[index].Key
		Remove(n, key)
		expect--
		actual := len(n.Attr)
		if expect != actual {
			t.Errorf("\ngot : %d, want: %d\n", actual, expect)
			break
		}
		if Has(n, key) {
			fmt.Printf("\nstill exist there\n")
		}
	}
}

func TestRemoveNode(t *testing.T) {
	n := new(html.Node)

	n.Attr = []html.Attribute{
		{Key: "key-0", Val: "value-0"},
		{Key: "key-1", Val: "value-1"},
		{Key: "key-2", Val: "value-2"},
		{Key: "key-3", Val: "value-3"},
		{Key: "key-4", Val: "value-4"},
	}

	expect := len(n.Attr)
	for len(n.Attr) > 0 {
		index := rand.Intn(len(n.Attr))
		a := n.Attr[index]
		RemoveNode(n, a)
		expect--
		actual := len(n.Attr)
		if expect != actual {
			t.Errorf("\ngot : %d, want: %d\n", actual, expect)
			break
		}
		if HasNode(n, a) {
			fmt.Printf("\nstill exist there\n")
		}
	}
}

func TestRemoveNS(t *testing.T) {
	n := new(html.Node)

	n.Attr = []html.Attribute{
		{Key: "key-0", Val: "value-0", Namespace: "namespace-0"},
		{Key: "key-0", Val: "value-0", Namespace: "namespace-1"},
		{Key: "key-1", Val: "value-1", Namespace: "namespace-1"},
		{Key: "key-2", Val: "value-2", Namespace: "namespace-1"},
		{Key: "key-0", Val: "value-0", Namespace: "namespace-2"},
		{Key: "key-2", Val: "value-2", Namespace: "namespace-2"},
		{Key: "key-0", Val: "value-0", Namespace: "namespace-3"},
		{Key: "key-1", Val: "value-1", Namespace: "namespace-3"},
		{Key: "key-3", Val: "value-3", Namespace: "namespace-3"},
	}

	expect := len(n.Attr)
	for len(n.Attr) > 0 {
		index := rand.Intn(len(n.Attr))
		ns, key := n.Attr[index].Namespace, n.Attr[index].Key
		RemoveNS(n, ns, key)
		expect--
		actual := len(n.Attr)
		if expect != actual {
			t.Errorf("\ngot : %d, want: %d\n", actual, expect)
			break
		}
		if HasNS(n, ns, key) {
			fmt.Printf("\nstill exist there\n")
		}
	}
}
