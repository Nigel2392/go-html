package gohtml

import (
	"fmt"
	"strings"
	"sync"
)

// Add adds a child to the element.
func Add[T1, T2 ElementConstraint](elem T1, children ...T2) *Element {
	var e = (*Element)(elem)
	var childrenLen = len(e.Children)
	var newChildren = make([]*Element, len(children)+childrenLen)
	copy(newChildren, e.Children)
	for i, child := range children {
		c := (*Element)(child)
		c.parent = e
		c.depth = e.depth + 1
		c.bufLen += len("\t") * c.depth
		if !c.AutoClose && c.Type != "" {
			c.bufLen += (len("\t") * c.depth)
		}
		e.increaseBufLen(c.bufLen)
		newChildren[i+childrenLen] = c
	}
	e.Children = newChildren
	return e
}

// Element is a single HTML element.
type Element struct {
	Type      string     // The type of element, e.g. "div", "span", "input", etc.
	InnerText string     // The inner text of the element, if any.
	AutoClose bool       // Whether or not the element is self closing.
	Children  []*Element // List of child elements.

	BoolAttrs    []Attr[bool]     // List of boolean attributes.
	SemicolAttrs []Attr[[]string] // List of attributes that are separated by semicolons.
	Attrs        []Attr[[]string] // List of attributes that are separated by spaces.

	depth      int                            // Depth of the element in the tree.
	bufLen     int                            // Length of the element's buffer.
	parent     *Element                       // The parent of the element.
	renderFunc func(*Element, *elementBuffer) // The render function for the element.
}

// NewElement creates a new element.
func newElement[T ElementConstraint](t string) *Element {
	var e = &Element{
		Type:         t,
		BoolAttrs:    make([]Attr[bool], 0),
		Attrs:        make([]Attr[[]string], 0),
		SemicolAttrs: make([]Attr[[]string], 0),
		Children:     make([]*Element, 0),
		renderFunc:   renderElement,
		InnerText:    "",
	}
	e.bufLen = len(t) * 2
	e.bufLen += len("<>\n</>\n")
	return e
}

// NewElement creates a new element.
func NewElement(t string) *Element {
	return newElement[*Element](t)
}

// Creates a new Element with the specified text as its inner text.
func textElement[T ElementConstraint](t string, parent T) *RestrainedElement {
	var e = &Element{
		InnerText:    t,
		BoolAttrs:    make([]Attr[bool], 0),
		Attrs:        make([]Attr[[]string], 0),
		SemicolAttrs: make([]Attr[[]string], 0),
		Children:     make([]*Element, 0),
		renderFunc:   renderText,
		parent:       (*Element)(parent),
	}

	// If the parent is not nil, set the parent and depth of the element.
	if parent != nil {
		e.parent = (*Element)(parent)
		e.depth = e.parent.depth + 1
	}

	// Set the buffer length to the length of the text plus one for the newline.
	e.bufLen = len(t) + len("\n")

	return (*RestrainedElement)(e)
}

// https://github.com/golang/go/blob/master/src/html/escape.go
var htmlEscaper = strings.NewReplacer(
	`&`, "&amp;",
	`'`, "&#39;", // "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
	`<`, "&lt;",
	`>`, "&gt;",
	`"`, "&#34;", // "&#34;" is shorter than "&quot;".
)

// Write text to the element.
//
// Returns the element on which this method is called.
func (e *Element) Text(s string) *Element {
	return e.HTML(htmlEscaper.Replace(s))
}

// Write text to the element using a format string.
//
// Returns the element on which this method is called.
func (e *Element) Textf(format string, a ...interface{}) *Element {
	return e.Text(fmt.Sprintf(format, a...))
}

// Writes raw HTML to the element.
//
// Returns the element on which this method is called.
func (e *Element) HTML(s string) *Element {
	var v = textElement(s, e)
	Add(e, v)
	return e
}

func renderText(e *Element, buf *elementBuffer) {
	if e.InnerText != "" {
		for i := 0; i < e.depth; i++ {
			buf.WriteString("\t")
		}
		buf.WriteString(e.InnerText)
		buf.WriteString("\n")
	}
}

func renderElement(e *Element, buf *elementBuffer) {
	for i := 0; i < e.depth; i++ {
		buf.WriteString("\t")
	}
	buf.WriteString("<")
	buf.WriteString(e.Type)
	for _, attr := range e.BoolAttrs {
		attr.ToBuffer(buf)
	}
	for _, attr := range e.Attrs {
		attr.ToBuffer(buf)
	}
	for _, attr := range e.SemicolAttrs {
		attr.ToBuffer(buf)
	}
	if e.AutoClose {
		buf.WriteString("/>")
	} else {
		buf.WriteString(">")
		buf.WriteString("\n")

		for _, child := range e.Children {
			child.render(buf)
		}

		for i := 0; i < e.depth; i++ {
			buf.WriteString("\t")
		}

		buf.WriteString("</")
		buf.WriteString(e.Type)
		buf.WriteString(">")
	}
	buf.WriteString("\n")
}

// NoClose sets the element to not close.
func (e *Element) NoClose() *Element {
	if !e.AutoClose {
		e.decreaseBufLen(len("<>\n") + len(e.Type))
		e.AutoClose = true
	}
	return e
}

func (e *Element) increaseBufLen(l int) {
	e.bufLen += l
	if e.parent != nil {
		e.parent.increaseBufLen(l)
	}
}

func (e *Element) decreaseBufLen(l int) {
	e.bufLen -= l
	if e.parent != nil {
		e.parent.decreaseBufLen(l)
	}
}

