/*
	* Felipe Freitas Silva
	* 08/09/2023

	* 1) Implemente um programa que ordene um vetor de inteiros usando o algoritmo InsertSort e cujo número de processos concorrentes seja dinâmico.
	* R: Código
*/

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func getArgs() (pipeSize int, upperBound int) {
	pipeSize = 15
	upperBound = 999

	if len(os.Args) > 1 {
		amount, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Erro ao ler argumentos de terminal")
			return;
		}
		if amount > 0 {
			pipeSize = amount
		}
		if amount > upperBound {
			upperBound *= 10
		} else if amount < upperBound / 10 {
			upperBound /= 10
		}
	}
	fmt.Printf(
		"Executando processo com %d canais e gerando números entre [%d, %d]\n",
		pipeSize, -upperBound, upperBound,
	)
	if len(os.Args) <= 1 {
		fmt.Println("Caso queira alterar a quantidade de processos concorrentes, execute o programa passando um argumento numérico.\nExemplo: go run Ex3_InsertSortConc.go 15")
	}

	return pipeSize, upperBound
}

func main() {
	fmt.Println("===== PIPE SORT =====")

	var PIPE_SIZE, UPPER_BOUND = getArgs()

	results := make(chan int, PIPE_SIZE)
	comparators := make([]chan int, PIPE_SIZE + 1)
	for i := 0; i <= PIPE_SIZE; i++ {
		comparators[i] = make(chan int, 2)
	}
	for i := 0; i < PIPE_SIZE; i++ {
		go cellSorter(i, comparators[i], comparators[i + 1], results, UPPER_BOUND)
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < PIPE_SIZE; i++ {
		valor := rand.Intn(UPPER_BOUND) - rand.Intn(UPPER_BOUND)
		comparators[0] <- valor
		fmt.Printf("(%d)Entra %d \n", i, valor)
	}
	comparators[0] <- UPPER_BOUND + 1
	for i := 0; i < PIPE_SIZE; i ++ {
		v := <- results
		fmt.Printf("\tResult[%d]: %d \n", i, v)
	}
	<- comparators[PIPE_SIZE]
}

func cellSorter(i int, in chan int, out chan int, result chan int, max int) {
	var myVal int
	undef := true

	for {
		read := <- in
		if read == max + 1 {
			result <- myVal
			out <- read
			break
		}
		if undef {
			myVal = read
			undef = false
		} else if read >= myVal {
			out <- read
		} else {
			out <- myVal
			myVal = read
		}
	}
}
