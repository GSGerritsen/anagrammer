package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

var layoutFuncs = template.FuncMap{
	"yield": func() (string, error) {
		return "", fmt.Errorf("yield used incorrectly")
	},
}

var layout = template.Must(
	template.New("layout.html").Funcs(layoutFuncs).ParseFiles("templates/layout.html"),
)

var templates = template.Must(template.New("t").ParseGlob("templates/**/*.html"))

var errorTemplate = `
<html>
	<body>
		<h1>Error rendering template %s</h1>
		<p>%s</p>
	</body>
</html>
`

func RenderTemplate(w http.ResponseWriter, r *http.Request, templateName string, data map[string]interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}

	funcs := template.FuncMap{
		"yield": func() (template.HTML, error) {
			buf := bytes.NewBuffer(nil)
			err := templates.ExecuteTemplate(buf, templateName, data)
			return template.HTML(buf.String()), err
		},
	}

	layoutClone, _ := layout.Clone()
	layoutClone.Funcs(funcs)
	err := layoutClone.Execute(w, data)

	if err != nil {
		http.Error(w, fmt.Sprintf(errorTemplate, templateName, err), http.StatusInternalServerError)
	}

}
