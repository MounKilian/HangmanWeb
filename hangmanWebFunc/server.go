package hangmanWeb

import (
	"net/http"

	"github.com/MounKilian/hangman"
)

func Server() {
	H := hangman.New("words.txt", "default")
	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		Form(w, r, H)
	})
	http.HandleFunc("/", Menu)
	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		GameBack(w, r, H)
	})
	http.HandleFunc("/help", Help)
	http.HandleFunc("/win", Win)
	http.HandleFunc("/level", Level)
	http.HandleFunc("/loose", func(w http.ResponseWriter, r *http.Request) {
		Loose(w, r, H)
	})
	http.HandleFunc("/easygame", func(w http.ResponseWriter, r *http.Request) {
		EasyGame(w, r, H)
	})
	http.HandleFunc("/mediumgame", func(w http.ResponseWriter, r *http.Request) {
		MediumGame(w, r, H)
	})
	http.HandleFunc("/hardgame", func(w http.ResponseWriter, r *http.Request) {
		HardGame(w, r, H)
	})
	http.HandleFunc("/login", Login)
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}
