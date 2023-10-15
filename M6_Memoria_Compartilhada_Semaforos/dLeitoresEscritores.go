package main

import (
	"fmt"
	"os"
	"time"
)

const (
	Red = "\033[31m"
	Green = "\033[32m"
	Yellow = "\033[33m"
	Reset = "\033[0m"
)

var (
	readSwitch = Lightswitch{
		counter: 0,
		mutex:   NewSemaphore(1),
	}
	roomEmpty = NewSemaphore(1)
	turnstile = NewSemaphore(1)
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

func (s *Semaphore) Signal(msg ...string) {
	s.sc <- struct{}{} // entra sc
	s.v++
	if s.v <= 0 { // tem processo bloqueado ?
		<-s.fila // desbloqueia
	}
	<-s.sc // libera SC para outra op
	if len(msg) > 0 {
		PrintRed(msg[0])
	}
}

type Lightswitch struct {
	counter int
	mutex   *Semaphore
}

func (l *Lightswitch) Lock(semaphore *Semaphore, name string) {
	l.mutex.Wait()
		l.counter++
		if l.counter == 1 {
			semaphore.Wait()
			fmt.Println(name, "bloqueou a sala")
		}
	l.mutex.Signal()
}

func (l *Lightswitch) Unlock(semaphore *Semaphore, name string) {
	l.mutex.Wait()
		l.counter--
		if l.counter == 0 {
			semaphore.Signal()
			fmt.Println(name, "desbloqueou a sala")
		}
	l.mutex.Signal()
}

func writer(name string) {
	for {
		turnstile.Wait()
			roomEmpty.Wait()
			// Critical section for writers
			write(name)
		turnstile.Signal()
	
		roomEmpty.Signal(fmt.Sprint(name, " saiu da sala"))
	}
}

func write(id string) {
	PrintWriter(fmt.Sprint(id, "==> started writing to file"))
	
	file, err := os.OpenFile("demo.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if isError(err) {
		return
	}
	defer file.Close()

	curTime := time.Now().Local().Format("2006/01/02 15:04:05 (-0700)")
	newFileContent := fmt.Sprintln(id, ",", curTime)

	// Write some text line-by-line to file.
	_, err = file.WriteString(newFileContent)
	if isError(err) {
		return
	}

	PrintWriter(fmt.Sprint(id, "==> done writing to file"))
}

func reader(id int) {
	for {
		turnstile.Wait()
		turnstile.Signal()
	
		readSwitch.Lock(roomEmpty, fmt.Sprint(id))
			// Critical section for readers
			criticalSectionContent := read(id)
			if len(criticalSectionContent) > 50 {
				criticalSectionContent = criticalSectionContent[:47] + "..."
			}
			PrintReader(fmt.Sprint(
				id,
				" read from file:\n",
				string(criticalSectionContent),
			))
		readSwitch.Unlock(roomEmpty, fmt.Sprint(id))
	}
}

func read(id int) string {
	PrintReader(fmt.Sprint(id, "==> started reading from file"))

	data, err := os.ReadFile("demo.txt")

	if isError(err) {
		return Red + err.Error() + Reset
	}

	return string(data)
}

func main() {
	readerAmount := 100
	writers := []string{"Edgar Alan Poe", "Stephen King", "H.P. Lovecraft", "J.R.R. Tolkien", "George R.R. Martin", "J.K. Rowling", "Neil Gaiman", "Anne Rice", "Mary Shelley", "Bram Stoker"}
	ResetFile()
	for i := 0; i < readerAmount; i++ {
		go reader(i)
	}
	for _, w := range writers {
		go writer(w)
	}
	<-time.After(1000 * time.Millisecond)
}

func ResetFile() {
	defaultData := []byte("autor , data\n")

	err := os.WriteFile("demo.txt", defaultData, 0777)

	if isError(err) {
		return
	}

	fmt.Println("==> done resetting file")
}

func PrintWriter(msg string) {
	fmt.Println(Green, msg, Reset)
}

func PrintReader(id string) {
	fmt.Println(Yellow, id, Reset)
}

func PrintRed(msg string) {
	fmt.Println(Red, msg, Reset)
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
