package gohtml

type ElementConstraint interface {
	*Element | *HTMLLIST | *HTMLSELECT | *HTMLTABLE | *RestrainedElement
}

type RestrainedElement Element

// This method is used to set an attribute on an element. It returns the element for inlining purposes.
func (e *RestrainedElement) Attr(k string, v ...string) *RestrainedElement {
	return (*RestrainedElement)(
		(*Element)(e).Attr(k, v...),
	)
}

// This method is used to set a boolean attribute on an element. It returns the element for inlining purposes.
func (e *RestrainedElement) AttrBool(k string, v bool) *RestrainedElement {
	return (*RestrainedElement)(
		(*Element)(e).AttrBool(k, v),
	)
}

// This method is used to set an attribute on an element. It returns the element for inlining purposes.
func (e *RestrainedElement) AttrSemicol(k string, v ...string) *RestrainedElement {
	return (*RestrainedElement)(
		(*Element)(e).AttrSemiCol(k, v...),
	)
}

func Restrain[T ElementConstraint](e T) *RestrainedElement {
	return (*RestrainedElement)(e)
}

func UnRestrain[T ElementConstraint](e T) *Element {
	return (*Element)(e)
}
