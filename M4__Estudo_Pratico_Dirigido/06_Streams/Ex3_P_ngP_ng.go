/*
	* Felipe Freitas Silva
	* 04/09/2023

	* 1) Escreva o programa que fica fazendo ifinitamente pang-peng-ping-pong-pung-pang-peng...
	* R: Código
 
	* 2) Quantas bolas podem ser colocadas neste sistema?
	* R: O limite é sempre 1 bolinha a menos do que o número de quadras
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Ball struct {
	Colour string
}

func p_ng(sound string, receivingCourt <-chan Ball, sendingCourt chan<- Ball) {
	for {
		ball := <-receivingCourt
		fmt.Print(ball.Colour,sound,"-",getColour("Reset"))
		sendingCourt <- ball
	}
}

func getColour(colour string) string {
	switch colour {
		case "Reset":
			return "\033[0m"
		case "Red":
			return "\033[31m"
		case "Green":
			return "\033[32m"
		case "Yellow":
			return "\033[33m"
		case "Blue":
			return "\033[34m"
		case "Purple":
			return "\033[35m"
		case "Cyan":
			return "\033[36m"
		case "Gray":
			return "\033[37m"
		case "White":
			return "\033[97m"
	}
	return "\033[0m"
}

func getBallColour(ball int) string {
	switch ball % 5 {
	case 0:
		return getColour("Red")
	case 1:
		return getColour("Green")
	case 2:
		return getColour("Yellow")
	case 3:
		return getColour("Blue")
	case 4:
		return getColour("Purple")
	}
	return getColour("Reset")
}

func getSound(court int) string {
	switch court % 5 {
	case 0:
		return "pang"
	case 1:
		return "peng"
	case 2:
		return "ping"
	case 3:
		return "pong"
	case 4:
		return "pung"
	}
	return "error"
}

func main() {
	const COURTS_AMOUNT = 5
	BALLS_AMOUNT := 4
	if len(os.Args) > 1 {
		amount, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Erro ao ler argumentos de terminal")
			return;
		}
		BALLS_AMOUNT = amount
	}
	if BALLS_AMOUNT >= COURTS_AMOUNT {
		BALLS_AMOUNT = COURTS_AMOUNT - 1
	}
	var courts [COURTS_AMOUNT]chan Ball

	for i := 0; i < COURTS_AMOUNT; i++ {
		courts[i] = make(chan Ball)
	}
	
	for i := 0; i < COURTS_AMOUNT; i++ {
		go p_ng(getSound(i), courts[i], courts[(i+1)%COURTS_AMOUNT])
	}

	fmt.Print(getColour("Reset"))
	fmt.Println("\n MENU\t")
	fmt.Println("=========")
	for i := 0; i < BALLS_AMOUNT; i++ {
		colour := getBallColour(i)
		fmt.Print(colour)
		fmt.Print("Bolinha ", i + 1)
		fmt.Println(getColour("Reset"))
	}
	fmt.Println("=========")

	for i := 0; i < BALLS_AMOUNT; i++ {
		colour := getBallColour(i)
		courts[i%COURTS_AMOUNT] <- Ball{colour}
		time.Sleep(500 * time.Millisecond)
	}

	<- make(chan struct{})
}
