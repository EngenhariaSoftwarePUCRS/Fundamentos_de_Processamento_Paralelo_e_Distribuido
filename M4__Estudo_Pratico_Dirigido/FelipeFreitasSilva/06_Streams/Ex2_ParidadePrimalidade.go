/*
	* Felipe Freitas Silva
	* 04/09/2023

	* 1) Completar o programa utilizando processos concorrentes que se comunicam por canais.
	* R: Código
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func geraValores(amount int, c chan<- int) {
	for i := 0; i < amount; i++ {
		numero := rand.Int31n(2000) + 1
		c <- int(numero)
	}
}

func separadorImparesPares(numbers <-chan int, evens, odds chan<- int) {
	for {
		n := <- numbers
		if isEven(n) {
			evens <- n
		} else {
			odds <- n
		}
	}
}

// Separa ímpares de acordo com primalidade
func separadorPrimos(odds chan int, primes chan<- int) {
	for {
		n := <- odds
		if isPrime(n) {
			primes <- n
		} else {
			odds <- n
		}
	}
}

func consumer(nome string, rec <-chan int) {
	for {
		dado := <-rec
		fmt.Printf("%s(%d)\n", nome, dado)
	}
}

func isEven(n int) bool {
	return n % 2 == 0
}

func isPrime(n int) bool {
	if isEven(n) {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	const AMOUNT_TO_GENERATE = 200
	const BUFFER_SIZE = 1

	rand.Seed(time.Now().UnixNano())

	numbers := make(chan int, BUFFER_SIZE)
	evens := make(chan int, BUFFER_SIZE)
	odds := make(chan int, BUFFER_SIZE)
	primes := make(chan int, BUFFER_SIZE)

	// Consome pares
	go consumer("Even", evens)
	// Consome ímpares não primos
	go consumer("\t\tOdd", odds)
	// Consome primos
	go consumer("\t\t\t\tPrime", primes)

	// Separam ímpares e primos
	go separadorPrimos(odds, primes)
	go separadorPrimos(odds, primes)
	
	// Separam pares e ímpares
	go separadorImparesPares(numbers, evens, odds)
	go separadorImparesPares(numbers, evens, odds)

	go geraValores(AMOUNT_TO_GENERATE, numbers)
	go geraValores(AMOUNT_TO_GENERATE, numbers)

	<-make(chan struct{})
}
