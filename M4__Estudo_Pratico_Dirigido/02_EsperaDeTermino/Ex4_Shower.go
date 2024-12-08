/*
	* Felipe Freitas Silva
	* 01/09/2023
	
	* 1) Todos os valores escritos no canal são lidos?
	* R: Não, visto que assim que o processo principal (main) termina ele 'mata' todas as sub-rotinas, e como há um buffer, a última iteração do laço de escrita sempre termina antes da última leitura

	* 2) Como isto poderia ser resolvido?
	* R: Existem diversas alternativas para resolver o problema, tais quais sincronizar as duas rotinas, por exemplo, e a solução optada foi a utilização de waitGroups, conforme aprendido no exercício anterior
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	const QTD_PROCESSOS = 20
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(QTD_PROCESSOS)

	go shower(ch, &wg)
	for i := 1; i <= QTD_PROCESSOS; i++ {
		ch <- i
	}

	wg.Wait()
}

func shower(c chan int, wg *sync.WaitGroup) {
	for {
		j := <-c
		fmt.Printf("%d\n", j)
		wg.Done()
	}
}
