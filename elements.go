package gohtml

func HTML() *Element {
	var v = NewElement("html")
	return v
}

func (e *Element) Element(elementType string) *Element {
	var v = NewElement(elementType)
	e.Add(v)
	return v
}

func (e *Element) MetaCharset(v string) *RestrainedElement {
	var meta = NewElement("meta")
	meta.Attr("charset", v)
	meta.NoClose()
	e.Add(meta)
	return (*RestrainedElement)(meta)
}

func (e *Element) MetaTag(name string, content string) *RestrainedElement {
	var meta = NewElement("meta")
	meta.Attr("name", name)
	meta.Attr("content", content)
	meta.NoClose()
	e.Add(meta)
	return (*RestrainedElement)(meta)
}

func (e *Element) A(href string) *Element {
	var v = NewElement("a")
	v.Attr("href", href)
	e.Add(v)
	return v
}
func (e *Element) Abbr() *Element {
	var v = NewElement("abbr")
	e.Add(v)
	return v
}
func (e *Element) Address() *Element {
	var v = NewElement("address")
	e.Add(v)
	return v
}
func (e *Element) Area() *Element {
	var v = NewElement("area")
	e.Add(v)
	return v
}
func (e *Element) Article() *Element {
	var v = NewElement("article")
	e.Add(v)
	return v
}
func (e *Element) Aside() *Element {
	var v = NewElement("aside")
	e.Add(v)
	return v
}
func (e *Element) Audio() *Element {
	var v = NewElement("audio")
	e.Add(v)
	return v
}
func (e *Element) B() *Element {
	var v = NewElement("b")
	e.Add(v)
	return v
}
func (e *Element) Bdi() *Element {
	var v = NewElement("bdi")
	e.Add(v)
	return v
}
func (e *Element) Bdo() *Element {
	var v = NewElement("bdo")
	e.Add(v)
	return v
}
func (e *Element) Blockquote() *Element {
	var v = NewElement("blockquote")
	e.Add(v)
	return v
}
func (e *Element) Body() *Element {
	var v = NewElement("body")
	e.Add(v)
	return v
}
func (e *Element) Br() *Element {
	var v = NewElement("br").NoClose()
	e.Add(v)
	return v
}
func (e *Element) Button() *Element {
	var v = NewElement("button")
	e.Add(v)
	return v
}
func (e *Element) Canvas() *Element {
	var v = NewElement("canvas")
	e.Add(v)
	return v
}
func (e *Element) Caption() *Element {
	var v = NewElement("caption")
	e.Add(v)
	return v
}
func (e *Element) Cite() *Element {
	var v = NewElement("cite")
	e.Add(v)
	return v
}
func (e *Element) Code() *Element {
	var v = NewElement("code")
	e.Add(v)
	return v
}
func (e *Element) Col() *Element {
	var v = NewElement("col")
	e.Add(v)
	return v
}
func (e *Element) Colgroup() *Element {
	var v = NewElement("colgroup")
	e.Add(v)
	return v
}
func (e *Element) Command() *Element {
	var v = NewElement("command")
	e.Add(v)
	return v
}
func (e *Element) Datalist() *Element {
	var v = NewElement("datalist")
	e.Add(v)
	return v
}
func (e *Element) Dd() *Element {
	var v = NewElement("dd")
	e.Add(v)
	return v
}
func (e *Element) Del() *Element {
	var v = NewElement("del")
	e.Add(v)
	return v
}
func (e *Element) Details() *Element {
	var v = NewElement("details")
	e.Add(v)
	return v
}
func (e *Element) Dfn() *Element {
	var v = NewElement("dfn")
	e.Add(v)
	return v
}
func (e *Element) Div() *Element {
	var v = NewElement("div")
	e.Add(v)
	return v
}
func (e *Element) Dl() *Element {
	var v = NewElement("dl")
	e.Add(v)
	return v
}
func (e *Element) Dt() *Element {
	var v = NewElement("dt")
	e.Add(v)
	return v
}
func (e *Element) Em() *Element {
	var v = NewElement("em")
	e.Add(v)
	return v
}
func (e *Element) Embed() *Element {
	var v = NewElement("embed")
	e.Add(v)
	return v
}
func (e *Element) Fieldset() *Element {
	var v = NewElement("fieldset")
	e.Add(v)
	return v
}
func (e *Element) Figcaption() *Element {
	var v = NewElement("figcaption")
	e.Add(v)
	return v
}
func (e *Element) Figure() *Element {
	var v = NewElement("figure")
	e.Add(v)
	return v
}
func (e *Element) Footer() *Element {
	var v = NewElement("footer")
	e.Add(v)
	return v
}
func (e *Element) Form() *Element {
	var v = NewElement("form")
	e.Add(v)
	return v
}
func (e *Element) H1() *Element {
	var v = NewElement("h1")
	e.Add(v)
	return v
}
func (e *Element) H2() *Element {
	var v = NewElement("h2")
	e.Add(v)
	return v
}
func (e *Element) H3() *Element {
	var v = NewElement("h3")
	e.Add(v)
	return v
}
func (e *Element) H4() *Element {
	var v = NewElement("h4")
	e.Add(v)
	return v
}
func (e *Element) H5() *Element {
	var v = NewElement("h5")
	e.Add(v)
	return v
}
func (e *Element) H6() *Element {
	var v = NewElement("h6")
	e.Add(v)
	return v
}
func (e *Element) Header() *Element {
	var v = NewElement("header")
	e.Add(v)
	return v
}
func (e *Element) Hr() *RestrainedElement {
	var v = NewElement("hr").NoClose()
	e.Add(v)
	return (*RestrainedElement)(v)
}
func (e *Element) I() *Element {
	var v = NewElement("i")
	e.Add(v)
	return v
}
func (e *Element) Iframe() *Element {
	var v = NewElement("iframe")
	e.Add(v)
	return v
}
func (e *Element) Img(src string, alt ...string) *Element {
	var v = NewElement("img").NoClose().Attr("src", src)
	if len(alt) != 0 {
		v.Attr("alt", alt[0])
	}
	e.Add(v)
	return v
}

