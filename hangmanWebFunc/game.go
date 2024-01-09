package hangmanWeb

import (
	"html/template"
	"log"
	"net/http"

	"github.com/MounKilian/hangman"
)

func Form(w http.ResponseWriter, r *http.Request, H *hangman.HangManData) {
	if GameLoop(H) == 1 {
		if H.Level == "easy" {
			H.Point += 1 * H.Attempts
			H.Win += 1
		} else if H.Level == "medium" {
			H.Point += 2 * H.Attempts
			H.Win += 1
		} else {
			H.Point += 3 * H.Attempts
			H.Win += 1
		}
		Update(H)
		Read(H)
		Refresh(H)
		http.Redirect(w, r, "/win", http.StatusFound)
	} else if GameLoop(H) == 0 {
		H.Loose += 1
		Update(H)
		Read(H)
		Refresh(H)
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
	template, err := template.ParseFiles("./pages/win.html", "./templates/ranking.html", "./templates/header.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, H)
}

func Loose(w http.ResponseWriter, r *http.Request, H *hangman.HangManData) {
	template, err := template.ParseFiles("./pages/loose.html", "./templates/ranking.html", "./templates/header.html")
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

func Login(w http.ResponseWriter, r *http.Request, H *hangman.HangManData) {
	template, err := template.ParseFiles("./pages/login.html", "./templates/header.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, H)
}

func Scoreboard(w http.ResponseWriter, r *http.Request, H *hangman.HangManData) {
	Read(H)
	Refresh(H)
	template, err := template.ParseFiles("./pages/scoreboard.html", "./templates/ranking.html", "./templates/header.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, H)
}

func Username(w http.ResponseWriter, r *http.Request, H *hangman.HangManData) {
	H.Username = r.FormValue("SignInUsername")
	H.Password = r.FormValue("SignInPassword")
	H.Email = r.FormValue("SignInEmail")
	H.Point = 0
	Account := []string{H.Username, H.Email, H.Password}
	if Email(Account) {
		AllAccount := ReadSignIn()
		Save(AllAccount, Account)
		http.Redirect(w, r, "/level", http.StatusFound)
	} else {
		H.TypeOfGame = true
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
