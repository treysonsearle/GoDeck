package main

import (
	"errors"
	"math/rand"
	"sort"
)

var deck []Card
var discardPile []Card
var suits = []string{"spades", "hearts", "clubs", "diamonds"}
var ranks = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}

func init() {
	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, Card{Suit: suit, Rank: rank})
		}
	}
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
	discardPile = nil
	orderDeck()
}

func orderDeck() {
	var newDeck []Card
	for _, suit := range suits {
		for _, rank := range ranks {
			newDeck = append(newDeck, Card{Suit: suit, Rank: rank})
		}
	}
	deck = sortBySuitRank(deck)
}

func sortBySuitRank(deck []Card) []Card {
	if len(deck) == 0 {
		return deck
	}

	sort.Slice(deck, func(i, j int) bool {
		suitRankI := findSuitRank(deck[i].Suit, suits)
		suitRankJ := findSuitRank(deck[j].Suit, suits)
		if suitRankI != suitRankJ {
			return suitRankI < suitRankJ
		}
		rankI := findRank(deck[i].Rank, ranks)
		rankJ := findRank(deck[j].Rank, ranks)
		return rankI < rankJ
	})

	return deck
}

func findSuitRank(suit string, suits []string) int {
	for i, s := range suits {
		if s == suit {
			return i
		}
	}
	return -1
}

func findRank(rank string, ranks []string) int {
	for i, r := range ranks {
		if r == rank {
			return i
		}
	}
	return -1
}
