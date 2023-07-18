package examples

import (
	"math/rand"
)

// Picks a random card and return the card as string
func PickTarrotCard() string {
	randomNumber := rand.Intn(8)

	tarrotReading := map[int]string{
		0: "The Fool",
		1: "The Magician",
		2: "The High Priestess",
		3: "The Empress",
		4: "The Emperor",
		5: "The Hierophant",
		6: "The Lovers",
		7: "The Chariot",
	}

	return tarrotReading[randomNumber]
}
