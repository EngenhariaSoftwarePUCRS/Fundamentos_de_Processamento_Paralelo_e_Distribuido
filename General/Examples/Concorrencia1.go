package main

import (
	"fmt"
	"time"
)

var N int = 40

func funcaoA(id int, s string) {
	for {
		fmt.Println(s, id)
	}
}

func geraNespacos(n int) string {
	s := "  "
	for j := 0; j < n; j++ {
		s = s + "   "
	}
	return s
}

func main() {
	for i := 0; i < N; i++ {
		go funcaoA(i, geraNespacos(i))
	}
	for true {
		time.Sleep(100 * time.Millisecond)
	}
}

// quantos processos concorrentes são gerados ?
// o que se pode supor sobre a velocidade relativa dos mesmos ?
// o sleep no método main serve para este nao acabar, o que acabaria todos processos em execucao
//     mais adiante veremos outras formas de sincronizar isto
