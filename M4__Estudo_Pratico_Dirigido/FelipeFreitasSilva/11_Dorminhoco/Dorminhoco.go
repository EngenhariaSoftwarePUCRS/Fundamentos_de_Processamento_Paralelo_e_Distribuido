/*
* Felipe Freitas Silva
* 08/09/2023

* 1) Adapte o esqueleto de código para simular o popular jogo de 'dorminhoco'.
* R: Código
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	CARD_AMOUNT = 3
	Ace Suit = iota
	Clubs
	Hearts
	Cups
	Joker = 10
)

var (
	GAME_CAN_START = false
)

type Suit uint8

func (suit Suit) String() string {
	switch suit {
	case Ace:
		return "Espadas"

	case Clubs:
		return "Paus"

	case Hearts:
		return "Ouros"

	case Cups:
		return "Copas"

	case Joker:
		return "Coringa"

	default:
		return "Naipe inválido"
	}
}

type Card struct {
	symbol string
	suit Suit
}

func (card Card) String() string {
	if card.suit == Joker {
		return fmt.Sprintf("%s", card.suit)
	}
	return fmt.Sprintf("%s de %s", card.symbol, card.suit)
}

func canFinish(cards []Card) bool {
	bySuit := false
	previousSuit := cards[0].suit
	for i := 1; i < len(cards); i++ {
		card := cards[i]
		if card.suit != Joker && card.suit != previousSuit {
			bySuit = false
		}
	}
	if bySuit {
		return true
	}

	byValue := false
	previousValue := cards[0].symbol
	for i := 1; i < len(cards); i++ {
		card := cards[i]
		if card.suit != Joker && card.symbol != previousValue {
			byValue = false
		}
	}
	if byValue {
		return true
	}

	return false
}

func jogador(
	id int,
	in <-chan Card,
	out chan<- Card,
	finished chan int,
	cartasIniciais []Card,
) {
	currentHand := cartasIniciais
    var receivedCard Card

	for {
		if !GAME_CAN_START {
			continue
		}
		// Espera de 5-10 segundos
		time.Sleep(time.Duration(rand.Intn(5000) + 5000) * time.Millisecond)
		fmt.Printf("(%d) Vou jogar\n", id)
		fmt.Println("CurrentHand: ", currentHand)
		select {
		case fasterPlayer := <- finished:
			fmt.Printf("%d bateu, eu (%d) vou bater também\n", fasterPlayer, id)
			finished <- id
			
		case receivedCard = <- in:
			fmt.Printf("(%d) Recebi %s\n", id, receivedCard)
			currentHand = append(currentHand, receivedCard)
			fmt.Println("CurrentHand: ", currentHand)
			if canFinish(currentHand) {
				fmt.Printf("(%d) Vou bater\n", id)
				finished <- id
			} else {
				worstCard := currentHand[0]
				fmt.Printf("(%d) Vou passar %s\n", id, worstCard)
				out <- worstCard
			}
		}
	}
}

func insertCard(deck chan<- Card, card Card) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	deck <- card
	// fmt.Printf("Inseri %s\n", card)
}

func getCard(deck <-chan Card) Card {
	return <- deck
}

func main() {
	const PLAYER_AMOUNT = 4

	players := make([]chan Card, PLAYER_AMOUNT)
	for i := 0; i < PLAYER_AMOUNT; i++ {
		players[i] = make(chan Card, CARD_AMOUNT)
	}

	finished := make(chan int, PLAYER_AMOUNT)

	deck := [PLAYER_AMOUNT * CARD_AMOUNT + 1]Card{
		{"12", Ace},
		{"12", Clubs},
		{"12", Hearts},
		{"12", Cups},
		{"11", Ace},
		{"11", Clubs},
		{"11", Hearts},
		{"11", Cups},
		{"10", Ace},
		{"10", Clubs},
		{"10", Hearts},
		{"10", Cups},
		{"@", Joker},
	}

	shuffled_deck := make(chan Card, len(deck))
	for _, card := range deck {
		go insertCard(shuffled_deck, card)
	}
	
	for i := 0; i < PLAYER_AMOUNT; i++ {
		playerCards := make([]Card, CARD_AMOUNT)
		for j := 0; j < CARD_AMOUNT; j++ {
			card := getCard(shuffled_deck)
			// fmt.Printf("Dando %s para o jogador %d\n", card, i)
			playerCards[j] = card
		}
		go jogador(
			i,
			players[i],
			players[(i+1)%PLAYER_AMOUNT],
			finished,
			playerCards,
		)
	}
	// fmt.Println("Última carta: ", getCard(shuffled_deck))
	players[0] <- getCard(shuffled_deck)
	GAME_CAN_START = true
	
	for i := 0; i < PLAYER_AMOUNT - 1; i++ {
		fmt.Printf("%d ficou em %dº lugar\n", <- finished, i+1)
	}

	fmt.Println(<- finished, "ficou em último lugar e leva rolhada!")
}
