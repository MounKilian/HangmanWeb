package hangmanWeb

import (
	"net/http"

	"github.com/MounKilian/hangman"
)

func Server() {
	H := hangman.New("words.txt", "default")
	hangman.FirstLetter(H)
	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		Form(w, r, H)
	})
	http.HandleFunc("/", Menu)
	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		GameBack(w, r, H)
	})
	http.HandleFunc("/help", Help)
	http.HandleFunc("/win", Win)
	http.HandleFunc("/loose", Loose)
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}
