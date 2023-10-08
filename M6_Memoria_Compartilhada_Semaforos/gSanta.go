package main

import "fmt"

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

type Semaphore chan struct{}

func NewSemaphore(initialSize int) *Semaphore {
	var s Semaphore = make(chan struct{}, 1)
	for i := 0; i < initialSize; i++ {
		s.signal()
	}
	return &s
}

func (s Semaphore) wait() {
	<-s
}

func (s Semaphore) signal() {
	s <- struct{}{}
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
