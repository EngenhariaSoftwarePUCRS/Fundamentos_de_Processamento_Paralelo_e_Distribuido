/*
	* Felipe Freitas Silva
	* 01/09/2023
	
	* 1) Avalie o comportamento do programa para tamBuff 0 e 10. Você consegue explicar a diferença?
	* R: Quando existe um buffer de tamanho 10, por exemplo, a saída é aletória, pois depende da velocidade de cada processo e há uma certa aleatoriedade nisso. Já quando não há este buffer, existe o padrão de sempre cada processo executar sua ação (imprimir texto), ler/escrever este texto, executar sua ação e esperar a liberação da leitura/escrita, o que gera este ciclo "fechado".

	* 2) Qual versão tem maior nível de concorrência?
	* R: Quanto maior o tamanho do buffer, maior a possibilidade de haver mais concorrência entre os processos, visto que não há tanta dependência quanto quando o buffer é 0 (canais sincronizantes)
	
	* 3) Faça uma versão que tem vários processos destino que podem consumir os dados de forma não determinística. Ou seja, processos diferentes podem consumir quantidades diferentes de itens, conforme sua velocidade. Como você coordenaria o término dos processos depois do consumo dos N valores?
	* R: Código
*/


package main

import "fmt"

func fonteDeDados(quantidadeDados int, saida chan<- int) {
	for i := 1; i < quantidadeDados; i++ {
		fmt.Println(i, "fonteDeDados -> c ")
		saida <- i
	}
}

func main() {
	const (
		QUANTIDADE_DADOS = 100
		// + 10 para questão 1
		TAMANHO_BUFFER = 0 // + 10
	)
	c := make(chan int, TAMANHO_BUFFER)
	go fonteDeDados(QUANTIDADE_DADOS, c)
	
	// Questão 3
	for i := 1; i < QUANTIDADE_DADOS; i++ {
		select {
		case x := <- c:
			fmt.Println("x <- c ", x)

		case y := <- c:
			fmt.Println("y <- c ", y)

		case z := <- c:
			fmt.Println("z <- c ", z)
		}
	}
}