// TYPE -> (NAME, ID) -> PLACEHOLDER -> VALUE
func (e *Element) Input() *RestrainedElement {
	var v = NewElement("input")
	v.NoClose()
	e.Add(v)
	return (*RestrainedElement)(v)
}

func (e *Element) Ins() *Element {
	var v = NewElement("ins")
	e.Add(v)
	return v
}
func (e *Element) Kbd() *Element {
	var v = NewElement("kbd")
	e.Add(v)
	return v
}
func (e *Element) Keygen() *Element {
	var v = NewElement("keygen")
	e.Add(v)
	return v
}
func (e *Element) Label() *Element {
	var v = NewElement("label")
	e.Add(v)
	return v
}
func (e *Element) Legend() *Element {
	var v = NewElement("legend")
	e.Add(v)
	return v
}
func (e *Element) Main() *Element {
	var v = NewElement("main")
	e.Add(v)
	return v
}
func (e *Element) Map() *Element {
	var v = NewElement("map")
	e.Add(v)
	return v
}
func (e *Element) Mark() *Element {
	var v = NewElement("mark")
	e.Add(v)
	return v
}
func (e *Element) Menu() *Element {
	var v = NewElement("menu")
	e.Add(v)
	return v
}
func (e *Element) Meter() *Element {
	var v = NewElement("meter")
	e.Add(v)
	return v
}
func (e *Element) Nav() *Element {
	var v = NewElement("nav")
	e.Add(v)
	return v
}
func (e *Element) Object() *Element {
	var v = NewElement("object")
	e.Add(v)
	return v
}
func (e *Element) Optgroup() *Element {
	var v = NewElement("optgroup")
	e.Add(v)
	return v
}