// AttrBool adds a boolean attribute to the element.
func (e *Element) AttrBool(k string, v bool) *Element {
	e.BoolAttrs = append(e.BoolAttrs, NewAttr(k, v))
	e.increaseBufLen(len(k) + len("=\"\""))
	if v {
		e.increaseBufLen(len("true"))
	} else {
		e.increaseBufLen(len("false"))
	}
	return e
}

// Attr adds an attribute to the element.
func (e *Element) Attr(k string, v ...string) *Element {
	for i := range e.Attrs {
		if strings.EqualFold(e.Attrs[i].Key, k) {
			e.Attrs[i].Value = append(e.Attrs[i].Value, v...)
			var totalLen int
			for _, v := range v {
				totalLen += len(v) + 1
			}
			e.increaseBufLen(totalLen)
			return e
		}
	}
	var a = NewStringAttr(k, v, " ")
	e.Attrs = append(e.Attrs, a)
	e.increaseBufLen(attrStringLen(a))

	return e
}

// AttrSemiCol adds an attribute to the element.
func (e *Element) AttrSemiCol(k string, v ...string) *Element {
	for i := range e.SemicolAttrs {
		if strings.EqualFold(e.SemicolAttrs[i].Key, k) {
			e.SemicolAttrs[i].Value = append(e.SemicolAttrs[i].Value, v...)
			var totalLen int
			for _, v := range v {
				totalLen += len(v) + 1
			}
			e.increaseBufLen(totalLen)
			return e
		}
	}
	var a = NewStringAttr(k, v, ";")
	e.SemicolAttrs = append(e.SemicolAttrs, a)
	e.increaseBufLen(attrStringLen(a))
	return e
}

// DelAttr removes an attribute from the element.
func (e *Element) DelAttr(k string) *Element {
	e.BoolAttrs = removeListItem(e.BoolAttrs, func(a, b Attr[bool]) bool {
		return strings.EqualFold(a.Key, k)
	})
	e.Attrs = removeListItem(e.Attrs, func(a, b Attr[[]string]) bool {
		return strings.EqualFold(a.Key, k)
	})
	e.SemicolAttrs = removeListItem(e.SemicolAttrs, func(a, b Attr[[]string]) bool {
		return strings.EqualFold(a.Key, k)
	})
	return e
}

func getAttr[T any](list []Attr[T], k string) (T, bool) {
	for _, attr := range list {
		if strings.EqualFold(attr.Key, k) {
			return attr.Value, true
		}
	}
	return *new(T), false
}

// GetAttr returns the attribute value.
func (e *Element) GetAttr(k string) (stringsAttr []string, boolAttr, found bool) {
	var wg = &sync.WaitGroup{}
	var mu = &sync.Mutex{}
	wg.Add(3)
	go func(w *sync.WaitGroup, m *sync.Mutex) {
		defer deferredDone(w, m)
		m.Lock()
		boolAttr, found = getAttr(e.BoolAttrs, k)
	}(wg, mu)
	go func(w *sync.WaitGroup, m *sync.Mutex) {
		defer deferredDone(w, m)
		m.Lock()
		stringsAttr, found = getAttr(e.Attrs, k)
	}(wg, mu)
	go func(w *sync.WaitGroup, m *sync.Mutex) {
		defer deferredDone(w, m)
		m.Lock()
		var attr, ok = getAttr(e.SemicolAttrs, k)
		if ok {
			stringsAttr = attr
		}
	}(wg, mu)
	wg.Wait()
	return stringsAttr, boolAttr, found
}

func deferredDone(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Unlock()
	wg.Done()
}

// removeListItem removes an item from a list of attributes.
func removeListItem[T any](list []T, cond func(a, b T) bool) []T {
	var deleted int
	for i := range list {
		var j = i - deleted
		if cond(list[j], list[i]) {
			list = append(list[:j], list[j+1:]...)
			deleted++
		}
	}
	return list
}

// AddChild adds a child to the element.
func (e *Element) Add(children ...*Element) *Element {
	return Add(e, children...)
}

// String returns the HTML string.
func (e *Element) String() string {
	var buf *elementBuffer = NewBuffer(e.bufLen)
	e.render(buf)
	return buf.String()
}

func (e *Element) render(buf *elementBuffer) {
	e.renderFunc(e, buf)
}

func (e *Element) For(start, iterations, step int, f func(int) *Element) *Element {
	for i := start; i < iterations; i += step {
		e.Add(f(i))
	}
	return e
}

func ForEach[T ElementConstraint](s T, f func(int, T)) {
	f(0, s)
	var sElem = (*Element)(s)
	var children = sElem.Children
	for i := range children {
		ForEach((T)(children[i]), f)
	}
}

// Loop over all inner elements recursively.
func (e *Element) ForEach(f func(int, *Element)) *Element {
	ForEach(e, f)
	return e

}

// Loop over all inner elements recursively asynchronously.
func (e *Element) AsyncForEach(fn func(*Element) bool) {
	var wg sync.WaitGroup
	var terminate = make(chan struct{}, 1)
	wg.Add(1)
	go e.asyncForEach(fn, &wg, terminate)
	wg.Wait()
	for len(terminate) > 0 {
		<-terminate
	}
	close(terminate)
}

func (e *Element) asyncForEach(fn func(*Element) bool, wg *sync.WaitGroup, terminate chan struct{}) {
	defer wg.Done()
	if fn(e) {
		terminate <- struct{}{}
		return
	}
	select {
	case <-terminate:
		return
	default:
		wg.Add(len(e.Children))
		for _, child := range e.Children {
			go child.asyncForEach(fn, wg, terminate)
		}
	}
}
