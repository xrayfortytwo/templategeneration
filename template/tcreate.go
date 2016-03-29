package template

/*
   Helper for template generation execution
*/

import (
	"os"
	"text/template"
)

type in struct {
	Token string
	Input string
}

func CreateTemplate() {
	template1 := in{Token: "name_1", Input: "input_1"}
	t := template.New("template")
	t, err := t.ParseFiles("template/template.templ")
	check(err)
	err = t.ExecuteTemplate(os.Stdout, "rule", template1)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
