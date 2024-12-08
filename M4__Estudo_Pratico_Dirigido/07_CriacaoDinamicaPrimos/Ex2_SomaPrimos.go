/*
	* Felipe Freitas Silva
	* 05/09/2023

	* 1) Modifique o programa para que o calculo de primos seja sobreposto temporalmente
	* R: Código
*/


package main

import "fmt"

func isPrime(v int, result chan<- int, end chan<- struct{}) {
	if v % 2 == 0 {
		end <- struct{}{}
		return
	}
	for i := 3; i*i <= v; i += 2 {
		if v % i == 0 {
			end <- struct{}{}
			return
		}
	}
	result <- v
	end <- struct{}{}
}

func primesUpTo(n int, primes chan<- int) {
	evaluatedPrimes := make(chan struct{})
	primes <- 2
	for p := 3; p <= n; p += 2 {
		go isPrime(p, primes, evaluatedPrimes)
	}
	for p := 3; p <= n; p += 2 {
		<- evaluatedPrimes
	}
	close(primes)
}

func addPrimesTo(n int) (total int) {
	primes := make(chan int)
	go primesUpTo(n, primes)
	for p := range primes {
		fmt.Printf(" %d ", p)
		total += p
	}
	return total
}

func main() {
	fmt.Println("\nTotal: ",addPrimesTo(100))
}
