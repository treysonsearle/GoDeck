package main

import (
	"errors"
	"math/rand"
)

var deck []Card
var discardPile []Card
var suits = []string{"spades", "hearts", "clubs", "diamonds"}
var ranks = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}

func init() {
	orderDeck()
}

func shuffleDeck() {
	for i := range deck {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
}

func cutDeck(cutLocation int) error {
	if cutLocation < 0 || cutLocation >= len(deck) {
		return errors.New("Invalid cut location.")
	}
	bottomHalf := deck[cutLocation:]
	deck = append(deck[:cutLocation], bottomHalf...)
	return nil
}

func rebuildDeck() {
	deck = append(deck, discardPile...)
	discardPile = nil // Clear discard pile
	orderDeck()
}

func orderDeck() {
	deck = nil
	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, Card{Suit: suit, Rank: rank})
		}
	}
}
