package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/harsh-m-patil/GeoSilhouette/internal/datatypes"
)

func repl() {
	reader := bufio.NewScanner(os.Stdin)
	scoreCard := datatypes.NewScoreCard()
	for {
		country := countryMaps()

		data, err := os.ReadFile("./ascii-countries/40-wide/" + country + ".txt")
		if err != nil {
			fmt.Printf("Error reading file %v", err)
		}

		fmt.Print(string(data))

		fmt.Printf("\nGeoSilhoutte (%d) > ", scoreCard.LivesLeft)
		reader.Scan()
		words := cleanInput(reader.Text())

		if len(words) == 0 {
			updateScore(false, scoreCard)
			countryName := getCountryName(country)
			fmt.Printf("You lost!! It was %s(%s)\n\n", countryName, country)
			checkLives(scoreCard)
			continue
		}

		switch words[0] {
		case country:
			updateScore(true, scoreCard)
		case "exit":
			exitAndPrintScore(scoreCard)
		default:
			updateScore(false, scoreCard)
			countryName := getCountryName(country)
			fmt.Printf("You lost!! It was %s(%s)\n\n", countryName, country)
			checkLives(scoreCard)
		}

	}
}

func exitAndPrintScore(scoreCard *datatypes.ScoreCard) {
	fmt.Printf("Right Guesses : %d\n", scoreCard.Right)
	fmt.Printf("Wrong Guesses : %d\n", scoreCard.Wrong)
	exit()
}

func updateScore(result bool, scoreCard *datatypes.ScoreCard) {
	if result {
		scoreCard.LivesLeft++
		scoreCard.Right++
		fmt.Println("You won")
	} else {
		scoreCard.LivesLeft--
		scoreCard.Wrong++
	}
}

func checkLives(scoreCard *datatypes.ScoreCard) {
	if scoreCard.LivesLeft == 0 {
		exitAndPrintScore(scoreCard)
	}
}
