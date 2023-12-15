package hangmanWeb

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/MounKilian/hangman"
)

type Informations struct {
	Email string
}

func Home(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./index.html", "./templates/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func Test(w http.ResponseWriter, r *http.Request, infos *Informations) {
	infos.Email = r.FormValue("Text input")
	fmt.Println(infos.Email)
	http.Redirect(w, r, "http://localhost:8080/", http.StatusFound)
}

func Server() {
	var infos Informations
	fmt.Print(hangman.RandomWordUnderscore("portemanteau"))
	http.HandleFunc("/", Home)
	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		Test(w, r, &infos)
	})
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}
