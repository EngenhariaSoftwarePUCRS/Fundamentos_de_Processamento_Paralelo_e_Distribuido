/*
	* Felipe Freitas Silva
	* 01/09/2023
	
	* 1) Modifique o número de processos e o numero de iteracoes de cada processo
	* R: Código

	* 2) Avalie o resultado obtido do ponto de vista da velocidade relativa entre os processos e da justiça.
	* R: Independentemente de velocidade, sempre todos os processos vão executar suas funções, independente de ordem

	* 4) Faria diferença se 'fin' fosse assíncrono, ou seja, se tivesse um buffer para armazenar itens?
	* R: Sim, mas não para a execução do programa, visto que ainda é necessário fazer N leituras, e o fato de uma escrita não esperar a outra não é muito importante já que o sinal de processo terminado só acontece ao final de cada processo e estes são independentes - então o programa principal (main) ainda teria que realizar N leituras.
*/


package main

import "fmt"

func algoConcorrente(id int, par int, fin chan struct{}) {
	for i := 0; i < par; i++ {
		fmt.Println(id, " fazendo algo ", i)
	}
	// Sinaliza final
	fin <- struct{}{}
}

func main() {
	const QTD_ROTINAS = 20
	fin := make(chan struct{}, QTD_ROTINAS)

	// Criar N rotinas concorrentes
	for i := 0; i < QTD_ROTINAS; i++ {
		// Repassa canal fin para avisar o termino
		go algoConcorrente(i, QTD_ROTINAS, fin)
	}

	// Espera o termino das rotinas
	for i := 0; i < QTD_ROTINAS; i++ {
		<-fin
	}

	fmt.Println("fim")
}
