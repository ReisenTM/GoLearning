package main

import (
	"html/template"
	"net/http"
)

func extend() {
	files, err := template.ParseFiles("templates/base.tmpl", "templates/child.tmpl")
	if err != nil {
		return
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := "帕鲁"
		err := files.ExecuteTemplate(w, "child.tmpl", name)
		if err != nil {
			return
		}

	})
}
