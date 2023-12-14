package hangmanWeb

import (
	"net/http"

	"github.com/MounKilian/hangman"
)

func Server() {
	H := hangman.New("words.txt", "rien")
	letteruse := ""
	for _, i := range hangman.LettersUse(H) {
		letteruse += i + " | "
	}
	H.Letters = letteruse
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Form(w, r, H)
	})
	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		GameBack(w, r, H)
	})
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}
