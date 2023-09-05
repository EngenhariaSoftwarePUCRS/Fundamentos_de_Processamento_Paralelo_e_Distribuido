/*
	* Felipe Freitas Silva
	* 04/09/2023

	* 1) Torne este programa concorrente, sobrepondo temporalmente o trabalho de computar se um valor é primo.
	* R: Código
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 2000

func main() {
	fmt.Println("------ DIFFERENT prime count IMPLEMENTATIONS -------")
	slice := generateSlice(N)
	p := contaPrimosSeq(slice)
	fmt.Println("  ------ n primos :  ", p)
}

// Generates a slice filled with random numbers
func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999999999)
	}
	return slice
}

func contaPrimosSeq(s []int) int {
	result := 0
	for i := 0; i < N; i++ {
		if isPrime(s[i]) {
			fmt.Println("  ------ primos :  ", s[i])
			result++
		}
	}
	return result
}

func isPrime(p int) bool {
	if p % 2 == 0 {
		return false
	}
	for i := 3; i*i <= p; i += 2 {
		if p % i == 0 {
			return false
		}
	}
	return true
}
