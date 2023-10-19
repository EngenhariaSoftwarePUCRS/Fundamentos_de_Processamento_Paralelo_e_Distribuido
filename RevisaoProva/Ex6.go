// Suponha que você dispõe de canais e precisa de uma implementação de semáforos, usando as operações  s=NewSemaphore(_);  s.Wait()  e  s.Signal()  __conforme  visto  em  aula__. A  seguinte implementação é sugerida em algum lugar da internet e nem tudo funciona como deveria. Qual a razão ?

package main

import (
	"fmt"
	"time"
)

var i int = 0

type Semaphore struct {
	sChan chan struct{}
}
func NewSemaphore(init int) *Semaphore {
	s := &Semaphore{
		sChan: make(chan struct{}, init),	
	}
	return s
}
func (s *Semaphore) Wait() {
	fmt.Println("Wait", i)
	i++
	s.sChan <- struct{}{}
}
func (s *Semaphore) Signal() {
	<-s.sChan
}

func main() {
	s := NewSemaphore(1)
	for i := 0; i < 10; i++ {
		go func() {
			s.Wait()
			s.Signal()
		}()
	}
	fmt.Println("Fim")
	<- time.After(2 * time.Second)
}
