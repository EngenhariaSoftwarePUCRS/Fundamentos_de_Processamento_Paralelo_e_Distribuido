/*
	* Felipe Freitas Silva
	* 04/09/2023

	* 1) Encontre e corrija o problema de deadlock que ocorre
	* R: Código

	* 2) Faça uma nova solução que garanta que no máximo N(duas) confirmações estão aguardando leitura pela fonte
	* R: Código
*/

package main

import "fmt"

func Gerador(amount int, solicitaEnvio chan<- int) {
	for i := 1; i < amount; i++ {
		solicitaEnvio <- i
	}
}

func Fonte(solicitaEnvio <-chan int, envia chan<- int, confirma <-chan struct{}) {
	contConf := 0
	for {
		select {
		case x := <-solicitaEnvio:
			envia <- x
		case <-confirma:
			contConf++
		}
	}
}

func FonteEx1(solicitaEnvio <-chan int, envia chan<- int, confirma <-chan struct{}) {
	contConf := 0
	for {
		// Ex 1: Escreve o que ler do solicitaEnvio "em uma operação, 'simular processo síncrono'; independente de buffer"
		envia <- <-solicitaEnvio
		<-confirma
		contConf++
	}
}

func FonteEx2(
	solicitaEnvio <-chan int,
	envia chan<- int,
	confirma <-chan struct{},
	readingLimit int,
) {
	contConf := 0
	amountReading := 0
	for {
		// Ex 2: Validar quantidade de leituras simultaneas
		if (amountReading < readingLimit) {
			select {
			case x := <-solicitaEnvio:
				amountReading++
				envia <- x
			case <-confirma:
				amountReading--
				contConf++
			}
		} else {
			<-confirma
			amountReading--
			contConf++
		}
	}
}

func Destino(envia <-chan int, confirma chan<- struct{}) {
	for {
		rec := <-envia         // recebe valor
		confirma <- struct{}{} // confirma
		fmt.Print(rec, ", ")
	}
}

func main() {
	const (
		// Observação: Programa sempre cai em 'deadlock' neste número
		AMOUNT_TO_GENERATE = 100
		BUFFER_SIZE = 5
	)
	var (
		solicitaEnvio = make(chan int, BUFFER_SIZE)
		envia = make(chan int, BUFFER_SIZE)
		confirma = make(chan struct{}, BUFFER_SIZE)
	)
	go Gerador(AMOUNT_TO_GENERATE, solicitaEnvio)
	// go Fonte(solicitaEnvio, envia, confirma)
	// go FonteEx1(solicitaEnvio, envia, confirma)
	go FonteEx2(solicitaEnvio, envia, confirma, BUFFER_SIZE)
	fmt.Println()
	go Destino(envia, confirma)
	<-make(chan struct{})
}
