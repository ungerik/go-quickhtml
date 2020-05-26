package quickhtml

import (
	"bytes"
	"encoding/json"
	"html"
)

type Renderer struct {
	body bytes.Buffer
}

func (r *Renderer) Bytes() []byte {
	head := []byte("todo")
	return append(head, r.body.Bytes()...)
}

// Write implements io.Writer
func (r *Renderer) Write(data []byte) (int, error) {
	return r.body.Write(data)
}

func (r *Renderer) H1(text string) {
	r.body.WriteString("<h1>")
	r.body.WriteString(html.EscapeString(text))
	r.body.WriteString("</h1>")
}

func (r *Renderer) H2(text string) {
	r.body.WriteString("<h2>")
	r.body.WriteString(html.EscapeString(text))
	r.body.WriteString("</h2>")
}

func (r *Renderer) H3(text string) {
	r.body.WriteString("<h3>")
	r.body.WriteString(html.EscapeString(text))
	r.body.WriteString("</h3>")
}

func (r *Renderer) JSON(jsonBytes []byte) {
	r.body.WriteString("<pre>")
	var pretty bytes.Buffer
	json.Indent(&pretty, jsonBytes, "", "  ")
	r.body.Write(pretty.Bytes())
	r.body.WriteString("</pre>")
}

func (r *Renderer) XML(xmlBytes []byte) {
	r.body.WriteString("<pre>")
	r.body.Write(xmlBytes)
	r.body.WriteString("</pre>")
}

func (r *Renderer) Table(rows [][]string) {
	panic("TODO")
}

func (r *Renderer) StructTable(structs interface{}) {
	panic("TODO")
}
