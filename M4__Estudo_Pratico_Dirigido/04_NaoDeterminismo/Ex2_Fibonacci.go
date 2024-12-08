/*
	* Felipe Freitas Silva
	* 04/09/2023

	* 1) Ler e entender o código
*/

package main

import "fmt"

func fibonacci(c chan<- int, quit <-chan struct{}) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Tired")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan struct{})

	go func() {
		for i := 0; i < 30; i++ {
			fmt.Println(<-c)
		}
		quit <- struct{}{}
	}()
	fibonacci(c, quit)
}
