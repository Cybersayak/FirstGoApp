package render

import (
	"fmt"
	"html/template"
	"net/http"
)

//If function starts with lower case letter like renderTemplate it is a
//private function , in case of captal RenderTemplate it is a public function

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	// ,"./templates/base.layout.tmpl"
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error Parsing template: ", err)
		return
	}
}
