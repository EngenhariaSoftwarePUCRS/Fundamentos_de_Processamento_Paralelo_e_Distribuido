/*
	* Felipe Freitas Silva
	* 01/09/2023
	
	* 1) Implementar um canal 'sem dados' que apenas passa a mensagem é uma solução para sincronizar o final do programa?
	* R: Sim, porém não é a solução mais ideal, visto que esta dependência é altamente sucetiva a erro humano, e que não há necessariamente ligação entre a espera e o final do programa - isto é - caso a função 'say' fosse mais complexa, por exemplo, é possível que existisse alguma condicial que evitasse a escrita/sinalização de pronto no canal c, e bastaria 1 instância/processo falhar para que o programa ficasse 'pendurado'

	* 2) Aumente para criar 10 processos concorrentes say(...), como fariamos a espera de todos?
	* R: Código atualizado
*/

package main

import "fmt"

func say(s string, c chan struct{}) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
	c <- struct{}{}
}

func main() {
	fin := make(chan struct{})
	/* Código questão 1
	go say("world", fin)
	go say("hello", fin)
	<-fin
	<-fin */
	
	// Código questão 2
	const QTD_PROCESSOS = 20
	sentence := [QTD_PROCESSOS]string{"The", "quick", "brown", "sympathetic", "fox", "jumps", "over", "the", "nice", "lazy", "dog", "as", "it", "says", "the", "following:", "hello", ",", "world", "!"}
	for _, word := range sentence {
		go say(word, fin)
	}

	for i := 0; i < QTD_PROCESSOS; i++ {
		<- fin
	}
}
