/*
	* Felipe Freitas Silva
	* 08/09/2023
*/

package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! - Olá, Mundo! - Hallo, Welt! - ...")
}

func help(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You can do it!")
}

func f1(w http.ResponseWriter, r *http.Request) {
	for i := 1; i < 5; i++ {
		fmt.Fprintf(w, "Computing. . .\n")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/help", help)
	http.HandleFunc("/f1", f1)

	port := 8080
	address := fmt.Sprintf(":%d", port)

	fmt.Printf("Starting server on address: localhost%s", address)
	
	http.ListenAndServe(address, nil)
}
