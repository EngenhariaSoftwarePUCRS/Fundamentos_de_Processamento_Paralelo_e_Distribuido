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
	Joker = 10
	Ace Suit = iota
	Clubs
	Hearts
	Cups
	// Show comments and explanations during the game
	LOG = true
	CARDS_PER_PLAYER = 3
	MIN_PLAYER_THINKING_TIME = 0.001 * 1000
	MAX_PLAYER_THINKING_TIME = 0.005 * 1000
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
		if card.symbol == "" {
			return "Carta vazia"
		}
		return fmt.Sprint(card.suit)
	}
	return fmt.Sprint(card.symbol, " de ", card.suit)
}

func canFinish(cards []Card) bool {
	suits := make(map[Suit]int)	
	symbols := make(map[string]int)

	for _, card := range cards {
		suits[card.suit]++
		symbols[card.symbol]++
	}
	
	for _, card := range cards {
		if suits[card.suit] == CARDS_PER_PLAYER || symbols[card.symbol] == CARDS_PER_PLAYER {
			return true
		}
	}

	return false
}

func announcePlay(player int, currentHand []Card, message string) {
	// Para facilitar o acompanhamento do jogo, cada jogador leva um tempo para pensar em sua jogada
	time.Sleep(time.Duration(rand.Intn(
		MAX_PLAYER_THINKING_TIME - MIN_PLAYER_THINKING_TIME,
	) + MIN_PLAYER_THINKING_TIME) * time.Millisecond)
	Print(fmt.Sprintf("\n(%d) Vou jogar\n", player))
	Print(fmt.Sprintln("CurrentHand: ", currentHand))
	Print(message)
}

func jogador(
	id int,
	prevPlayer <-chan Card,
	nextPlayer chan<- Card,
	finished chan int,
	startingHand []Card,
) {
	currentHand := startingHand
	haveFinished := false
    var receivedCard, worstCard Card

	for {
		select {
		case fasterPlayer := <- finished: {
			announcePlay(id,
				currentHand,
				fmt.Sprintln(fasterPlayer, " bateu, vou bater também"),
			)
			finished <- id
			haveFinished = true
			// Nesse ponto, o jogador já bateu e está apenas esperando os outros terminarem
			// Normalmente, o jogador seguinte vai esperar uma carta de mim, mas como já bati, vou passar uma "carta" vazia
			// Para simular ele olhando para minhas cartas na mesa
			nextPlayer <- Card{}
			return
		}

		default: {
			receivedCard = <- prevPlayer
			announcePlay(id,
				currentHand,
				fmt.Sprintln("Recebi", receivedCard),
			)
			currentHand = append(currentHand, receivedCard)
			Print(fmt.Sprintln("New CurrentHand: ", currentHand))
			if haveFinished {
				// Nesse ponto, eu já bati e estou apenas esperando os outros terminarem
				// Normalmente, o jogador seguinte vai esperar uma carta de mim, mas como já bati, vou passar uma "carta" vazia
				// Para simular ele olhando para minhas cartas na mesa
				nextPlayer <- Card{}
			} else if canFinish(currentHand) || receivedCard.symbol == "" {
				// Ou posso bater
				// Ou o jogador anterior bateu, e só vi agora
				Print(fmt.Sprintln("Vou bater"))
				finished <- id
				haveFinished = true
				// Nesse ponto, eu já bati e estou apenas esperando os outros terminarem
				// Normalmente, o jogador seguinte vai esperar uma carta de mim, mas como já bati, vou passar uma "carta" vazia
				// Para simular ele olhando para minhas cartas na mesa
				nextPlayer <- Card{}
			} else {
				// Do contrário, vou passar adiante a pior carta da minha mão
				cardToRemove := 0
				worstCard, currentHand = RemoveIndex(cardToRemove, currentHand)
				Print(fmt.Sprintln("Vou passar ", worstCard))
				nextPlayer <- worstCard
			}
		}
		}
	}
}

func insertCard(deck chan<- Card, card Card) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	deck <- card
	Print(fmt.Sprintln("Inseri ", card), false)
}

func getCard(deck <-chan Card) Card {
	return <- deck
}

func main() {
	const PLAYER_AMOUNT = 4

	Print(fmt.Sprintln("======================================"), true)
	Print(fmt.Sprintf("Iniciando dorminhoco com %d jogadores.\n", PLAYER_AMOUNT), true)
	Print(fmt.Sprintln("======================================"), true)

	players := make([]chan Card, PLAYER_AMOUNT)
	for i := 0; i < PLAYER_AMOUNT; i++ {
		players[i] = make(chan Card, CARDS_PER_PLAYER)
	}

	finished := make(chan int, PLAYER_AMOUNT)

	deck := [PLAYER_AMOUNT * CARDS_PER_PLAYER + 1]Card{
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
		playerCards := make([]Card, CARDS_PER_PLAYER)
		for j := 0; j < CARDS_PER_PLAYER; j++ {
			card := getCard(shuffled_deck)
			Print(fmt.Sprintf("Dando %s para o jogador %d\n", card, i))
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
	lastCard := getCard(shuffled_deck)
	close(shuffled_deck)
	Print(fmt.Sprintf("Entregando %s para o 1º jogador\n", lastCard))
	players[0] <- lastCard

	podium := make([]int, PLAYER_AMOUNT)
	Print(fmt.Sprintln("Esperando jogadores terminarem"))
	for i := 0; i < PLAYER_AMOUNT; i++ {
		podium[i] = <- finished
		Print(fmt.Sprintf("Jogador %d terminou\n", podium[i]))
	}
	close(finished)

	Print(fmt.Sprintln("======================================"), true)
	for i := 0; i < PLAYER_AMOUNT - 1; i++ {
		fmt.Printf("Jogador %d ficou em %dº lugar\n", podium[i], i+1)
	}

	fmt.Printf("Jogador %d ficou em último lugar e leva rolhada!\n", podium[PLAYER_AMOUNT - 1])


	Print(fmt.Sprintln("======================================"), true)
	Print(fmt.Sprintln("\tObrigado por jogar!"), true)
}

func RemoveIndex(index int, s []Card) (Card, []Card) {
    return s[index], append(s[:index], s[index+1:]...)
}

// Print prints a message if LOG is true
// 
// If override is true, it prints regardless of LOG
func Print(message string, override ...bool) {
	overrides := (len(override) > 0 && override[0])
	if !overrides {
		return
	}
	if LOG || overrides {
		fmt.Print(message)
	}
}
