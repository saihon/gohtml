package saihon

import "golang.org/x/net/html"

// Collection
type Collection struct {
	Nodes []*html.Node
}

// Length
func (e Collection) Length() int {
	return len(e.Nodes)
}

// Get returns the "*Element" given index
func (e Collection) Get(index int) *Element {
	if len(e.Nodes) > 0 && index < len(e.Nodes) {
		return &Element{e.Nodes[index]}
	}
	return nil
}

// Enumerator can calls with for..range
// for element := range elements.Enumerator()...
func (e Collection) Enumerator() chan *Element {
	ch := make(chan *Element, 1)
	go e.enumeration(ch)
	return ch
}

func (e Collection) enumeration(ch chan *Element) {
	for i := 0; i < len(e.Nodes); i++ {
		ch <- &Element{e.Nodes[i]}
	}
	close(ch)
}

type ForEachFunc = func(value *Element, index int, collection Collection)

func (e Collection) ForEach(fn ForEachFunc) {
	for i := 0; i < len(e.Nodes); i++ {
		fn(&Element{e.Nodes[i]}, i, Collection{Nodes: e.Nodes})
	}
}
