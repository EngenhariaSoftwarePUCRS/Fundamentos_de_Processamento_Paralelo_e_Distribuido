package main

import (
	"fmt"
	"os"
)

func main() {
	mydata := []byte("Meu texto de\nteste")

	err := os.WriteFile("test.txt", mydata, 0777)

	if err != nil {
		fmt.Println(err)
		return
	}
	
	data, err := os.ReadFile("test.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println(string(data))

	f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString("\nNova linha de testee"); err != nil {
		panic(err)
	}

	data, err = os.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}
