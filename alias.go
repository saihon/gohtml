package gohtml

import (
	"github.com/saihon/gohtml/find"
)

// Deprecated: GetByTag alias `GetElementsByTagName'
func (d Document) GetByTag(tagname string) Collection {
	return Collection{find.ByTag(d.Node, tagname)}
}

// Deprecated: GetByName alias `GetElementsByName'
func (d Document) GetByName(name string) Collection {
	return Collection{find.ByName(d.Node, name)}
}

// Deprecated: GetByClass alias `GetElementsByClassName'
func (d Document) GetByClass(classname string) Collection {
	return Collection{find.ByClass(d.Node, classname)}
}

// Deprecated: QueryAll alias `QuerySelectorAll'
func (d Document) QueryAll(selector string) Collection {
	return Collection{find.QueryAll(d.Node, selector)}
}

// Deprecated: GetById alias `GetElementById'
func (d Document) GetById(id string) *Element {
	if n := find.ById(d.Node, id); n != nil {
		return &Element{n}
	}
	return nil
}

// Deprecated: Query alias `QuerySelector'
func (d Document) Query(selector string) *Element {
	if n := find.Query(d.Node, selector); n != nil {
		return &Element{Node: n}
	}
	return nil
}

// Deprecated: GetByTag alias `GetElementsByTagName'
func (e Element) GetByTag(tagname string) Collection {
	return Collection{find.ByTag(e.Node, tagname)}
}

// Deprecated: GetByName alias `GetElementsByName'
func (e Element) GetByName(name string) Collection {
	return Collection{find.ByName(e.Node, name)}
}

// Deprecated: GetByClass alias `GetElementsByClassName'
func (e Element) GetByClass(classname string) Collection {
	return Collection{find.ByClass(e.Node, classname)}
}

// Deprecated: QueryAll alias `QuerySelectorAll'
func (e Element) QueryAll(selector string) Collection {
	return Collection{find.QueryAll(e.Node, selector)}
}

// Deprecated: GetById alias `GetElementById'
func (e Element) GetById(id string) *Element {
	if n := find.ById(e.Node, id); n != nil {
		return &Element{n}
	}
	return nil
}

// Deprecated: Query alias `QuerySelector'
func (e Element) Query(selector string) *Element {
	if n := find.Query(e.Node, selector); n != nil {
		return &Element{Node: n}
	}
	return nil
}