func (e *Element) Output() *Element {
	var v = NewElement("output")
	e.Add(v)
	return v
}
func (e *Element) P() *Element {
	var v = NewElement("p")
	e.Add(v)
	return v
}
func (e *Element) Param() *Element {
	var v = NewElement("param")
	e.Add(v)
	return v
}
func (e *Element) Pre() *Element {
	var v = NewElement("pre")
	e.Add(v)
	return v
}
func (e *Element) Progress() *Element {
	var v = NewElement("progress")
	e.Add(v)
	return v
}
func (e *Element) Q() *Element {
	var v = NewElement("q")
	e.Add(v)
	return v
}
func (e *Element) Rp() *Element {
	var v = NewElement("rp")
	e.Add(v)
	return v
}
func (e *Element) Rt() *Element {
	var v = NewElement("rt")
	e.Add(v)
	return v
}
func (e *Element) Ruby() *Element {
	var v = NewElement("ruby")
	e.Add(v)
	return v
}
func (e *Element) S() *Element {
	var v = NewElement("s")
	e.Add(v)
	return v
}
func (e *Element) Samp() *Element {
	var v = NewElement("samp")
	e.Add(v)
	return v
}
func (e *Element) Section() *Element {
	var v = NewElement("section")
	e.Add(v)
	return v
}

func (e *Element) Small() *Element {
	var v = NewElement("small")
	e.Add(v)
	return v
}
func (e *Element) Source() *Element {
	var v = NewElement("source")
	e.Add(v)
	return v
}
func (e *Element) Span() *Element {
	var v = NewElement("span")
	e.Add(v)
	return v
}
func (e *Element) Strong() *Element {
	var v = NewElement("strong")
	e.Add(v)
	return v
}
func (e *Element) Sub() *Element {
	var v = NewElement("sub")
	e.Add(v)
	return v
}
func (e *Element) Summary() *Element {
	var v = NewElement("summary")
	e.Add(v)
	return v
}
func (e *Element) Sup() *Element {
	var v = NewElement("sup")
	e.Add(v)
	return v
}
func (e *Element) Textarea() *Element {
	var v = NewElement("textarea")
	e.Add(v)
	return v
}

func (e *Element) Time() *Element {
	var v = NewElement("time")
	e.Add(v)
	return v
}
func (e *Element) Track() *Element {
	var v = NewElement("track")
	e.Add(v)
	return v
}
func (e *Element) U() *Element {
	var v = NewElement("u")
	e.Add(v)
	return v
}
func (e *Element) Var() *Element {
	var v = NewElement("var")
	e.Add(v)
	return v
}
func (e *Element) Video() *Element {
	var v = NewElement("video")
	e.Add(v)
	return v
}
func (e *Element) Wbr() *Element {
	var v = NewElement("wbr")
	e.Add(v)
	return v
}

type HTMLSELECT Element

// Takes the text to be displayed and a list of items to be selected from
func (e *Element) Select(name string) *HTMLSELECT {
	var v = NewElement("select")
	e.Add(v)
	return (*HTMLSELECT)(v)
}

func (e *HTMLSELECT) Option(value, text string) *Element {
	var v = NewElement("option")
	v.Attr("value", value).InnerText = text
	Add(e, v)
	return v
}

type HTMLTABLE Element

func (e *Element) Table() *HTMLTABLE {
	var v = NewElement("table")
	e.Add(v)
	return (*HTMLTABLE)(v)
}

func (e *HTMLTABLE) Tbody() *Element {
	var v = NewElement("tbody")
	Add(e, v)
	return v
}

func (e *HTMLTABLE) Td() *Element {
	var v = NewElement("td")
	Add(e, v)
	return v
}

func (e *HTMLTABLE) Tfoot() *Element {
	var v = NewElement("tfoot")
	Add(e, v)
	return v
}

func (e *HTMLTABLE) Th() *Element {
	var v = NewElement("th")
	Add(e, v)
	return v
}

func (e *HTMLTABLE) Thead() *Element {
	var v = NewElement("thead")
	Add(e, v)
	return v
}

func (e *HTMLTABLE) Tr() *Element {
	var v = NewElement("tr")
	Add(e, v)
	return v
}

type HTMLLIST Element

func (e *Element) Ul() *HTMLLIST {
	var v = NewElement("ul")
	e.Add(v)
	return (*HTMLLIST)(v)
}
func (e *Element) Ol() *HTMLLIST {
	var elem = (*Element)(e)
	var v = NewElement("ol")
	elem.Add(v)
	return (*HTMLLIST)(v)
}
func (e *HTMLLIST) Li() *Element {
	var elem = (*Element)(e)
	var v = NewElement("li")
	elem.Add(v)
	return v
}
