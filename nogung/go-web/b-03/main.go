package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var filePath = path.Join("views", "index.html")
		var tmpl, err = template.ParseFiles(filePath)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var data = map[string]interface{}{
			"title": "learning golang web",
			"name":  "Arnov",
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	var address = "localhost:9000"
	fmt.Println("server started at", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
