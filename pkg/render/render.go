package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, html string) {

	parsedTemplates, _ := template.ParseFiles("./temp/" + html)
	err := parsedTemplates.Execute(w, nil)

	if err != nil {
		fmt.Println("error message:", err)
		return
	}
}