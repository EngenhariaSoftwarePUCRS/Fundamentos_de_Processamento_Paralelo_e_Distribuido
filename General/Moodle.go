package main
import "fmt"

var x, y, z int = 0, 0, 0

// -----------------------

func px() {
	x = 1
	x = 2
}

func py() {
	y = 1
	y = 2
}

func pz() {
	z = 1
	z = 2
}

// -----------------------

func main() {
	go px()
	go py()
	go pz()
	// fmt.Println(x, y, z)
	for {}
}
