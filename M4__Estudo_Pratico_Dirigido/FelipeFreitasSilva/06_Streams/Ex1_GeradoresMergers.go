/*
	* Felipe Freitas Silva
	* 04/09/2023

	* 1) Com estes processos, instancie outras topologias. Por exemplo, faça com que o resultado de um merger va para outro merger que recebe também de um terceiro gerador.
	* R: 
*/

package main

func geraValores(amount int, c chan<- int) {
	for i := 0; i < amount; i++ {
		c <- int(i)
	}
}

func merger(rec1, rec2 <-chan int, c chan<- int) {
	for {
		select {
		case dado := <-rec1:
			println("Gerador 1, dado:", dado)
			c <- dado
		case dado := <-rec2:
			println("Gerador 2, dado:", dado)
			c <- dado
		}
	}
}

func consumer(nome string, rec <-chan int) {
	for {
		dado := <-rec
		println("Em ", nome, " chegou ", dado)
	}
}

func main() {
	const (
		// Obs: Programa sempre para em deadlock ao chegar neste número
		AMOUNT_TO_GENERATE = 200
		BUFFER_SIZE = 10
	)
	geradorUm := make(chan int, BUFFER_SIZE)
	geradorDois := make(chan int, BUFFER_SIZE)
	geradorTres := make(chan int, BUFFER_SIZE)
	merger1Send := make(chan int, BUFFER_SIZE)
	merger2Send := make(chan int, BUFFER_SIZE)

	go consumer("geral", merger2Send)

	go merger(geradorUm, geradorDois, merger1Send)
	go merger(geradorUm, geradorDois, merger2Send)

	go geraValores(AMOUNT_TO_GENERATE, geradorUm)
	go geraValores(AMOUNT_TO_GENERATE, geradorDois)
	go geraValores(AMOUNT_TO_GENERATE, geradorTres)

	<-make(chan struct{})
}
