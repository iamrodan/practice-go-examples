package ifelse

import (
	"fmt"
	"math/rand"
)

func PlayGuessTheNumber() {
	randomNumber := rand.Intn(11)
	var guessedNumber int
	fmt.Println("Guess the number: ")
	fmt.Scanln(&guessedNumber)
	if guessedNumber == randomNumber {
		fmt.Println("You won the game in first go !")
	} else {
		attempts := 1 // user already tried once before
		for guessedNumber != randomNumber {
			if guessedNumber < randomNumber {
				fmt.Println("You have guessed less!")
			} else {
				fmt.Println("You have guessed High!")
			}

			fmt.Scanln(&guessedNumber)
			attempts++
		}
		fmt.Printf("You won the game with %v attempts", attempts)
	}
}
