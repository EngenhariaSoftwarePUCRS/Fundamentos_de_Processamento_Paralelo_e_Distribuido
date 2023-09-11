/*
	* Felipe Freitas Silva
	* 11/09/2023

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
	// Setting card suits
	Joker = 10
	Spades Suit = iota
	Clubs
	Hearts
	Diamonds
	// Show comments from players during the game
	LOG_PLAYER = true
	// Show comments from the system
	LOG_SYSTEM = false
	// Show prettier (menu-like) messages
	LOG_PRETTY = true
	// Number of players in the game (should be up to 4 due to deck limitations)
	PLAYER_AMOUNT = 4
	// Number of cards each player starts with (should be up to 3 due to deck limitations)
	CARDS_PER_PLAYER = 3
	// Determines the delay between each player's turn (in milliseconds)
	MIN_PLAYER_THINKING_TIME = 0.5 * 1000
	MAX_PLAYER_THINKING_TIME = 1 * 1000
	// Average dealing time per card (in seconds)
	AVG_DEAL_TIME_CARD = 3
	// Determines the delay between each card being dealt (in milliseconds); The greater the value, the more shuffled the deck will be
	DEALING_TIME = (((PLAYER_AMOUNT * CARDS_PER_PLAYER) + 1) / AVG_DEAL_TIME_CARD) * 1000
)

var (
	// Used to print the podium at the end of the game and not disrupt the players "seing" the others finishing
	globalPodium chan int
)

type Suit uint8

func (suit Suit) String() string {
	switch suit {
	case Spades:
		return "Espadas"

	case Clubs:
		return "Paus"

	case Hearts:
		return "Ouros"

	case Diamonds:
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
	if card.symbol == "" {
		return "'Carta vazia'"
	}
	if card.suit == Joker {
		return fmt.Sprint(card.suit)
	}
	return fmt.Sprint(card.symbol, " de ", card.suit)
}

type Deck []Card

func (deck Deck) String() string {
	if len(deck) == 0 {
		return "[Baralho vazio]"
	}
	str := "["
	for _, card := range deck {
		str += fmt.Sprint(card, ", ")
	}
	// Removes the last ", " from the string
	str = str[:len(str) - 2]
	str += "]"
	return str
}

func (deck Deck) RemoveIndex(index int) (Card, Deck) {
	PrintSystem(fmt.Sprintf("Removendo %s[%d]\n", deck, index))
	removedCard := deck[index]
	deck = append(deck[:index], deck[index+1:]...)
	PrintSystem(fmt.Sprintln("Removi", removedCard))
	PrintSystem(fmt.Sprintln("Ficou", deck))
	return removedCard, deck
}

func (deck Deck) Contains(c Card) bool {
	for _, card := range deck {
		if card == c {
			return true
		}
	}
	return false
}

func (deck Deck) Append(c Card) Deck {
	return append(deck, c)
}

func (deck Deck) getWorstCard(jokerManager chan bool) (card Card, index int) {
	PrintSystem(fmt.Sprintln("Analyzing hand:", deck))
	suits := make(map[Suit]int)
	symbols := make(map[string]int)

	for _, card := range deck {
		suits[card.suit]++
		symbols[card.symbol]++
	}

	worstCardIndex := 0
	worstCard := deck[worstCardIndex]
	updateWorstCard := func (index int, card Card) {
		PrintSystem(fmt.Sprintf("Atualizando pior carta de %s para %s\n", worstCard, card))
		worstCardIndex = index
		worstCard = card
	}

	for i := 1; i < len(deck); i++ {
		card := deck[i]

		// If the card is a joker, check if I can forward it
		if (card.suit == Joker || card.symbol == "@") {
			if canForwardJoker(jokerManager) {
				updateWorstCard(i, card)
			} else {
				PrintPlayer(fmt.Sprintln("Infelizmente, ainda não posso passar o coringa adiante"))
			}
			continue
		}

		// Make sure not to mark one of the winnig cards as the worst card and then if there are less cards of the current card's symbol or suit than the worst card's, update the worst card
		if (!(symbols[card.symbol] == CARDS_PER_PLAYER || suits[card.suit] == CARDS_PER_PLAYER) && symbols[card.symbol] < symbols[worstCard.symbol] || suits[card.suit] < suits[worstCard.suit]) {
			updateWorstCard(i, card)
		}
		// // Discarted idea, but could be explored
		// // If there are the same amount of cards of the current card's symbol or suit than the worst card's, randomly update the worst card
		// if suits[card.suit] == suits[worstCard.suit] && symbols[card.symbol] == symbols[worstCard.symbol] {
		// 	if rand.Intn(2) == 0 {
		// 		updateWorstCard(i, card)
		// 	}
		// }
	}

	return worstCard, worstCardIndex
}

func canForwardJoker(jokerManager chan bool) bool {
	if jokerManager == nil {
		return false
	}
	select {
	// If I've passed a turn with it, I can pass it forward
	case <- jokerManager:
		return true

	default:
		// If I can't pass it forward, I'll set a flag to pass it forward next turn
		jokerManager <- true
		return false
	}
}

func insertCard(deck chan<- Card, card Card) {
	time.Sleep(time.Duration(rand.Intn(DEALING_TIME)) * time.Millisecond)
	PrintSystem(fmt.Sprintln("Inserindo", card))
	deck <- card
}

func getCard(deck <-chan Card) Card {
	return <- deck
}

type Player struct {
	// Unique identifier
	id int
	// Player's hand
	hand Deck
	// Channel used to receive cards from the previous player
	receivingChan chan Card
	// Channel used to pass cards to the next player
	sendingChan chan Card
}

func (p Player) String() string {
	return fmt.Sprint("Jogador ", p.id)
}

func (p Player) receiveCard(card Card) Deck {
	return p.hand.Append(card)
}

func (p Player) sendCard(card Card) {
	p.sendingChan <- card
}

func (p Player) canFinish() (bool, Deck) {
	suits := make(map[Suit]int)	
	symbols := make(map[string]int)

	for _, card := range p.hand {
		suits[card.suit]++
		symbols[card.symbol]++
	}
	
	hasAnnounced := false
	winningHandIndex := 0
	winningHand := make(Deck, CARDS_PER_PLAYER)
	for _, card := range p.hand {
		if symbols[card.symbol] == CARDS_PER_PLAYER {
			if !hasAnnounced {
				PrintPlayer(fmt.Sprintf("Posso bater, pois tenho %d %s's\n", CARDS_PER_PLAYER, card.symbol))
				hasAnnounced = true
			}
			winningHand[winningHandIndex] = card
			winningHandIndex++
		} else if suits[card.suit] == CARDS_PER_PLAYER {
			if !hasAnnounced {
				PrintPlayer(fmt.Sprintf("Posso bater, pois tenho %d cartas de %s\n", CARDS_PER_PLAYER, card.suit))
				hasAnnounced = true
			}
			winningHand[winningHandIndex] = card
			winningHandIndex++
		}
	}

	if winningHand[0].symbol != "" {
		return true, winningHand
	}

	return false, nil
}

func (p Player) finish(finished chan<- int) {
	// Wait for a random amount of time to simulate a player waiting for the right moment to finish
	thinkingTime := MIN_PLAYER_THINKING_TIME +
		rand.Intn(MAX_PLAYER_THINKING_TIME - MIN_PLAYER_THINKING_TIME)
	time.Sleep(time.Duration(thinkingTime) * time.Millisecond)
	// Finishes
	finished <- p.id
	PrintPlayer(fmt.Sprintf("(%d) Terminei!\n", p.id))
	// "Alerts" the main goroutine for the podium
	globalPodium <- p.id
}

func (p Player) announcePlay(message string) {
	PrintPlayer(fmt.Sprintf("\n(%d) Vou jogar\n", p.id))
	PrintPlayer(fmt.Sprintln("Minha mão:", p.hand))
	PrintPlayer(message)
}

func (p Player) play(
	nextPlayer chan Card,
	jokerManager chan bool,
	finished chan int,
	startingHand Deck,
) {
	p.hand = startingHand
	p.sendingChan = nextPlayer

    var receivedCard, worstCard Card
	var worstCardIndex int

	for {
		Print(fmt.Sprintf("(%d) De olho na mesa...\n", p.id), false)
		select {
		case fasterPlayer := <- finished: {
			p.announcePlay(fmt.Sprintln(fasterPlayer, "bateu, vou bater também"))
			go p.finish(finished)

			close(jokerManager)
			return
		}

		case receivedCard = <- p.receivingChan: {
			p.announcePlay(fmt.Sprintln("Recebi", receivedCard))
			p.hand = p.receiveCard(receivedCard)

			if possible, winningHand := p.canFinish(); possible {
				PrintPlayer(fmt.Sprintln("Vou bater"))
				go p.finish(finished)

				for _, card := range p.hand {
					if !winningHand.Contains(card) {
						worstCard = card
						break
					}
				}

				// Make sure not to pass a winning card
				if !winningHand.Contains(worstCard) {
					sendCard := func() {
						PrintPlayer(fmt.Sprintln("Vou passar o", worstCard))
						p.sendCard(worstCard)
					}
					// Make sure not to pass the joker if I can't forward it
					if (worstCard.suit == Joker || worstCard.symbol == "@") {
						if canForwardJoker(jokerManager) {
							sendCard()
						}
					} else {
						sendCard()
					}
				}

				close(jokerManager)
				return
			} else {
				// If I can't finish, I'll pass the worst card
				_, worstCardIndex = p.hand.getWorstCard(jokerManager)
				worstCard, p.hand = p.hand.RemoveIndex(worstCardIndex)
				PrintPlayer(fmt.Sprintln("Minha pior carta é o", worstCard))
				PrintPlayer(fmt.Sprintln("Vou passar", worstCard))
				p.sendCard(worstCard)
			}
		}

		default: {
			thinkingTime := MIN_PLAYER_THINKING_TIME +
				rand.Intn(MAX_PLAYER_THINKING_TIME - MIN_PLAYER_THINKING_TIME)

			// Simulates the player thinking about life; not really necessary, but makes the game more interesting as the players aren't all in sync waiting to either take a card or see someone finish immediatly
			time.Sleep(time.Duration(thinkingTime) * time.Millisecond)
		}
		}
	}
}

func main() {
	PrintPretty(fmt.Sprintln("======================================"))
	PrintPretty(fmt.Sprintf("Iniciando dorminhoco com %d jogadores.\n", PLAYER_AMOUNT))
	PrintPretty(fmt.Sprintln("======================================"))

	players := make([]Player, PLAYER_AMOUNT)
	for i := 0; i < PLAYER_AMOUNT; i++ {
		players[i] = Player{
			id: i + 1,
			receivingChan: make(chan Card, CARDS_PER_PLAYER),
		}
	}

	finished := make(chan int, PLAYER_AMOUNT)
	globalPodium = make(chan int, PLAYER_AMOUNT)

	deck := [PLAYER_AMOUNT * CARDS_PER_PLAYER + 1]Card{
		{"12", Spades},
		{"12", Clubs},
		{"12", Hearts},
		{"12", Diamonds},
		{"11", Spades},
		{"11", Clubs},
		{"11", Hearts},
		{"11", Diamonds},
		{"10", Spades},
		{"10", Clubs},
		{"10", Hearts},
		{"10", Diamonds},
		{"@", Joker},
	}

	shuffled_deck := make(chan Card, len(deck))
	for _, card := range deck {
		go insertCard(shuffled_deck, card)
	}
	
	for i, player := range players {
		startingHand := make(Deck, CARDS_PER_PLAYER)
		jokerManager := make(chan bool, 1)
		for j := 0; j < CARDS_PER_PLAYER; j++ {
			card := getCard(shuffled_deck)
			PrintSystem(fmt.Sprintf("Dando %s para o jogador %d\n", card, player.id))
			startingHand[j] = card
		}
		go player.play(
			players[(i+1)%PLAYER_AMOUNT].receivingChan,
			jokerManager,
			finished,
			startingHand,
		)
	}
	lastCard := getCard(shuffled_deck)
	close(shuffled_deck)

	PrintSystem(fmt.Sprintln("Cartas entregues, começando jogo"))
	PrintSystem(fmt.Sprintf("Entregando %s para o 1º jogador\n", lastCard))
	players[0].receivingChan <- lastCard

	PrintSystem(fmt.Sprintln("Esperando jogadores terminarem"))
	podium := make([]int, PLAYER_AMOUNT)
	for i := 0; i < PLAYER_AMOUNT; i++ {
		finishedPlayer := <- globalPodium
		podium[i] = finishedPlayer
		PrintSystem(fmt.Sprintf("Jogador %d terminou\n", finishedPlayer))
	}
	close(finished) 
	close(globalPodium)

	PrintPretty(fmt.Sprintln("\n======================================"))
	for i := 0; i < PLAYER_AMOUNT - 1; i++ {
		Print(fmt.Sprintf("Jogador %d ficou em %dº lugar\n", podium[i], i+1), true)
	}
	Print(fmt.Sprintf("Jogador %d ficou em último lugar e leva rolhada!\n", podium[PLAYER_AMOUNT - 1]), true)


	PrintPretty(fmt.Sprintln("======================================"))
	PrintPretty(fmt.Sprintln("\tObrigado por jogar!"))
}

// Print prints a message if variable LOG_ is true
func Print(message string, log bool) {
	if log {
		fmt.Print(message)
	}
}

func PrintSystem(message string, override... bool) {
	if (len(override) > 0) {
		Print(fmt.Sprint("[Sys]", message), override[0])
	} else {
		Print(fmt.Sprint("[Sys]", message), LOG_SYSTEM)
	}
}

func PrintPlayer(message string, override... bool) {
	if (len(override) > 0) {
		Print(message, override[0])
	} else {
		Print(message, LOG_PLAYER)
	}
}

func PrintPretty(message string, override... bool) {
	if (len(override) > 0) {
		Print(message, override[0])
	} else {
		Print(message, LOG_PRETTY)
	}
}
