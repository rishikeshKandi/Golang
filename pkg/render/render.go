package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, html string) {

	parsedTemplate, _ := template.ParseFiles("./temp/" + html)

	err := parsedTemplate.Execute(w, nil)

	if err != nil {

		fmt.Println("error msg", err)

		return

	}
}
