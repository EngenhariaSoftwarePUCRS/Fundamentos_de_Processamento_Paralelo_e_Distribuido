package main

import (
	"fmt"
	"time"
)

const (
	Red = "\033[31m"
	Green = "\033[32m"
	Yellow = "\033[33m"
	Reset = "\033[0m"
)

var (
	// Amount of elves currently waiting for Santa
	elves     uint8 = 0
	// Amount of reindeers currently waiting for Santa
	reindeers uint8 = 0
	// Semaphore(0) - Santa waits for either 3 elves or 9 reindeers
	santaSem = NewSemaphore(0)
	// Semaphore(0) - Reindeers wait for Santa to hitch them
	reindeerSem = NewSemaphore(0)
	// Semaphore(1) - Elves wait for other elves to be helped
	elfTex = NewSemaphore(1)
	// Semaphore(1) - Mutex for accessing reindeers and elves counters
	mutex = NewSemaphore(1)
)

type Semaphore struct {
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

func santa() {
	for {
		PrintSanta("Santa is sleeping")
		santaSem.Wait()
		PrintSanta("Santa is awake")
		PrintSanta("Santa is waiting to help elves or reindeers")
		mutex.Wait()
		PrintSanta("Santa is ready to help")
			if reindeers == 9 {
				prepareSleigh()
				for i := 0; i < 9; i++ {
					PrintSanta(fmt.Sprintf("Hitching reindeer %d", i))
					reindeerSem.Signal()
				}
			} else if elves == 3 {
				helpElves()
			}
		PrintSanta("Santa is done helping")
		mutex.Signal()
		PrintSanta("Santa is going to sleep")
	}
}

func prepareSleigh() {
	PrintSanta("Santa is preparing the sleigh")
	reindeers = 0
}

func helpElves() {
	PrintSanta("Santa is helping elves")
}

func reindeer(name string) {
	for {
		PrintReindeer(name + " is waiting")
		mutex.Wait()
		PrintReindeer(name + " is in the critical section")
			reindeers += 1
			if reindeers == 9 {
				PrintReindeer(name + " is waking santa")
				santaSem.Signal()
			}
		PrintReindeer(name + " is out the critical section")
		mutex.Signal()

		PrintReindeer(fmt.Sprintf("%s is waiting for hitch (%d/9)", name, reindeers))
		reindeerSem.Wait()
		getHitched(name)
	}
}

func getHitched(name string) {
	PrintReindeer(name + " is getting hitched")
}

func elf(id string) {
	for {
		elfTex.Wait()
		PrintElf(id + " is waiting")
		mutex.Wait()
		PrintElf(id + " is in the critical section")
			elves += 1
			if elves == 3 {
				PrintElf(fmt.Sprintf("%s is waking santa (%d/3)", id, elves))
				santaSem.Signal()
			} else {
				PrintElf(fmt.Sprintf("%s is waiting for other elves (%d/4)", id, elves))
				elfTex.Signal()
			}
		PrintElf(id + " is out the critical section")
		mutex.Signal()

		// Should wait for others
		getHelp(id)

		mutex.Wait()
		PrintElf(id + " is leaving")
			elves -= 1
			if elves == 0 {
				elfTex.Signal()
			}
		PrintElf(id + " is out")
		mutex.Signal()
	}
}

func getHelp(id string) {
	PrintElf(id + " is getting help")
}

func main() {
	reindeerNames := []string{"Dasher", "Dancer", "Prancer", "Vixen", "Comet", "Cupid", "Donder", "Blitzen", "Rudolph"}
	elfAmount := 4
	fmt.Println("Christmas is coming")
	go santa()
	for i := 0; i < len(reindeerNames); i++ {
		fmt.Println("Making reindeer ", i)
		go reindeer(reindeerNames[i])
	}
	for i := 0; i < elfAmount; i++ {
		fmt.Println("Making elf ", i)
		go elf(fmt.Sprintf("%d", i))
	}
	<-time.After(500 * time.Millisecond)
	goodbye := " Christmas is here "
	colors := []string{Red, Green, Yellow}
	for i := 0; i < len(goodbye); i++ {
		fmt.Print(colors[i%3] + "=")
	}
	fmt.Println()
	for i := 0; i < len(goodbye); i++ {
		fmt.Print(colors[i%3] + string(goodbye[i]))
	}
	fmt.Println()
	for i := 0; i < len(goodbye); i++ {
		fmt.Print(colors[i%3] + "=")
	}
	fmt.Println(Reset)
}

func PrintSanta(s string) {
	fmt.Println(Red + s + Reset)
}

func PrintReindeer(s string) {
	fmt.Println(Yellow + s + Reset)
}

func PrintElf(s string) {
	fmt.Println(Green + s + Reset)
}
