/*
	* Felipe Freitas Silva
	* 12/09/2023

	* 1) Monte uma topologia de processos e canais formando um anel onde os processos passam circularmente um sinal (token). Cada vez que o processo recebe o sinal, ele faz print do seu identificador em uma coluna diferente da tela.
	* R: Código
*/

package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	AMOUNT_NODES = 40
	LOG = true
	DEBUG = true
)

// Node is a representation of a node in a ring
type Node struct {
	id   		int
	mainChan	chan struct{}
	prev 		*Node
	next 		*Node
}

// Returns a String representation of the node; example: (1)
func (n Node) String() string {
	return fmt.Sprintf("(%d)", n.id)
}

// Returns a new node with the given id and a channel
func NewNode(id int) Node {
	return Node{
		id: id,
		mainChan: make(chan struct{}),
		prev: nil,
		next: nil,
	}
}

// Waits for a signal from the previous node
func (n Node) AwaitSignal() {
	<-n.mainChan
}

// Sends a signal to the next node
func (n Node) SendSignal() {
	n.next.mainChan <- struct{}{}
}

// Runs the node; it will wait for a signal from the previous node, execute an action, and then send a signal to the next node
func (n Node) Run() {
	for {
		n.AwaitSignal()
		Print(n.id)
		n.SendSignal()
	}
}

// Ring is a representation of a ring of nodes
type Ring struct {
	nodes [AMOUNT_NODES]Node
	end chan struct{}
}

// Returns a representation of the ring
// 
// Example:
// 
// -> (1) -> (2) -> (3) -> (4) -> (5) -> (1) ->
func (ring Ring) String() string {
	str := "-> "
	for i := 0; i < len(ring.nodes); i++ {
		str += fmt.Sprint(ring.nodes[i], " -> ")
	}
	str += fmt.Sprint(ring.nodes[0], " ->")
	return str
}

func (ring Ring) SetupNodes() *Ring {
	for i := 0; i < AMOUNT_NODES; i++ {
		node := NewNode(i + 1)
		ring.nodes[i] = node
	}

	for i := 0; i < AMOUNT_NODES; i++ {
		ring.nodes[i].prev = &ring.nodes[(i - 1 + AMOUNT_NODES) % AMOUNT_NODES]
		ring.nodes[i].next = &ring.nodes[(i + 1) % AMOUNT_NODES]
	}

	return &ring
}

func (ring Ring) startNodes() *Ring {
	for i := 0; i < len(ring.nodes); i++ {
		go ring.nodes[i].Run()
	}
	return &ring
}

func (ring Ring) Start() *Ring {
	ring.startNodes().SignalNode(0)
	return &ring
}

func (ring Ring) Stop() {
	if DEBUG {
		fmt.Println("\033[31mEncerrando...\033[0m")
	}
	ring.end <- struct{}{}
}

func (ring Ring) SignalNode(i int) *Ring {
	ring.nodes[i].prev.SendSignal()
	if DEBUG {
		fmt.Println("\033[32mSinal enviado paro nó", ring.nodes[i], "\033[0m")
	}
	return &ring
}

// Waits for a given amount of time (in milliseconds) before executing a function
func (ring Ring) Delay(delay int, delayMsg string) *Ring {
	if DEBUG {
		fmt.Println(delayMsg)
	}
	<- time.After(time.Duration(delay) * time.Millisecond)
	return &ring
}

func main() {
	end := make(chan struct{}, 1)
	ring := Ring{
		nodes: [AMOUNT_NODES]Node{},
		end: end,
	}
	ring.
		SetupNodes().
		Start().
		Delay(5000,
			"\033[36mEsperando 5 segundos para enviar mais um sinal...\033[0m").
		SignalNode(AMOUNT_NODES / 2 - 1).
		Delay(5000,
			"\033[36mEsperando 5 segundos para encerrar...\033[0m").
		Stop()

	<- end
}

// Prints a String with i spaces in the beginning
func Print(i int) {
	if LOG {
		fmt.Println(strings.Repeat("  ", i) + fmt.Sprint(i))
	}
}
