/*
	* Felipe Freitas Silva
	* 05/09/2023


	* Dada uma árvore inicializada e uma operação de caminhamento, pede-se fazer:

	* 1) A operação que soma todos elementos da árvore.
	* R: Linhas [89-95]

	* 2) Uma operação concorrente que soma todos elementos da árvore
	* R: Linhas [97-112]

	* 3) A operação de busca de um elemento v, dizendo true se encontrou v na árvore, ou falso
	* R: Linhas [114-124]

	* 4) A operação de busca concorrente de um elemento, que informa imediatamente por um canal se encontrou o elemento (sem acabar a busca), ou informa que não encontrou ao final da busca
	* R: Linhas [126-140]

	* 5) A operação que escreve todos pares em um canal de saidaPares e todos impares em um canal saidaImpares, e ao final avisa que acabou em um canal fin
	* R: Linhas [221-239]

	* 6) A versão concorrente da operação acima, ou seja, os varios nodos sao testados concorrentemente se pares ou impares, escrevendo o valor no canal adequado
	* R: Linhas [142-148]
*/

package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

type Time struct {
	label string
	start time.Time
	end   time.Time
}

func startTimer(label string) *Time {
	title(label)
	t := &Time{label: label}
	t.start = time.Now()
	return t
}

func (t *Time) stopTimer() {
	t.end = time.Now()
	timeElapsed := t.end.Sub(t.start)
	fmt.Printf(
		"\nTempo de execução: %s (%f segundos)\n",
		timeElapsed,
		timeElapsed.Seconds(),
	)
}

type Nodo struct {
	v int
	e *Nodo
	d *Nodo
}

func caminhaERD(r *Nodo) {
	if r != nil {
		caminhaERD(r.e)
		fmt.Print(r.v, ", ")
		caminhaERD(r.d)
	}
}

func saveCaminhaERD(r *Nodo, values *[]int) {
	if r != nil {
		saveCaminhaERD(r.e, values)
		*values = append(*values, r.v)
		saveCaminhaERD(r.d, values)
	}
}

func saveCaminhaERDConcurrent(r *Nodo, ch chan<- int) {
	if r != nil {
		go saveCaminhaERDConcurrent(r.e, ch)
		ch <- r.v
		go saveCaminhaERDConcurrent(r.d, ch)
	}
}

func soma(r *Nodo) int {
	if r != nil {
		// fmt.Print(r.v, ", ")
		return r.v + soma(r.e) + soma(r.d)
	}
	return 0
}

func somaConc(r *Nodo) int {
	s := make(chan int)
	go somaConcCh(r, s)
	return <-s
}

func somaConcCh(r *Nodo, s chan<- int) {
	if r != nil {
		s1 := make(chan int)
		go somaConcCh(r.e, s1)
		go somaConcCh(r.d, s1)
		s <- (r.v + <-s1 + <-s1)
	} else {
		s <- 0
	}
}

func busca(r *Nodo, n int) bool {
	treeValues := make([]int, 0)
	saveCaminhaERD(r, &treeValues)

	for _, v := range treeValues {
		if n == v {
			return true
		}
	}
	return false
}

func buscaConcorrente(
	r *Nodo,
	n int,
	treeValuesLen int,
	treeValues chan int,
) {
	saveCaminhaERDConcurrent(r, treeValues)
	for i := 0; i < treeValuesLen; i++ {
		if n == <- treeValues {
			printContains(true, n)
			return
		}
	}
	printContains(false, n)
}

func parityConcurrent(v int, saidaPares chan<- int, saidaImpares chan<- int) {
	if isEven(v) {
		saidaPares <- v
	} else {
		saidaImpares <- v
	}
}

