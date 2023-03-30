package gohtml

import (
	"strings"
	"sync"
	"unsafe"
)

// identifier is a unique identifier for an element.
//
// This is used to compare elements to each other.
//
// It is the address of the element in memory.
func (e *Element) identifier() uintptr {
	return uintptr(unsafe.Pointer(e))
}

// Get an element by another element!
func (e *Element) GetElementByElement(elem *Element) *Element {
	var elems []*Element = getElementsByAttr(e, "", elem.identifier(), true, identifierCmpFunc)
	if len(elems) > 0 {
		return elems[0]
	}
	return nil
}

// NextChild returns the next child element of the current elements parent.
// If the current element is the last child, nil is returned.
func (e *Element) NextChild() *Element {
	var parentChildren, index = indexSelf(e)
	if index+1 >= len(e.parent.Children) || index == -1 {
		return nil
	}
	return parentChildren[index+1]
}

// PrevChild returns the previous child element of the current elements parent.
// If the current element is the first child, nil is returned.
func (e *Element) PrevChild() *Element {
	var parentChildren, index = indexSelf(e)
	if index == 0 || index == -1 {
		return nil
	}
	return parentChildren[index-1]
}

// indexSelf returns the index of the current element in its parent's children slice.
// If the current element has no parent, -1 is returned.
func indexSelf(e *Element) ([]*Element, int) {
	if e.parent == nil {
		return nil, -1
	}
	var index int
	for i, child := range e.parent.Children {
		if child.identifier() == e.identifier() {
			index = i
			break
		}
	}
	return e.parent.Children, index
}

// ID sets the id attribute of the element.
// If the element already has an id attribute, it will be overwritten.
// If the id is an empty string, the id attribute will be removed.
func (e *Element) ID(id string) *Element {
	if id == "" {
		e.DelAttr("id")
		return e
	}
	e.Attr("id", id)
	return e
}

// Class sets the class attribute of the element.
// If the element already has a class attribute, it will be overwritten.
// If the class is an empty string, the class attribute will be removed.
func (e *Element) Class(class ...string) *Element {
	if len(class) == 0 {
		e.DelAttr("class")
		return e
	}
	e.Attr("class", class...)
	return e
}

// Get inner elements by tag name
func (e *Element) GetElementsByTagName(tagName string) []*Element {
	return getElementsByAttr(e, "type", tagName, false, tagCmpFunc)
}

// Get inner elements by classname
func (e *Element) GetElementsByClassName(className string) []*Element {
	return getElementsByAttr(e, "class", className, false, attrCmpFunc)
}

// Get inner elements by ID
func (e *Element) GetElementByID(id string) *Element {
	var elems []*Element = getElementsByAttr(e, "id", id, true, attrCmpFunc)
	if len(elems) > 0 {
		return elems[0]
	}
	return nil
}

func getElementsByAttr[T comparable](e *Element, fieldname string, cmp T, stopAfterFound bool, cmpFunc func(*Element, string, T) bool) []*Element {
	var elems []*Element = make([]*Element, 0)
	var mu = sync.Mutex{}
	e.AsyncForEach(func(e *Element) bool {
		if cmpFunc(e, fieldname, cmp) {
			mu.Lock()
			defer mu.Unlock()
			elems = append(elems, e)
			return stopAfterFound
		}
		return false
	})
	return elems
}

func attrCmpFunc(e *Element, fieldname, cmp string) bool {
	var attrs, _ = getAttr(e.Attrs, fieldname)
	for _, c := range attrs {
		if c == cmp {
			return true
		}
	}
	return false
}

func tagCmpFunc(e *Element, fieldname, cmp string) bool {
	return strings.EqualFold(e.Type, cmp)
}

func identifierCmpFunc(e *Element, fieldname string, cmp uintptr) bool {
	return e.identifier() == cmp
}
