/*
	* Felipe Freitas Silva
	* 01/09/2023
	
	* 1) O que você conclui da execução do programa?
	* R: A execução é completamente aleatória, e pode gerar qualquer combinação de "hello" e "worlds", sendo que ao imprimir o último "hello" o programa termina, independente de quantos "world" foram impressos
*/

package main

import "fmt"

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
