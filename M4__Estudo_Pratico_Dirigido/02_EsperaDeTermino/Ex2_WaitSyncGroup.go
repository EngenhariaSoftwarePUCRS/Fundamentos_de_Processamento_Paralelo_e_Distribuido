/*
	* Felipe Freitas Silva
	* 01/09/2023
	
	* Esta solução é similar à do exercício anterior, porém evita o uso de um laço de repetição a mais para esperar o fim dos processos.
*/

package main

import (
	"fmt"
	"sync"
)

func say(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
	wg.Done()
}

func main() {
	var waitgroup sync.WaitGroup
	/* Código original
	waitgroup.Add(2)
	go say("world", &waitgroup)
	go say("hello", &waitgroup)
	waitgroup.Wait() */
	
	// Código questão 2
	const QTD_PROCESSOS = 20
	waitgroup.Add(QTD_PROCESSOS)
	sentence := [QTD_PROCESSOS]string{"The", "quick", "brown", "sympathetic", "fox", "jumps", "over", "the", "nice", "lazy", "dog", "as", "it", "says", "the", "following:", "hello", ",", "world", "!"}
	for _, word := range sentence {
		go say(word, &waitgroup)
	}

	waitgroup.Wait()
}
