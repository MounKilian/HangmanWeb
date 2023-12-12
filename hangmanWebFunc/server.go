package hangmanWeb

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/MounKilian/hangman"
)

type Informations struct {
	Email   string
	Subject string
	Message string
}

func Home(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./index.html", "./templates/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func Infos(w http.ResponseWriter, r *http.Request, infos *Informations) {
	template, err := template.ParseFiles("./pages/infos.html", "./templates/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, infos)
}

func Server() {
	fmt.Print(hangman.RandomWordUnderscore("portemanteau"))
	infos := &Informations{"kilianmoun@gmail.com", "sujet message", "message"}
	http.HandleFunc("/", Home)
	http.HandleFunc("/infos", func(w http.ResponseWriter, r *http.Request) {
		Infos(w, r, infos)
	})
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}
