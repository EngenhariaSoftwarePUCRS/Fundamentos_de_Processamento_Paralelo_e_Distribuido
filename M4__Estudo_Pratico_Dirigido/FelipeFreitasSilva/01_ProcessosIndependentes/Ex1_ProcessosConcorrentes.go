/*
	* Felipe Freitas Silva
	* 31/08/2023
	
	* 1) Quantos processos concorrentes são gerados?
	* R: São gerados N (40) processos concorrentes, sem contar o processo principal (main)

	* 2) O que se pode supor sobre a velocidade relativa dos processos?
	* R: Se pode observar após a execução que a saída é completamente aleatória
*/

package main

import "fmt"

const N = 40

func funcaoA(id int, s string) {
	for {
		fmt.Println(s, id)
	}
}

func geraNespacos(n int) string {
	s := "  "
	for j := 0; j < n; j++ {
		s += "   "
	}
	return s
}

func main() {
	for i := 0; i < N; i++ {
		go funcaoA(i, geraNespacos(i))
	}
	for { }
}
