package main

import "fmt"

var (
	elves          uint8 = 0
	reindeers uint8 = 0
	// Semaphore(0)
	santaSem = NewSemaphore(0)
	// Semaphore(0)
	reindeerSem = NewSemaphore(0)
	// Semaphore(1)
	elfTex = NewSemaphore(1)
	// Semaphore(1)
	mutex = NewSemaphore(1)
)

type Semaphore struct {
	channel chan struct{}
}

func NewSemaphore(initialSize int) *Semaphore {
	s := Semaphore{make(chan struct{}, 1)}
	for i := 0; i < initialSize; i++ {
		s.signal()
	}
	return &s
}

func (s Semaphore) wait() {
	<-s.channel
}

func (s Semaphore) signal() {
	s.channel <- struct{}{}
}

func santa() {
	for {
		santaSem.wait()
		mutex.wait()
			if reindeers == 9 {
				fmt.Println("Preparing sleigh")
				// prepareSleigh()
				for i := 0; i < 9; i++ {
					reindeerSem.signal()
				}
			} else if elves == 3 {
				fmt.Println("Helping elves")
				// helpElves()
			}
		mutex.signal()
	}
}

func reindeer(name string) {
	mutex.wait()
		reindeers += 1
		if reindeers == 9 {
			santaSem.signal()
		}
	mutex.signal()

	reindeerSem.wait()
	fmt.Printf("%s is getting hitched\n", name)
	// getHitched()
}

func elf(id int) {
	elfTex.wait()
	mutex.wait()
		elves += 1
		if elves == 3 {
			santaSem.signal()
		} else {
			elfTex.signal()
		}
	mutex.signal()

	fmt.Printf("%d is getting help\n", id)
	// getHelp()

	mutex.wait()
		elves -= 1
		if elves == 0 {
			elfTex.signal()
		}
	mutex.signal()
}

func main() {
	reindeerAmount := 9
	reindeerNames := []string{"Dasher", "Dancer", "Prancer", "Vixen", "Comet", "Cupid", "Donder", "Blitzen", "Rudolph"}
	elfAmount := 4
	go santa()
	for i := 0; i < reindeerAmount; i++ {
		go reindeer(reindeerNames[i])	
	}
	for i := 0; i < elfAmount; i++ {
		go elf(i)
	}
}
