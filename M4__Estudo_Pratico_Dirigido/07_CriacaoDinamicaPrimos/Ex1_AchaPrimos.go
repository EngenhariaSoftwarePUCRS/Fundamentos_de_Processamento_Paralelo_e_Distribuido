/*
	* Felipe Freitas Silva
	* 05/09/2023

	* 1) Torne este programa concorrente, sobrepondo temporalmente o trabalho de computar se um valor é primo.
	* R: Código
*/

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

const AMOUNT_TO_GENERATE = 2000

func main() {
	runtime.GOMAXPROCS(8)

	fmt.Println("===== Conta Primos =====")
	slice := generateSlice(AMOUNT_TO_GENERATE)

	start := time.Now()
	p := contaPrimosSync(slice)
	fmt.Println("Sync - Tempo Decorrido: ", time.Since(start).Seconds())
	fmt.Println("Quantidade de Primos: ", p)

	fmt.Println()

	start = time.Now()
	p = contaPrimosConcurrent(slice)
	fmt.Println("Concurrent - Tempo Decorrido: ", time.Since(start).Seconds())
	fmt.Println("Quantidade de Primos: ", p)
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

func contaPrimosSync(s []int) int {
	result := 0
	for i := 0; i < AMOUNT_TO_GENERATE; i++ {
		if isPrime(s[i]) {
			result++
		}
	}
	return result
}

func contaPrimosConcurrent(ns []int) int {
	result := 0
	ret := make(chan bool, AMOUNT_TO_GENERATE)

	for i := 0; i < AMOUNT_TO_GENERATE; i++ {
		go isPrimeConcurrent(ns[i], ret)
	}
	for i := 0; i < AMOUNT_TO_GENERATE; i++ {
		if <- ret {
			result++
		}
	}
	return result
}

func isPrimeConcurrent(v int, ch chan<- bool) {
	ch <- isPrime(v)
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
