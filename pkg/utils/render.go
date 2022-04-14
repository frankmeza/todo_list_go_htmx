package utils

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/thedevsaddam/renderer"
)

var render *renderer.Render

func init() {
	options := renderer.Options{
		ParseGlobPattern: "pkg/templates/*.html",
	}

	render = renderer.New(options)
}

func RenderHtmlTemplate(
	w http.ResponseWriter,
	statusCode int,
	templateName string,
) {
	render.HTML(w, statusCode, templateName, nil)
}

func RenderHome(w http.ResponseWriter, r *http.Request) {
	RenderHtmlTemplate(w, http.StatusOK, "home")
}

// example templatesData map
// map[string]string{
// 	"addInput": "pkg/templates/add_input.html",
// 	"addTask":  "pkg/templates/add_task.html",
// }
func RenderHtmlTemplates(
	w http.ResponseWriter,
	statusCode int,
	templatesData map[string]string,
) {
	var err error
	var templatesToSend *template.Template

	for key, value := range templatesData {
		templatesToSend, err = template.ParseFiles(value)

		if err != nil {
			fmt.Fprintf(w, "cannot find template: %s", err)
		}

		templatesToSend.ExecuteTemplate(w, key, nil)
	}
}

func RenderHtmlString(
	w http.ResponseWriter,
	statusCode int,
	html string,
) {
	render.HTMLString(w, http.StatusOK, html)
}
