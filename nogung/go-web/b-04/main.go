package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type M map[string]interface{}

func main() {
	// // with template.ParseGlob().
	// var tmpl, err = template.ParseGlob("views/*")
	// if err != nil {
	// 	panic(err.Error())
	// 	return
	// }

	// http.HandleFunc("/index", func(rw http.ResponseWriter, r *http.Request) {
	// 	var data = M{"name": "batman"}
	// 	err = tmpl.ExecuteTemplate(rw, "index", data)
	// 	if err != nil {
	// 		http.Error(rw, err.Error(), http.StatusInternalServerError)
	// 	}
	// })

	// with template.ParseFiles
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Batman"}
		var tmpl = template.Must(template.ParseFiles(
			"views/index.html",
			"views/_header.html",
			"views/_message.html",
		))

		var err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// http.HandleFunc("/about", func(rw http.ResponseWriter, r *http.Request) {
	// 	var data = M{"name": "batman"}
	// 	err = tmpl.ExecuteTemplate(rw, "about", data)
	// 	if err != nil {
	// 		http.Error(rw, err.Error(), http.StatusInternalServerError)
	// 	}
	// })

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Batman"}
		var tmpl = template.Must(template.ParseFiles(
			"views/about.html",
			"views/_header.html",
			"views/_message.html",
		))

		var err = tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)

}
