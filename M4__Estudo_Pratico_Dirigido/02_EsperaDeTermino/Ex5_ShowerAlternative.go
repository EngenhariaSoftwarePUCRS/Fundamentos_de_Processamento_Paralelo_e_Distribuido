/*
	* Felipe Freitas Silva
	* 01/09/2023
 
	* 1) Esta é uma solução para a questão anterior?
	* R: Sim, visto que sincroniza a leitura com a escrita, e só permite o termino ao escrever no canal quit, após escrever e ler todos os valores

	* 2) O que garante que todos os valores serão lidos antes do final do programa?
	* R: A sincronização citada acima
*/

package main

import "fmt"

func main() {
	ch := make(chan int)
	quit := make(chan struct{})
	go shower(ch, quit)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	quit <- struct{}{}
}

func shower(c <-chan int, quit <-chan struct{}) {
	for {
		select {
		case j := <-c:
			fmt.Printf("%d\n", j)
		case <-quit:
			break
		}
	}
}