func main() {
	const MAX_CONCURRENT_PROCS = 8
	runtime.GOMAXPROCS(MAX_CONCURRENT_PROCS)

	root := &Nodo{v: 10,
		e: &Nodo{v: 5,
			e: &Nodo{v: 3,
				e: &Nodo{v: 1, e: nil, d: nil},
				d: &Nodo{v: 4, e: nil, d: nil}},
			d: &Nodo{v: 7,
				e: &Nodo{v: 6, e: nil, d: nil},
				d: &Nodo{v: 8, e: nil, d: nil}}},
		d: &Nodo{v: 15,
			e: &Nodo{v: 13,
				e: &Nodo{v: 12, e: nil, d: nil},
				d: &Nodo{v: 14, e: nil, d: nil}},
			d: &Nodo{v: 18,
				e: &Nodo{v: 17, e: nil, d: nil},
				d: &Nodo{v: 19, e: nil, d: nil}}}}

	time := startTimer("Valores na árvore")
	caminhaERD(root)
	time.stopTimer()

	time = startTimer("Caminhamento Sync")
	treeValues := make([]int, 0)
	saveCaminhaERD(root, &treeValues)
	for _, v := range treeValues {
		fmt.Print(v, ", ")
	}
	time.stopTimer()
	
	time = startTimer("Caminhamento Concorrente")
	treeValuesC := make(chan int, MAX_CONCURRENT_PROCS)
	saveCaminhaERDConcurrent(root, treeValuesC)
	for i := 0; i < len(treeValues); i++ {
		fmt.Print(<-treeValuesC, ", ")
	}
	time.stopTimer()

	
	time = startTimer("Sum - Sync")
	fmt.Print(soma(root))
	time.stopTimer()

	time = startTimer("Sum - Concurrent")
	fmt.Print(somaConc(root))
	time.stopTimer()


	x_to_search, y_to_search := 10, 2
	
	time = startTimer(fmt.Sprintf("busca(%d) | busca(%d)", x_to_search, y_to_search))
	contains_x := busca(root, x_to_search)
	printContains(contains_x, x_to_search)
	fmt.Print(" | ")
	contains_y := busca(root, y_to_search)
	printContains(contains_y, y_to_search)
	time.stopTimer()

	time = startTimer(fmt.Sprintf("buscaConcorrente(%d) | buscaConcorrente(%d)", x_to_search, y_to_search))
	buscaConcorrente(root, x_to_search, len(treeValues), make(chan int, MAX_CONCURRENT_PROCS))
	fmt.Print(" | ")
	buscaConcorrente(root, y_to_search, len(treeValues), make(chan int, MAX_CONCURRENT_PROCS))
	time.stopTimer()


	saidaPares := make(chan int, len(treeValues))
	saidaImpares := make(chan int, len(treeValues))
	evenAmount, oddAmount := 0, 0

	time = startTimer("Pares e Ímpares - Sync")
	for _, v := range treeValues {
		if isEven(v) {
			saidaPares <- v
			evenAmount++
		} else {
			saidaImpares <- v
			oddAmount++
		}
	}
	fmt.Print("Pares: ")
	for i := 0; i < evenAmount; i++ {
		fmt.Print(<- saidaPares, ", ")
	}
	fmt.Print("\nÍmpares: ")
	for i := 0; i < oddAmount; i++ {
		fmt.Print(<- saidaImpares, ", ")
	}
	time.stopTimer()

	time = startTimer("Pares e Ímpares - Concurrent")
	for _, v := range treeValues {
		go parityConcurrent(v, saidaPares, saidaImpares)
	}
	fmt.Print("Pares: ")
	for i := 0; i < evenAmount; i++ {
		fmt.Print(<- saidaPares, ", ")
	}
	fmt.Print("\nÍmpares: ")
	for i := 0; i < oddAmount; i++ {
		fmt.Print(<- saidaImpares, ", ")
	}
	time.stopTimer()


	title("Obrigado por utilizar!")
}

func title(title string) {
	const MAX_LEN = 50
	char_amount := MAX_LEN - len(title)
	if char_amount < 2 {
		char_amount = 2
	}
	fmt.Printf(
		"\n\n%s %s %s\n",
		strings.Repeat("=", char_amount/2),
		title,
		strings.Repeat("=", char_amount/2),
	)
}

func printContains(contains bool, n int) {
	if contains {
		fmt.Print("Árvore contém ", n)
	} else {
		fmt.Print("Árvore não contém ", n)
	}
}

func isEven(n int) bool {
	return n % 2 == 0
}
