/*
	* Felipe Freitas Silva
	* 04/09/2023

	* 1) Crie outros canais para eventos temporizados e declare reações aos mesmos junto ao mesmo select
	* R: Código
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	tic := time.Tick(100 * time.Millisecond)
	tac := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	silence := time.After(400 * time.Millisecond)
	isSilenced := false
	for {
		select {
		case <-tic:
			if !isSilenced {
				fmt.Println("tic")
			}
		case <-tac:
			if !isSilenced {
				fmt.Println("tac")
			}
		case <-boom:
			fmt.Println("\tBOOM!")
			return
		case <-silence:
			isSilenced = true
			fmt.Printf("\n\n\n\n\n")
		default:
			if !isSilenced {
				fmt.Println("\t.")
				time.Sleep(50 * time.Millisecond)
			}
		}
	}
}
