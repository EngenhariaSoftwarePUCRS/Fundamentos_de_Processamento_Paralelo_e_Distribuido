// (4 pontos) Suponha que você tem os processos pA pB pC pD abaixo que ciclicamente fazem computação local, representada por "...", e então realizam respectivamente as atividades A, B, C e D que devem manter a seguinte relação de ordem: A antes de B e C; B e C concorrentes; B antes de D; C antes de D; D deve acabar antes de A fazer um novo ciclo. Ou seja, manter a precedência dada pelo grafo abaixo.

package main

import (
	"fmt"
	"time"
)

var (
	semA = NewSemaphore(1)
	semBC = NewSemaphore(0)
	renBC = NewSemaphore(0)
	renCB = NewSemaphore(0)
	semD = NewSemaphore(0)
)

func pA() {
	for {
		semA.Wait()
		fmt.Println("atividadeA()")
		semBC.Signal()
		semBC.Signal()
	}
}

func pB() {
	for {
		semBC.Wait()
		fmt.Println("atividadeB()")
		renBC.Signal()
		renCB.Wait()
		semD.Signal()
	}
}

func pC() {
	for {
		semBC.Wait()
		fmt.Println("atividadeC()")
		renCB.Signal()
		renBC.Wait()
		semD.Signal()
	}
}

func pD() {
	for {
		semD.Wait()
		semD.Wait()
		fmt.Println("atividadeD()")
		semA.Signal()
	}
}

func main() {
	go pA()
	go pB()
	go pC()
	go pD()
	<- time.After(10 * time.Second)
}

type Semaphore struct { // este semáforo implementa quaquer numero de creditos em "v"
	v    int           // valor do semaforo: negativo significa proc bloqueado
	fila chan struct{} // canal para bloquear os processos se v < 0
	sc   chan struct{} // canal para atomicidade das operacoes wait e signal
}

func NewSemaphore(init int) *Semaphore {
	s := &Semaphore{
		v:    init,                   // valor inicial de creditos
		fila: make(chan struct{}),    // canal sincrono para bloquear processos
		sc:   make(chan struct{}, 1), // usaremos este como semaforo para SC, somente 0 ou 1
	}
	return s
}

func (s *Semaphore) Wait() {
	s.sc <- struct{}{} // SC do semaforo feita com canal
	s.v--              // decrementa valor
	if s.v < 0 {       // se negativo era 0 ou menor, tem que bloquear
		<-s.sc               // antes de bloq, libera acesso
		s.fila <- struct{}{} // bloqueia proc
	} else {
		<-s.sc // libera acesso
	}
}

func (s *Semaphore) Signal() {
	s.sc <- struct{}{} // entra sc
	s.v++
	if s.v <= 0 { // tem processo bloqueado ?
		<-s.fila // desbloqueia
	}
	<-s.sc // libera SC para outra op
}
