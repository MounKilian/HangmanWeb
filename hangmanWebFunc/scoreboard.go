package hangmanWeb

import (
	"encoding/csv"
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
		log.Fatal(err)
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

	file, err := os.Create("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	writer.WriteAll(H.Scoreboard)
}

func Refresh(H *hangman.HangManData) {
	sort.Slice(H.Scoreboard, func(i, j int) bool {
		value1, err1 := strconv.Atoi(H.Scoreboard[i][1])
		value2, err2 := strconv.Atoi(H.Scoreboard[j][1])

		if err1 != nil || err2 != nil {
			log.Fatal(err1, err2)
		}

		return value1 > value2
	})

	file, err := os.Create("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	for index, record := range H.Scoreboard {
		if index < 10 {
			writer.Write(record)
		}
	}
}
