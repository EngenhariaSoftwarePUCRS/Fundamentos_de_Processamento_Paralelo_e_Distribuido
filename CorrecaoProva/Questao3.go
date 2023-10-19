// (3 pontos) Suponha que seu colega dispõe da implementação de pilha como abaixo. Os "Nodo"s são simplesmente encadeados, cada um tem um inteiro e aponta para o próximo Nodo. "empilha" coloca um nodo antes do apontado por "pilha". "desempilha" retira o primeiro apontado, como abaixo. A pilha funciona corretamente se utilizada de forma sequencial. Caso considere haver algum problema no uso sequencial, sinalize ou pergunte ao professor.
// Entretanto, seu colega ainda não fez esta disciplina e ele ingenuamente quer usar estas operações com múltiplas threads. Ele está obtendo um comportamento estranho da pilha e não sabe a razão.

// Código disponibilizado pelo professor:
// ***** Pilha - tipo de dado e operações

// 01: tipo Nodo {
// 02:   valor: int,
// 03:   next: referecia para Nodo
// 04: }

// 05: funcao novoNodo(valor) retorna: referencia para Nodo {
// 06: 		retorna: referencia para novo Nodo(valor, nulo)
// 07: }

// 08: funcao empilha(valor: int, pilha: referencia para Nodo) retorna: referencia para Nodo {
// 09:		se pilha == nulo
// 10:		entao retorna novoNodo(valor)
// 11:			senao aux = novoNodo(valor)
// 12:				aux.next = pilha
// 13:				retorna aux
// 14:
// 15: }

// 16: funcao desempilha(pilha: referencia para nodo)
// 17: 	retorna: referencia para Nodo, boolean sucesso, valor retirado {
// 18:		se pilha == nulo
// 19:		entao retorna nulo, falso, 0
// 20:		senao retorna pilha.next, verdadeiro, pilha.valor
// 21: }

// *************************************************************
// ***** usando a pilha

// 22: Var pilha referencia para Nodo // variavel gloal
// 23: Thread produtora() {
// 24:		Loop {
// 25:			pilha = empilha(valorRandomico(), pilha)
// 26:		}
// 27: }

// 28: Thread consumidora() {
// 29:		Loop {
// 30:			pilha, ok, v = desempilha(pilha)
// 31:			se ok entao print(" valor ",v)
// 32:		}
// 33: }

// 34: Main() {
// 35:		Inicia thread p1=produtora()
// 36:		Inicia thread p2=produtora()
// 37:		Inicia thread c1=consumidora()
// 38:		Inicia thread c2=consumidora()
// 39:		Espera termino de p1,p2,c1,c2
// 40: }

// Código corrigido em Go:

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var pilha *Nodo

type Nodo struct {
  valor int
  next *Nodo
}

func novoNodo(valor int) *Nodo {
		return &Nodo{valor, nil}
}

func empilha(valor int, pilha *Nodo) *Nodo {
	if pilha == nil {
		return novoNodo(valor)
	} else {
		aux := novoNodo(valor)
		aux.next = pilha
		return aux
	}
}

func desempilha(pilha *Nodo) (*Nodo, bool, int) {
	if pilha == nil {
		return nil, false, 0
	} else {
		return pilha.next, true, pilha.valor
	}
}

func produtora() {
	for {
		randomValue := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(99999)
		fmt.Println("Empilhando", randomValue)
		pilha = empilha(randomValue, pilha)
	}
}

func consumidora() {
	for {
		var ok bool
		var v int
		pilha, ok, v = desempilha(pilha)
		if !ok {
			fmt.Println("Não consegui desempilhar : (")
		} else {
			fmt.Println("Desempilhando ", v)
		}
	}
}

func main() {
	go produtora()
	go produtora()
	go consumidora()
	go consumidora()
	<- time.After(500 * time.Millisecond)
}

// Questões:

// a) Aponte ao menos duas situações em que o uso concorrente da pilha com as operações empilha e desempilha vai gerar resultado diferente no conjunto de valores da pilha, em relação a aplicação sequencial das mesmas operações. Considere as threads produtoras e consumidoras p1, p2, c1, c2 e diga em que linha cada thread envolvida deve estar para que o problema ocorra, e qual o problema que ocorre. Use o formato: "threads envolvidas p1 e p2: linha de p1, linha de p2:"

// a) threads envolvidas p1 e c2: linha de p1 (13), linha de c2 (30)
// Assumindo início = [5, 8, 4]
// 1. ao tentar empilhar o valor 6, salva referência para o 5
// 2. Enquanto isso, removi 5 da pilha
// 3. Pilha é atualizada na linha 30 como [8, 4]
// 4. Ao atualizar novamente a variável pilha, temos [6, 5, 8, 4] o que "invalidou" o desempilha

// b) threads envolvidas: c1 e c2: linhas 30/31
// Aqui, é possível n(2) threads removerem a mesma referência, tornando n(2) comandos "desempilha" em um só

// Resolva o problema do acesso concorrente com canais, e depois com semáforos. Questões abaixo.

// b) Declare e use canais para resolver o problema de sincronização. Diga quais operações de canais voce colocaria em quais linhas. Por exemplo voce pode dizer:
// 	depois de 22, declare canal c do tipo ... com capacidade ...
//	entre linhas ... e ...: c <- ... (escrita de valor do tipo)
// 	entre linhas ... e ...: ... = <- c (leitura do canal)

// depois de 22, insira const fakeSem = make(chan struct{}, 1)
// entre linhas 24 e 25, insira fakeSem <- struct{}{}
// entre linhas 25 e 26, insira <- fakeSem
// entre linhas 29 e 30, insira fakeSem <- struct{}{}
// entre linhas 30 e 31, insira <- fakeSem

// c) Declare e use semáforos para resolver o problema de sincronização. Diga quais operações de semáforos voce colocaria em quais linhas. Por exemplo voce pode dizer:
// 	depois de 22, declare semaforo s inicializado com ...
// 	entre linhas ... e ... use s.wait()

// antes de 22, insira const sem = NewSemaphore(1)
// entre linhas 24 e 25, insira sem.Wait()
// entre linhas 25 e 26, insira sem.Signal()
// entre linhas 29 e 30, insira sem.Wait()
// entre linhas 30 e 31, insira sem.Signal()
