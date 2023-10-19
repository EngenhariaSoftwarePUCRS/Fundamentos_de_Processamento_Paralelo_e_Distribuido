// (2 pontos) Considerando estado a tupla [x,y], monte os diagramas de estados dos programas abaixo.

package main

func letraA() {
	var x, y int = 0, 0

	func px() {
		x = 1
		x = 2
	}
	
	func py() {
		y = 1
		y = 2
	}

	func main() {
		go px()
		go py()
		<- time.After(1 * time.Second)
	}
	// Diagrama:
	// https://dreampuf.github.io/GraphvizOnline/#digraph%20G%20%7B%0A%20%20%20%20%220%2C%200%22%20-%3E%20%220%2C%201%22%2C%20%221%2C%200%22%0A%20%20%20%20%220%2C%201%22%2C%20%221%2C%200%22%20-%3E%20%221%2C%201%22%0A%20%20%20%20%220%2C%201%22%20-%3E%20%220%2C%202%22%0A%20%20%20%20%221%2C%200%22%20-%3E%20%222%2C%200%22%0A%20%20%20%20%221%2C%201%22%20-%3E%20%221%2C%202%22%2C%20%222%2C%201%22%0A%20%20%20%20%220%2C%202%22%20-%3E%20%221%2C%202%22%0A%20%20%20%20%222%2C%200%22%20-%3E%20%222%2C%201%22%0A%20%20%20%20%221%2C%202%22%2C%20%222%2C%201%22%20-%3E%20%222%2C%202%22%0A%7D%0A
}

func letraB() {
	var x, y int = 0, 0

	func pxs(c chan int) {
		x = 1
		<- c
		c <- x
		x = 2
	}

	func pys(c chan int) {
		y = 1
		c <- y
		<- c
		y = 2
	}

	func main() {
		c := make(chan int)
		go pxs(c)
		go pys(c)
		<- time.After(1 * time.Second)
	}
	
	// Diagrama:
	// https://dreampuf.github.io/GraphvizOnline/#digraph%20G%20%7B%0A%20%20%20%20%220%2C%200%22%20-%3E%20%220%2C%201%22%2C%20%221%2C%200%22%0A%20%20%20%20%220%2C%201%22%2C%20%221%2C%200%22%20-%3E%20%221%2C%201%22%0A%20%20%20%20%221%2C%201%22%20-%3E%20%221%2C%202%22%2C%20%222%2C%201%22%0A%20%20%20%20%221%2C%202%22%2C%20%222%2C%201%22%20-%3E%20%222%2C%202%22%0A%7D%0A
}
