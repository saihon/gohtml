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

// Get get Element by index from Collection
func (e Collection) Get(index int) *Element {
	return &Element{e.Nodes[index]}
}

// Enumerator
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
