package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type message struct {
	Message string `json:"message"`
}

func encodeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func decodeJSON(body io.Reader, data interface{}) error {
	return json.NewDecoder(body).Decode(data)
}

func dealCardHandler(w http.ResponseWriter, r *http.Request) {
	card, err := dealCard()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	encodeJSON(w, card)
}
func shuffleHandler(w http.ResponseWriter, r *http.Request) {
	shuffleDeck()
	encodeJSON(w, message{"Deck Shuffled"})
}

func discardHandler(w http.ResponseWriter, r *http.Request) {
	var card Card
	err := decodeJSON(r.Body, &card)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	discardCard(card)
	encodeJSON(w, message{"Card Discarded"})
}

func cutHandler(w http.ResponseWriter, r *http.Request) {
	type cutRequest struct {
		Location int `json:"location"`
	}
	var req cutRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = cutDeck(req.Location)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	encodeJSON(w, message{"Deck Cut"})
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
	orderDeck()
	encodeJSON(w, message{"Deck Ordered"})
}

func rebuildHandler(w http.ResponseWriter, r *http.Request) {
	rebuildDeck()
	encodeJSON(w, message{"Deck Rebuilt"})
}

func cheatHandler(w http.ResponseWriter, r *http.Request) {
	card, err := checkNextCard()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	encodeJSON(w, card)
}

func checkNextCard() (Card, error) {
	if len(deck) == 0 {
		return Card{}, errors.New("Deck is empty. Can't peek.")
	}
	return deck[0], nil
}

func dealCard() (Card, error) {
	if len(deck) == 0 {
		return Card{}, errors.New("Deck is empty. Can't deal cards.")
	}
	dealtCard := deck[0]
	deck = deck[1:]
	return dealtCard, nil
}

func discardCard(card Card) {
	discardPile = append(discardPile, card)
}
