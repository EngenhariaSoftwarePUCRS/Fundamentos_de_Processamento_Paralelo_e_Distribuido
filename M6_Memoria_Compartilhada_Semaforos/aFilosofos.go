package main

import (
	"fmt"
	"strconv"
	"time"
)

const (
	PHILOSOPHERS = 5
	FORKS        = 5
)

type Fork struct {}

func philosopher(id int, first_fork, second_fork chan Fork) {
	for {
		fmt.Println(strconv.Itoa(id) + " senta")
		<-first_fork
		fmt.Println(strconv.Itoa(id) + " pegou direita")
		<-second_fork
		fmt.Println(strconv.Itoa(id) + " pegou esquerda")
		fmt.Println(strconv.Itoa(id) + " come")
		first_fork <- Fork{}
		second_fork <- Fork{}
		fmt.Println(strconv.Itoa(id) + " levanta e pensa")
	}
}

func main() {
	var fork_channels [FORKS]chan Fork
	for i := 0; i < FORKS; i++ {
		fork_channels[i] = make(chan Fork, 1)
		fork_channels[i] <- Fork{} // no inicio garfo esta livre
	}
	for i := 0; i < (PHILOSOPHERS); i++ {
		fmt.Println("Filosofo " + strconv.Itoa(i))
		if i % 2 == 0 {
			go philosopher(i, fork_channels[i], fork_channels[(i+1)%PHILOSOPHERS])
		} else {
			go philosopher(i, fork_channels[(i+1)%PHILOSOPHERS], fork_channels[i])
		}
	}
	<-time.After(500 * time.Millisecond)
}
