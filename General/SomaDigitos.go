package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	BUFFER_SIZE = 1
)

var (
	result = 0
)

func main() {
	num := "123"
	if len(os.Args) > 1 {
		num = os.Args[1]
	}

	for index, num := range(num) {
		val, err := strconv.Atoi(string(num))
		if err != nil {
			fmt.Println("Error during conversion")
			return
		}
		go add(val, index)
	}

	fmt.Printf("Result: %d", result)
}

func add(val, id int) {
	fmt.Printf("%d) %d\n", id, val)
	result += val
}
