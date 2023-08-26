package main;

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a := "123"
	if len(os.Args) > 1 {
		a = os.Args[1]
	}

	var aArr [3]int
	for i, j := len(a) - 1, 0; i >= 0; i, j = i-1, j+1 {
		val, err := strconv.Atoi(string(a[i]))
		if err != nil {
			fmt.Println("Error during conversion")
			return
		}
		aArr[j] = val
	}

	result := 0
	for i := 0; i < len(a); i++ {
		result += int(aArr[i]) * IntPow(10, i)
	}

	fmt.Printf("%d", result)
}

func IntPow(n, m int) int {
    if m == 0 {
        return 1
    }
    result := n
    for i := 2; i <= m; i++ {
        result *= n
    }
    return result
}
