package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/harsh-m-patil/GeoSilhouette/internals/types"
)

func repl() {
	reader := bufio.NewScanner(os.Stdin)
	scoreCard := types.NewScoreCard()
	for {
		country := countryMaps()

		data, err := os.ReadFile("./ascii-countries/40-wide/" + country + ".txt")
		if err != nil {
			fmt.Printf("Error reading file %v", err)
		}

		fmt.Print(string(data))

		fmt.Printf("\nGeoSilhoutte (%d) > ", scoreCard.CurrentScore)
		reader.Scan()
		words := cleanInput(reader.Text())

		if len(words) == 0 || words[0] != country {
			scoreCard.CurrentScore -= 10
			fmt.Printf("You lost!! It was %s\n\n", country)

			if scoreCard.CurrentScore == 0 {
				exitAndPrintScore(scoreCard)
			}

			continue
		} else {
			scoreCard.CurrentScore += 10
			fmt.Println("You won")
		}

		if words[0] == "exit" {
			exitAndPrintScore(scoreCard)
		}
	}
}

func exitAndPrintScore(scoreCard *types.ScoreCard) {
	fmt.Printf("Your final score is %d\n", scoreCard.CurrentScore)
	exit()
}
