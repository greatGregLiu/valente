package elements

import "html"

//Button element produces an HTML button
type Button struct {
	Base
	Text       string
	HTMLEncode bool
	PostBack   []string
}

//String return a string with button
func (btn Button) String() string {
	ret := "<button" + btn.Attrs()
	if len(btn.PostBack) > 0 {
		ret += " onclick=\"sendEvent(" + params(btn.PostBack) + ")\""
	}
	ret += ">"
	if btn.HTMLEncode {
		ret += html.EscapeString(btn.Text)
	} else {
		ret += btn.Text
	}
	ret += "</button>"
	return ret
}

//TextArea element produces an HTML textarea
type TextArea struct {
	Base
	Text string
}

//String return the string to textarea tag
func (ta TextArea) String() string {
	ret := "<textarea" + ta.Attrs() + ">"
	ret += html.EscapeString(ta.Text)
	ret += "</textarea>"
	return ret
}
