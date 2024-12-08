/*
	* Felipe Freitas Silva
	* 08/09/2023 - 12/09/2023

	* 1) Quantos clientes podem estar sendo tratados concorrentemente?
	* R: Podem ser tratados tantos clientes quanto a capacidade de processamento do servidor permitir.

	* 2) Agora suponha que o seu servidor pode estar tratando no máximo 10 clientes concorrentemente. Como você faria?
	* R: Criaria um canal com capacidade de 10 clientes, e a cada requisição, o servidor adicionaria um cliente ao canal, e quando o cliente terminasse de ser atendido, o servidor removeria o cliente do canal, o que permite que no máximo 10 clientes sejam atendidos ao mesmo tempo, enquanto os outros ficam "na fila", esperando para escrever no canal.
*/

package main

import (
	"fmt"
	"math/rand"
)

const (
	// Upper limit of clients
	CLIENT_AMOUNT = 100
	// Greatest amount of concurrent client requests
	CLIENT_POOL = 10
	// Show display messages regarding the processess
	LOG = false
)

type Request struct {
	value   int
	retChan	chan int
}

type Client struct {
	id int
	mainChan chan int
}

// Runs the client, sending requests and receiving responses
// 
// If a limit is passed, the client will stop after sending that amount of requests
func (c Client) Run(req chan Request, limit... int) {
	while := func(i int) bool {
		if len(limit) > 0 {
			return i < limit[0]
		}
		return true
	}
	for i, value, res := 0, 0, 0; while(i); i++ {
		value = rand.Intn(1000)
		req <- Request{value, c.mainChan}
		res = <-c.mainChan
		if LOG {
			fmt.Println("cli: ", c.id, " req: ", value, "  resp:", res)
		}
	}
}

type Server struct {
	mainChan chan Request
	currentConcurrentRequests *int
	maxConcurrentRequests int
	concurrentRequestsHandler chan struct{}
}

func NewServer(mainChan chan Request, maxConcurrentRequests int) Server {
	currentConcurrentRequests := 0
	return Server{
		mainChan:                  mainChan,
		currentConcurrentRequests: &currentConcurrentRequests,
		maxConcurrentRequests:     maxConcurrentRequests,
		concurrentRequestsHandler: make(chan struct{}, maxConcurrentRequests),
	}
}

// Runs the server, receiving requests and handling them
// 
// If a limit is passed, the server will stop after handling that amount of requests
func (s Server) Run(limit... int) {
	while := func(i int) bool {
		if len(limit) > 0 {
			return i < limit[0]
		}
		return true
	}
	for i := 0; while(i); i++ {
		if LOG {
			fmt.Println("Current Concurrent Requests: ", *s.currentConcurrentRequests)
		}
		req := <-s.mainChan
		s.AddConcurrentRequest()
		go s.handleRequest(i, req)
	}
}

// Builds up to MAX(10) concurrent requests and then waits until one is free
func (s Server) AddConcurrentRequest() {
	(*s.currentConcurrentRequests)++
	s.concurrentRequestsHandler <- struct{}{}
}

// Frees one of the concurrent requests
func (s Server) RemoveConcurrentRequest() {
	<-s.concurrentRequestsHandler
	(*s.currentConcurrentRequests)--
}

// Receives a request and returns it's value doubled
func (s Server) handleRequest(id int, req Request) {
	if LOG {
		fmt.Println("                                 handleRequest ", id)
	}
	req.retChan <- req.value * 2
	s.RemoveConcurrentRequest()
}

func main() {
	fmt.Println("------ Servidores - Criação dinâmica -------")
	server := NewServer(make(chan Request), CLIENT_POOL)
	go server.Run()
	for i := 1; i <= CLIENT_AMOUNT; i++ {
		client := Client{i, make(chan int)}
		go client.Run(server.mainChan, 10)
	}

	// Runs indefinetly
	<-make(chan struct{})
}
