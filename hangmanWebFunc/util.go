package hangmanWeb

import (
	"log"
	"strconv"

	"github.com/MounKilian/hangman"
)

func InitGame(H *hangman.HangManData) {
	if H.Point == 0 {
		H.Loose = 0
		H.Win = 0
		DetectLevel(H)
		H.NewScore = []string{H.Username, strconv.Itoa(H.Point), H.Level, strconv.Itoa(H.Win), strconv.Itoa(H.Loose)}
		Write(H)
		Read(H)
		Refresh(H)
	}
	H.ToFind = hangman.RandomWord(string(("dic/" + H.WordFile)))
	H.Word = hangman.RandomWordUnderscore(H.ToFind)
	H.LetterInput = ""
	H.Attempts = 10
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

func DetectLevel(H *hangman.HangManData) {
	if H.WordFile == "words.txt" {
		H.Level = "easy"
	} else if H.WordFile == "words2.txt" {
		H.Level = "medium"
	} else if H.WordFile == "words3.txt" {
		H.Level = "hard"
	}
}

func DetectFile(H *hangman.HangManData) {
	if H.Level == "easy" {
		H.WordFile = "words.txt"
	} else if H.Level == "medium" {
		H.WordFile = "words2.txt"
	} else if H.Level == "hard" {
		H.WordFile = "words3.txt"
	}
}

func Email(H []string) bool {
	AllAccount := ReadSignIn()
	for _, Account := range AllAccount {
		if Account[1] == H[1] {
			return false
		}
		if Account[0] == H[0] {
			return false
		}
	}
	return true
}

func AcccountUse(Account []string, H *hangman.HangManData) bool {
	AllAccount := ReadSignIn()
	for _, Acc := range AllAccount {
		if Acc[1] == Account[1] && Acc[2] == Account[2] {
			H.Username = Acc[0]
			return true
		}
	}
	return false
}

func Log(H *hangman.HangManData) {
	for _, rec := range H.Scoreboard {
		if H.Username == rec[0] {
			var err error
			H.Point, err = strconv.Atoi(rec[1])
			if err != nil {
				log.Fatal(err)
			}
			H.Win, err = strconv.Atoi(rec[3])
			if err != nil {
				log.Fatal(err)
			}
			H.Loose, err = strconv.Atoi(rec[4])
			if err != nil {
				log.Fatal(err)
			}
			H.Level = rec[2]
			DetectFile(H)
			H.NewScore = []string{H.Username, strconv.Itoa(H.Point), H.Level, strconv.Itoa(H.Win), strconv.Itoa(H.Loose)}
		}
	}
}
