/*
	* Felipe Freitas Silva
	* 01/09/2023
	
	* Suponha que voce tem muitos processos bloqueados tentando escrever em um canal. Será que o processo que está bloqueado há mais tempo é o que ganha o direito de escrever no canal antes dos demais bloqueados? Ou seja, o atendimento dos bloqueados é FIFO? A especificação da linguagem, tanto quanto se pode avaliar, não aborda este aspecto. Isto é um teste. Não serve como prova de que o comportamento é este.
	* 1) Leia o programa, execute algumas vezes e veja o resultado. Qual a sua conclusão ?
	* R: Parece que o programa implementa algum tipo de fila para este caso. Do contrário, deveriamos ver os processos 'retornando' em uma ordem completamente aleatória.
*/


package main

import (
	"fmt"
	"time"
)

// Uma rotina concorrente que tenta escrever em um canal sincronizante
func seBloqueiaNoCanal(i int, c chan<- int) {
	c <- i
}

func main() {
	const QUANTIDADE_ROTINAS = 200
	var chanBloq chan int = make(chan int) // Canal para bloquear as N goRotinas
	for i := 0; i <= QUANTIDADE_ROTINAS; i++ {
		go seBloqueiaNoCanal(i, chanBloq) // Lança N goRotinas
		time.Sleep(20 * time.Microsecond) // Espera um pouco; microsec = (1/1.000.000) sec para que a rotina lançada chegue ao bloqueio. Não é garantido, apenas um teste - dependendo da máquina e OS (Sistema Operacional) deve regular este valor para gerar 0 respostas fora de ordem;
	}
	// Chegando aqui, devemos ter N processos "seBloqueiaNoCanal" bloqueados no canal chanBloq e com alta probabilidade resguardando a ordem de bloqueio conforme seu identificador i.
	// Agora vamos ver se as sincronizações acontecem mantendo a ordem de bloqueio ou não
	outOfOrder := 0 // Contador de elementos fora de ordem
	for i := 0; i <= QUANTIDADE_ROTINAS; i++ {
		// Desbloqueia os N processos, obtendo seu id
		v := <- chanBloq
		// Se o id não for i, está fora de ordem
		if i != v {
			// Neste caso, avisa com um . qual id está fora de ordem
			fmt.Print(" .", v, " ")
			// Computa quantidade fora de ordem
			outOfOrder++
		} else {
			// Mostra o id na ordem
			fmt.Print(v, " ")
		}
	}
	fmt.Println("\n -- \n -- Total de respostas fora de ordem: ", outOfOrder)
	fmt.Println(" -- Caso seja maior que 0, tente aumentar um pouco o tempo de sleep.\n --")
}
