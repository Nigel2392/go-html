package gohtml

import (
	"fmt"
	"io"
)

// Default document type declaration
const DOCTYPE string = "<!DOCTYPE html>\r\n"

// Document represents an HTML document
//
// It is a wrapper around a predefined HTML() element.
type Document struct {
	root *Element
}

// NewDocument creates a new HTML document
func NewDocument() *Document {
	return &Document{
		root: HTML(),
	}
}

// Len returns the length of the document
func (d *Document) Len() int {
	return d.root.bufLen + len(DOCTYPE)
}

// String returns the document as a string
func (d *Document) String() string {
	return d.Buffer().String()
}

// Bytes returns the document as a byte slice
func (d *Document) Bytes() []byte {
	return d.Buffer().Bytes()
}

// Buffer returns the document as an elementBuffer
func (d *Document) Buffer() *elementBuffer {
	var b *elementBuffer = NewBuffer(d.root.bufLen + len(DOCTYPE))
	var bufLen int = d.root.bufLen + len(DOCTYPE)
	b.WriteString(DOCTYPE)
	d.root.render(b)
	if b.Len() != bufLen {
		panic(fmt.Sprintf("Document: Buffer length mismatch: %d != %d", b.Len(), bufLen))
	}
	return b
}

// Render writes the document to the given io.Writer
func (d *Document) Render(w io.Writer) error {
	var b *elementBuffer = d.Buffer()
	_, err := b.WriteTo(w)
	return err
}

// Get the document's root element.
func (d *Document) Root() *Element {
	return d.root
}

// Get the document's head element. If it does not exist, create it.
func (d *Document) Head() *Element {
	if len(d.root.Children) < 1 {
		var head *Element = NewElement("head")
		d.root.Add(head)
		return head
	}
	return d.root.Children[0]
}

// Get the document's body element. If it does not exist, create it.
func (d *Document) Body() *Element {
	if len(d.root.Children) < 2 {
		var body *Element = NewElement("body")
		d.root.Add(body)
		return body
	}
	return d.root.Children[1]
}

// Get an element by its ID.
func (d *Document) GetElementByID(id string) *Element {
	return d.root.GetElementByID(id)
}

// Get elements by their tag name.
func (d *Document) GetElementsByTagName(tagName string) []*Element {
	return d.root.GetElementsByTagName(tagName)
}

// Get elements by their class name.
func (d *Document) GetElementsByClassName(className string) []*Element {
	return d.root.GetElementsByClassName(className)
}
