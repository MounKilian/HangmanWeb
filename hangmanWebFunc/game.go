package hangmanWeb

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/MounKilian/hangman"
)

func Home(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./index.html", "./templates/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func Test(w http.ResponseWriter, r *http.Request, H *hangman.HangManData) {
	H.LetterInput = r.FormValue("Text input")
	fmt.Println(H.LetterInput)
	http.Redirect(w, r, "http://localhost:8080/", http.StatusFound)
}
