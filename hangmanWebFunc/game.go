package hangmanWeb

import (
	"html/template"
	"log"
	"net/http"

	"github.com/MounKilian/hangman"
)

func Form(w http.ResponseWriter, r *http.Request, H *hangman.HangManData) {
	if !hangman.VerifIfAlreadyUse(H) && (H.LetterInput >= "a" && H.LetterInput <= "z") {
		H.Letters += H.LetterInput + " | "
		if len(H.LetterInput) == 1 {
			hangman.Verification(H)
			if hangman.WordFind(H) {
				http.Redirect(w, r, "/win", http.StatusFound)
			}
		} else if len(H.LetterInput) > 1 {
			win := hangman.EnterWord(H)
			if win {
				http.Redirect(w, r, "/win", http.StatusFound)
			}
		}
	}
	if H.Attempts <= 0 {
		http.Redirect(w, r, "/loose", http.StatusFound)
	}
	template, err := template.ParseFiles("./pages/game.html", "./templates/header.html", "./templates/informations.html")
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

func Win(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./pages/win.html", "./templates/header.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func Loose(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./pages/loose.html", "./templates/header.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func Level(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./pages/level.html", "./templates/header.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}
