package hangmanWeb

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./index.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func Server() {
	http.HandleFunc("/", Home)
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}
