package hangmanWeb

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

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

func Update(H *hangman.HangManData) {
	for _, record := range H.Scoreboard {
		if record[0] == H.NewScore[0] && record[2] == H.NewScore[2] {
			record[1] = strconv.Itoa(H.Point)
		}
	}

	sort.Slice(H.Scoreboard, func(i, j int) bool {
		return H.Scoreboard[i][1] > H.Scoreboard[j][1]
	})

	file, err := os.Create("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	writer.WriteAll(H.Scoreboard)
}
