package renderer

import (
	"fmt"
	"html/template"
	"net/http"
)

const defaultLayout = "layout.html"

type Renderer struct {
	layout string
	w      http.ResponseWriter
}

func New(w http.ResponseWriter) *Renderer {
	return &Renderer{layout: defaultLayout, w: w}
}

func (r *Renderer) Render(templateName string, data interface{}) {
	t, err := template.ParseFiles("templates/layouts/"+r.layout, "templates/"+templateName+".html")
	if nil != err {
		fmt.Println(err)
		return
	}
	t.Execute(r.w, data)
}

func (r *Renderer) Layout(layout string) *Renderer {
	r.layout = layout + ".html"
	return r
}
