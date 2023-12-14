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
				http.Redirect(w, r, "http://localhost:8080/infos", http.StatusFound)
			}
		} else if len(H.LetterInput) > 1 {
			win := hangman.EnterWord(H)
			if win {
				http.Redirect(w, r, "http://localhost:8080/infos", http.StatusFound)
			}
		}
	}
	template, err := template.ParseFiles("./index.html", "./templates/footer.html", "./templates/informations.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, H)
}

func GameBack(w http.ResponseWriter, r *http.Request, H *hangman.HangManData) {
	H.LetterInput = r.FormValue("Text input")
	http.Redirect(w, r, "http://localhost:8080/", http.StatusFound)
}
