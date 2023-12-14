package hangmanWeb

import (
	"net/http"

	"github.com/MounKilian/hangman"
)

func Server() {
	H := hangman.New("words.txt", "rien")
	http.HandleFunc("/", Home)
	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		Test(w, r, H)
	})
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}
