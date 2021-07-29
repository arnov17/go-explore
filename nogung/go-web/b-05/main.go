package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Info struct {
	Affiliation string
	Address     string
}

type Person struct {
	Name    string
	Gender  string
	Hobbies []string
	Info    Info
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = Person{
			Name:    "Bruce Wayne",
			Gender:  "male",
			Hobbies: []string{"Reading book", "Traveling", "Buying things"},
			Info:    Info{"Wayne Enterprice", "Gotham City"},
		}

		var tmplt = template.Must(template.ParseFiles("view.html"))
		if err := tmplt.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	var address = "localhost:9000"
	fmt.Printf("server startet at %s", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (t Info) GetAffiliationDetailInfo() string {
	return "have 31 divisions"
}
