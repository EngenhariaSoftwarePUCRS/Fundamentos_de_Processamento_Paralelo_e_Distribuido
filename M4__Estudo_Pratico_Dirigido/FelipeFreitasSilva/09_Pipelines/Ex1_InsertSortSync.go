/*
	* Felipe Freitas Silva
	* 05/09/2023

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
	const (
		PIPE_SIZE = 10
		LIMIT = 999
	)
	var pipe [PIPE_SIZE + 1]int
	fmt.Println("===== InsertSortSync =====")
	rand.Seed(time.Now().UnixNano())
	
	var j int
	for i := 0; i < PIPE_SIZE - 1; i++ {
		valor := rand.Intn(LIMIT) - rand.Intn(LIMIT)
		fmt.Printf("Inserindo %d\n", valor)
		
		// acha posicao em relacao aos demais ja colocados
		for j = 0; j < i; j++ {
			if pipe[j] >= valor {
				break
			}
		}

		// desloca restante para a direita
		for k := i + 1; PIPE_SIZE > k && k >= j; k-- {
			pipe[k+1] = pipe[k]
		}
		pipe[j] = valor
	}
	fmt.Println(pipe)
}
