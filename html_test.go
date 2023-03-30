package gohtml_test

import (
	"bytes"
	"testing"

	gohtml "github.com/Nigel2392/go-html"
)

func TestHTML(t *testing.T) {
	var document = gohtml.NewDocument()
	var head = document.Head()
	var body = document.Body()

	head.Element("title").Text("Hello World!")
	head.MetaCharset("utf-8")
	head.MetaTag("viewport", "width=device-width, initial-scale=1, shrink-to-fit=no")
	body.Div().Attr("class", "container").Text("Hello World!")
	body.Div().Attr("class", "container").Attr("class", "test").Text("Hello World!")
	body.Div().Attr("class", "container").Attr("class", "test").Text("Hello World!")
	body.Div().Attr("class", "container").Attr("class", "test").Text("Hello World!")
	body.Div().Attr("class", "container").Attr("class", "test").Text("Hello World!")
	body.Div().Div().Div().Div().Input()
	body.Div().Div().Div().Div().Input()
	body.Div().Div().Div().Div().Input().Attr("id", "test")
	body.Div().Div().Div().Div().Input()
	body.Div().Div().Div().Div().Input()
	div := body.Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div()
	div = div.Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div()
	div = div.Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div()
	div = div.Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div()
	div = div.Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div()
	div = div.Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div()
	div = div.Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div()
	div = div.Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div()
	div = div.Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div()
	div = div.Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div()
	div = div.Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div()
	div = div.Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div()
	div = div.Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div()
	div = div.Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div()
	div = div.Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div()
	div = div.Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div()
	div = div.Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div().Div()
	div.Class("TEST", "test")

	var b bytes.Buffer
	document.Render(&b)
	if document.Len() != b.Len() {
		t.Errorf("Expected %d bytes, got %d bytes", document.Len(), b.Len())
	}
	if document.GetElementByID("test") == nil {
		t.Errorf("Expected to find element with id 'test'")
	}
	if len(document.GetElementsByClassName("test")) != 5 {
		t.Errorf("Expected to find 5 elements with class 'test'")
	}
	if len(document.GetElementsByTagName("input")) != 5 {
		t.Errorf("Expected to find 5 elements with tag 'input'")
	}
}

func TestNextPrevChild(t *testing.T) {
	var document = gohtml.NewDocument()
	var body = document.Body()
	var div = body.Div()
	div.Type = "firstDiv"
	var div2 = div.Div()
	div2.Type = "secondDiv"
	var div3 = div2.Div()
	div3.Type = "thirdDiv"
	var div4 = div3.Div()
	div4.Type = "fourthDiv"
	var div5 = div4.Div()
	div5.Type = "fifthDiv"

	if d := div.NextChild(); d != nil && d.Type != div2.Type {
		t.Errorf("Expected div2 to be next sibling of div")
	}
	if d := div2.NextChild(); d != nil && d.Type != div3.Type {
		t.Errorf("Expected div3 to be next sibling of div2")
	}
	if d := div3.NextChild(); d != nil && d.Type != div4.Type {
		t.Errorf("Expected div4 to be next sibling of div3")
	}
	if d := div4.NextChild(); d != nil && d.Type != div5.Type {
		t.Errorf("Expected div5 to be next sibling of div4")
	}
	if d := div5.NextChild(); d != nil {
		t.Errorf("Expected div5 to be last sibling")
	}
	if d := div.PrevChild(); d != nil {
		t.Errorf("Expected div to be first sibling")
	}
	if d := div2.PrevChild(); d != nil && d.Type != div.Type {
		t.Errorf("Expected div to be previous sibling of div2")
	}
	if d := div3.PrevChild(); d != nil && d.Type != div2.Type {
		t.Errorf("Expected div2 to be previous sibling of div3")
	}
	if d := div4.PrevChild(); d != nil && d.Type != div3.Type {
		t.Errorf("Expected div3 to be previous sibling of div4")
	}
	if d := div5.PrevChild(); d != nil && d.Type != div4.Type {
		t.Errorf("Expected div4 to be previous sibling of div5")
	}
}

func BenchmarkHTML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var document = gohtml.NewDocument()
		var head = document.Head()
		var body = document.Body()

		head.Element("title").Text("Hello World!")
		head.MetaCharset("utf-8")
		head.MetaTag("viewport", "width=device-width, initial-scale=1, shrink-to-fit=no")
		body.Div().Attr("class", "container").Text("Hello World!")
		body.Div().Attr("class", "container").Attr("class", "test").Text("Hello World!")
		body.Div().Attr("class", "container").Attr("class", "test").Text("Hello World!")
		body.Div().Attr("class", "container").Attr("class", "test").Text("Hello World!")
		body.Div().Attr("class", "container").Attr("class", "test").Text("Hello World!")
		body.Div().Div().Div().Div().Input()
		body.Div().Div().Div().Div().Input()
		body.Div().Div().Div().Div().Input()
		body.Div().Div().Div().Div().Input()
		body.Div().Div().Div().Div().Input()
		var buf bytes.Buffer
		document.Render(&buf)
		if document.Len() != buf.Len() {
			b.Errorf("Expected %d bytes, got %d bytes", document.Len(), buf.Len())
		} else {
			b.Logf("Got %d bytes", buf.Len())
		}
	}
}

var benchmarkElementCount = []int{
	1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000,
}

func BenchmarkHTMLRender(b *testing.B) {
	for _, count := range benchmarkElementCount {
		var document = gohtml.NewDocument()
		//var head = document.Head()
		//head.Element("title").Text("Hello World!")
		var body = document.Body()
		for i := 0; i < count; i++ {
			body.Div().Attr("class", "container").Text("Hello World!")
		}
		b.Run("Render", func(b *testing.B) {
			startHTMLRenderBenchmark(b, document)
		})
	}
}

func startHTMLRenderBenchmark(b *testing.B, document *gohtml.Document) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		document.Render(&buf)
		if document.Len() != buf.Len() {
			b.Errorf("Expected %d bytes, got %d bytes", document.Len(), buf.Len())
		} else {
			b.Logf("Got %d bytes", buf.Len())
		}
	}
}
