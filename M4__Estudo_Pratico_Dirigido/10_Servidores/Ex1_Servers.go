/*
	* Felipe Freitas Silva
	* 08/09/2023

	* 1) Adapte o código para que o servidor trate os clientes concorrentemente.
	* R: Código
*/

package main

import (
	"fmt"
	"math/rand"
)

type Request struct {
	v      int
	ch_ret chan int
}


func cliente(i int, req chan Request) {
	var v, r int
	my_ch := make(chan int)
	for {
		v = rand.Intn(1000)
		req <- Request{v, my_ch}
		r = <-my_ch
		fmt.Println("cli: ", i, " req: ", v, "  resp:", r)
	}
}

func handleResponse(in <-chan Request) {
	req := <-in
	fmt.Println("                       trataReq ", req)
	go sendResponse(req.ch_ret, req.v * 2)
}

func sendResponse(receiver chan<- int, response int) {
	receiver <- response
}

func servidorSeq(in <-chan Request) {
	for {
		go handleResponse(in)
	}
}


func main() {
	const CLIENT_AMOUNT = 10
	serv_chan := make(chan Request)
	for i := 0; i < CLIENT_AMOUNT; i++ {
		go cliente(i, serv_chan)
	}
	servidorSeq(serv_chan)
}
