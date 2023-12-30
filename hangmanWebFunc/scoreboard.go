package hangmanWeb

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/MounKilian/hangman"
)

func Read(H *hangman.HangManData) {
	f, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading records")
	}

	H.Scoreboard = records
}

func Write(H *hangman.HangManData) {
	file, err := os.Create("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	writer.WriteAll(H.Scoreboard)
	writer.Write(H.NewScore)
}
