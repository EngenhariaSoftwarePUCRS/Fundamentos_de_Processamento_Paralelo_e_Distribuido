/*
	* Felipe Freitas Silva
	* 08/09/2023

	* 1) Quantos clientes podem estar sendo tratados concorrentemente?
	* R: Podem ser tratados tantos clientes quanto a capacidade de processamento do servidor permitir.

	* 2) Agora suponha que o seu servidor pode estar tratando no máximo 10 clientes concorrentemente. Como você faria?
	* R: Criaria uma variável que guarda o valor da quantidade de processos ativos em um momento do tempo, que é incrementada toda vez que um processo (cliente) é chamado e decrementada ao enviar a resposta para o cliente (Olhar Código)
*/

package main

import (
	"fmt"
	"math/rand"
)

const LOG = false

type Request struct {
	v      int
	ch_ret chan int
}

func cliente(i int, req chan Request) {
	var v, r int
	clientChan := make(chan int)
	for {
		v = rand.Intn(1000)
		req <- Request{v, clientChan}
		r = <-clientChan
		if LOG {
			fmt.Println("cli: ", i, " req: ", v, "  resp:", r)
		}
	}
}

func handleRequest(id int, req Request, currentConcurrentProcesses *int) {
	if LOG {
		fmt.Println("                                 handleRequest ", id)
	}
	req.ch_ret <- req.v * 2
	*currentConcurrentProcesses--
}

func servidorConc(in chan Request, maxConcurrentClients int) {
	currentConcurrentProcesses, j := 0, 0
	for currentConcurrentProcesses < maxConcurrentClients {
		currentConcurrentProcesses++
		if LOG {
			fmt.Println("Current concurrent processes: ", currentConcurrentProcesses)
		}
		req := <-in
		j++
		go handleRequest(j, req, &currentConcurrentProcesses)
	}
}

func main() {
	const (
		CLIENT_AMOUNT = 100
		CLIENT_POOL = 10
	)
	fmt.Println("------ Servidores - Criação dinâmica -------")
	serverChan := make(chan Request) // CANAL POR ONDE SERVIDOR RECEBE PEDIDOS
	go servidorConc(serverChan, CLIENT_POOL)      // LANÇA PROCESSO SERVIDOR
	for i := 0; i < CLIENT_AMOUNT; i++ {      // LANÇA DIVERSOS CLIENTES
		go cliente(i, serverChan)
	}
	// <-make(chan struct{})
	for {}
}
