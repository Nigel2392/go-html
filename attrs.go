package gohtml

import (
	"fmt"
	"strings"
)

// Attr represents an attribute of an element.
type Attr[T any] struct {
	Key        string                        // The key of the attribute.
	Value      T                             // The value of the attribute.
	Delimiter  string                        // The delimiter to use when rendering the value.
	renderFunc func(Attr[T], *elementBuffer) // The function that renders the attribute.
	lenFunc    func(Attr[T]) int             // The function that measures the length of the attribute.
}

// NewBoolAttr creates a new boolean attribute.
func NewBoolAttr(k string, v bool) Attr[bool] {
	return Attr[bool]{
		Key:        strings.ToLower(k),
		Value:      v,
		renderFunc: renderBoolAttr,
		lenFunc:    attrBoolLen,
	}
}

// NewStringAttr creates a new string attribute.
func NewStringAttr(k string, v []string, d string) Attr[[]string] {
	return Attr[[]string]{
		Key:        strings.ToLower(k),
		Value:      v,
		Delimiter:  d,
		renderFunc: renderStringAttr,
		lenFunc:    attrStringLen,
	}
}

func NewAttr[T any](k string, v T) Attr[T] {
	return Attr[T]{Key: strings.ToLower(k), Value: v}
}

func (a Attr[T]) String() string {
	return fmt.Sprintf(" %s=\"%v\"", a.Key, a.Value)
}

func (a Attr[T]) ToBuffer(b *elementBuffer) {
	a.renderFunc(a, b)
}

func renderBoolAttr(a Attr[bool], buf *elementBuffer) {
	if a.Value {
		buf.WriteString(fmt.Sprintf(" %s", a.Key))
	}
}

func renderStringAttr(a Attr[[]string], buf *elementBuffer) {
	if len(a.Value) > 0 {
		buf.WriteString(fmt.Sprintf(" %s=\"%s\"", a.Key, strings.Join(a.Value, a.Delimiter)))
	}
}

func attrBoolLen(a Attr[bool]) int {
	if a.Value {
		return len(a.Key) + 1
	}
	return 0
}

func attrStringLen(a Attr[[]string]) int {
	if len(a.Value) > 0 {
		var totalLen = len(a.Key) + 4
		for i, v := range a.Value {
			totalLen += len(v)
			if i < len(a.Value)-1 {
				totalLen += len(a.Delimiter)
			}
		}
		return totalLen
		// return len(a.Key) + len(strings.Join(a.Value, a.Delimiter)) + 4
	}
	return 0
}
