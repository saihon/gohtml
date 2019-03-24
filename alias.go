package saihon

import (
	"github.com/saihon/saihon/find"
)

// GetByTag alias `GetElementsByTagName'
func (d Document) GetByTag(tagname string) Collection {
	return Collection{find.ByTag(d.Node, tagname)}
}

// GetByName alias `GetElementsByName'
func (d Document) GetByName(name string) Collection {
	return Collection{find.ByName(d.Node, name)}
}

// GetByClass alias `GetElementsByClassName'
func (d Document) GetByClass(classname string) Collection {
	return Collection{find.ByClass(d.Node, classname)}
}

// QueryAll alias `QuerySelectorAll'
func (d Document) QueryAll(selector string) Collection {
	var c Collection
	c.Nodes = find.QueryAll(d.Node, selector)
	return c
}

// GetById alias `GetElementById'
func (d Document) GetById(id string) *Element {
	if n := find.ById(d.Node, id); n != nil {
		return &Element{n}
	}
	return nil
}

// Query alias `QuerySelector'
func (d Document) Query(selector string) *Element {
	if n := find.Query(d.Node, selector); n != nil {
		return &Element{Node: n}
	}
	return nil
}

// GetByTag alias `GetElementsByTagName'
func (e Element) GetByTag(tagname string) Collection {
	return Collection{find.ByTag(e.Node, tagname)}
}

// GetByName alias `GetElementsByName'
func (e Element) GetByName(name string) Collection {
	return Collection{find.ByName(e.Node, name)}
}

// GetByClass alias `GetElementsByClassName'
func (e Element) GetByClass(classname string) Collection {
	return Collection{find.ByClass(e.Node, classname)}
}

// QueryAll alias `QuerySelectorAll'
func (e Element) QueryAll(selector string) Collection {
	var c Collection
	c.Nodes = find.QueryAll(e.Node, selector)
	return c
}

// GetById alias `GetElementById'
func (e Element) GetById(id string) *Element {
	if n := find.ById(e.Node, id); n != nil {
		return &Element{n}
	}
	return nil
}

// Query alias `QuerySelector'
func (e Element) Query(selector string) *Element {
	if n := find.Query(e.Node, selector); n != nil {
		return &Element{Node: n}
	}
	return nil
}
