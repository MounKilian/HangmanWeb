package hangmanWeb

import (
	"html/template"
	"log"
	"net/http"

	"github.com/MounKilian/hangman"
)

func Form(w http.ResponseWriter, r *http.Request, H *hangman.HangManData) {
	if GameLoop(H) == 1 {
		H.Point += 1
		Update(H)
		Read(H)
		http.Redirect(w, r, "/win", http.StatusFound)
	} else if GameLoop(H) == 0 {
		http.Redirect(w, r, "/loose", http.StatusFound)
	}
	template, err := template.ParseFiles("./pages/game.html", "./templates/informations.html", "./templates/stickman.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, H)
}

func GameBack(w http.ResponseWriter, r *http.Request, H *hangman.HangManData) {
	H.LetterInput = r.FormValue("Text input")
	http.Redirect(w, r, "/game", http.StatusFound)
}

func Menu(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./index.html", "./templates/header.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func Help(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./pages/help.html", "./templates/header.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func Win(w http.ResponseWriter, r *http.Request, H *hangman.HangManData) {
	template, err := template.ParseFiles("./pages/win.html", "./templates/header.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, H)
}

func Loose(w http.ResponseWriter, r *http.Request, H *hangman.HangManData) {
	template, err := template.ParseFiles("./pages/loose.html", "./templates/header.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, H)
}

func Level(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./pages/level.html", "./templates/header.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func EasyGame(w http.ResponseWriter, r *http.Request, H *hangman.HangManData) {
	H.WordFile = "words.txt"
	InitGame(H)
	http.Redirect(w, r, "/game", http.StatusFound)
}

func MediumGame(w http.ResponseWriter, r *http.Request, H *hangman.HangManData) {
	H.WordFile = "words2.txt"
	InitGame(H)
	http.Redirect(w, r, "/game", http.StatusFound)
}

func HardGame(w http.ResponseWriter, r *http.Request, H *hangman.HangManData) {
	H.WordFile = "words3.txt"
	InitGame(H)
	http.Redirect(w, r, "/game", http.StatusFound)
}

func Login(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./pages/login.html", "./templates/header.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func Scoreboard(w http.ResponseWriter, r *http.Request, H *hangman.HangManData) {
	template, err := template.ParseFiles("./pages/scoreboard.html", "./templates/header.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, H)
}

func Username(w http.ResponseWriter, r *http.Request, H *hangman.HangManData) {
	H.Username = r.FormValue("User")
	H.Point = 0
	http.Redirect(w, r, "/level", http.StatusFound)
}
