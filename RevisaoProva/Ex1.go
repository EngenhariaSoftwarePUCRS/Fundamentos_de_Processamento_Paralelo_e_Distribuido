// 1. Explique a razão do código abaixo gerar resultado diferente de 0. Exemplifique e explique uma situaçã oem que o valor de sharedTest torna-se inconsistente.
// R:

package main

import (
	"fmt"
)

var sharedTest int = 0 // variavel comparilhada
var ch_fim chan struct{} = make(chan struct{})
func myFunc(inc int) {
	for k := 0; k < 1000; k++ {
		fmt.Println("Antes: ", sharedTest)
		sharedTest = sharedTest + inc
		fmt.Println("Depois: ", sharedTest)
	}
	ch_fim <- struct{}{}
}

func questao1() {
	for i := 0; i < 10; i++ {
		go myFunc(1)
		go myFunc(-1)
	}
	fmt.Println("Criei 20 processos")
	for i := 0; i < 20; i++ {
		<-ch_fim
	}
	fmt.Println("Processos acabaram. Resultado ", sharedTest)
}

func main() {
	// for i := 0; i < 100; i++ {
		questao1()
	// }
}
