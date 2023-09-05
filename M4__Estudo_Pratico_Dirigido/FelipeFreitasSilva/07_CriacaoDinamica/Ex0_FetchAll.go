/*
	* Felipe Freitas Silva
	* 04/09/2023

	* 1) Faça o fetch de cada endereco em sequencia e veja a diferença para a versao concorrente.
	* R: Código

	* 2) Execute varias vezes cada uma (aprox 5) e faça uma media dos valores sequenciais e concorrentes e veja se há ganho com a concorrencia (tempo conc / tempo seq)
	* R: 
*/

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Nenhum argumento passado")
		fmt.Println("Por favor, passe um ou mais argumentos\n Exemplo: ")
		fmt.Println("go run Ex0_FetchAll.go https://golang.org http://gopl.io https://godoc.org https://www.google.com.br https://www.youtube.com")
		return
	}

	// Sequential proccess
	start := time.Now()
	
	syncChan := make(chan string)
	for _, url := range args {
		fetchSync(url, syncChan)
	}

	for range args {
		data := <-syncChan
		appendToFile("syncOutput.txt", data)
		fmt.Print(data)
	}

	result := fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())
	appendToFile("syncOutput.txt", result)
	fmt.Printf("Sync - %s", result)
	
	// // Concurrent proccess
	// start := time.Now()

	// concurrentChan := make(chan string)
	// for _, url := range args {
	// 	go fetch(url, concurrentChan)
	// }

	// for range args {
	// 	data := <-concurrentChan
	// 	appendToFile("concurrentOutput.txt", data)
	// 	fmt.Print(data)
	// }

	// result := fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())
	// appendToFile("concurrentOutput.txt", result)
	// fmt.Printf("Concurrent - %s", result)
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("\nConcurrent - %s", err)
		return
	}
	
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	// Don't leak resources
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("\nConcurrent - while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s\n", secs, nbytes, url)
}

func fetchSync(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("\nSync - %s", err)
		return
	}
	
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	// Don't leak resources
	resp.Body.Close()
	if err != nil {
		fmt.Printf("\nSync - while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("\nSync - %.2fs  %7d  %s\n", secs, nbytes, url)
}

// If file doesn't exist, creates it
func appendToFile(fileName string, data string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	
	// If file is empty, write header
	if stat, err := f.Stat(); err == nil && stat.Size() == 0 {
		header := fmt.Sprintf("%7s  %7s  %s", "tempo", "bytes", "url", "\n")
		if _, err := f.WriteString(header); err != nil {
			fmt.Println(err)
		}
	}

	if _, err := f.WriteString(data); err != nil {
		fmt.Println(err)
	}
}
