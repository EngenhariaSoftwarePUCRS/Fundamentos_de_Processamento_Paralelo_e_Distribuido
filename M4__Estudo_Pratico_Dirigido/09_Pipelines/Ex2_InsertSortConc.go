/*
	* Felipe Freitas Silva
	* 06/09/2023

	* 1) Implemente um programa que ordene um vetor de inteiros usando o algoritmo InsertSort.
	* R: Código
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("===== PIPE SORT =====")
	
	const (
		PIPE_SIZE = 15
		UPPER_BOUND = 999
	)

	results := make(chan int, PIPE_SIZE)
	var comparators [PIPE_SIZE + 1]chan int
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
