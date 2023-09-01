/*
	* Felipe Freitas Silva
	* 01/09/2023
	
	* 1) Esta é uma solução para a questão anterior?
	* R: 

	* 2) O que garante que todos os valores serão lidos antes do final do programa?
	* R: 
*/

package main

import "fmt"

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go shower(ch, quit)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	quit <- false // or true, does not matter
}

func shower(c chan int, quit chan bool) {
	for {
		select {
		case j := <-c:
			fmt.Printf("%d\n", j)
		case <-quit:
			break
		}
	}
}
