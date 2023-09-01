/*
	* Felipe Freitas Silva
	* 01/09/2023
	
	* 1) Qual o diagrama de estados que representa a questão 'questaoStSp3'?
	* R: https://dreampuf.github.io/GraphvizOnline/#digraph%20G%20%7B%0A%20%20%20%20%22(x%2C%20y)%22%20-%3E%20%22(0%2C%201)%22%0A%20%20%20%20%22(0%2C%201)%22%20-%3E%20%22(1%2C%201)%22%0A%20%20%20%20%22(1%2C%201)%22%20-%3E%20%22(2%2C%201)%22%2C%22(1%2C%202)%22%0A%20%20%20%20%22(1%2C%202)%22%20-%3E%20%22(2%2C%202)%22%0A%20%20%20%20%22(2%2C%201)%22%20-%3E%20%22(2%2C%202)%22%0A%7D

	* 2) Qual o diagrama de estados que representa a questão 'questaoStSp4'?
	* R: https://dreampuf.github.io/GraphvizOnline/#digraph%20G%20%7B%0A%20%20%20%20%22(x%2C%20y)%22%20-%3E%20%22(1%2C%200)%22%2C%22(0%2C%201)%22%0A%20%20%20%20%22(0%2C%201)%22%20-%3E%20%22(1%2C%201)%22%0A%20%20%20%20%22(1%2C%200)%22%20-%3E%20%22(1%2C%201)%22%0A%20%20%20%20%22(1%2C%201)%22%20-%3E%20%22(2%2C%201)%22%2C%22(1%2C%202)%22%0A%20%20%20%20%22(1%2C%202)%22%20-%3E%20%22(2%2C%202)%22%0A%20%20%20%20%22(2%2C%201)%22%20-%3E%20%22(2%2C%202)%22%0A%7D
*/


package main

var x, y int = 0, 0

func pxs(c <-chan int) {
	x = <- c
	x = 2
}

func pys(c chan<- int) {
	y = 1
	c <- y
	y = 2
}

func questaoStSp3() {
	c := make(chan int, 0)
	go pxs(c)
	go pys(c)
}

func pxs2(c <-chan int) {
	x = 1
	<-c
	x = 2
}

func pys2(c chan<- int) {
	y = 1
	c <- y
	y = 2
}

func questaoStSp4() {
	c := make(chan int, 0)
	go pxs2(c)
	go pys2(c)
}

func main() {
	questaoStSp3()
	questaoStSp4()
	for { }
}
