package hangmanWeb

import (
	"net/http"

	"github.com/MounKilian/hangman"
)

func Server() {
	H := hangman.New("words.txt", "default")
	H.TypeOfGame = false
	Read(H)
	Refresh(H)
	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		Form(w, r, H)
	})
	http.HandleFunc("/", Menu)
	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		GameBack(w, r, H)
	})
	http.HandleFunc("/help", Help)
	http.HandleFunc("/win", func(w http.ResponseWriter, r *http.Request) {
		Win(w, r, H)
	})
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
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		Login(w, r, H)
	})
	http.HandleFunc("/username", func(w http.ResponseWriter, r *http.Request) {
		Username(w, r, H)
	})
	http.HandleFunc("/scoreboard", func(w http.ResponseWriter, r *http.Request) {
		Scoreboard(w, r, H)
	})
	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {
		Change(w, r, H)
	})
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}
