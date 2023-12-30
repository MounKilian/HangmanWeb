package hangmanWeb

import (
	"strconv"

	"github.com/MounKilian/hangman"
)

func InitGame(H *hangman.HangManData) {
	H.ToFind = hangman.RandomWord(string(("dic/" + H.WordFile)))
	H.Word = hangman.RandomWordUnderscore(H.ToFind)
	H.LetterInput = ""
	H.Attempts = 10
	H.Point = 0
	H.NewScore = []string{H.Username, strconv.Itoa(H.Point), H.Level}
	Write(H)
	Read(H)
	hangman.FirstLetter(H)
}

func GameLoop(H *hangman.HangManData) int {
	if !hangman.VerifIfAlreadyUse(H) && (H.LetterInput >= "a" && H.LetterInput <= "z") {
		H.Letters += H.LetterInput + " | "
		if len(H.LetterInput) == 1 {
			hangman.Verification(H)
			if hangman.WordFind(H) {
				return 1
			}
		} else if len(H.LetterInput) > 1 {
			win := hangman.EnterWord(H)
			if win {
				return 1
			}
		}
	}
	if H.Attempts <= 0 {
		return 0
	} else {
		return 3
	}
}
