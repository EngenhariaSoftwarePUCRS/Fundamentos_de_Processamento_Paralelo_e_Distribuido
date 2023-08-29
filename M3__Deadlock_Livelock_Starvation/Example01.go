package main

import "fmt"

func proc(s string, rx, ry chan struct{}) {
	for {
		<-rx
		<-ry
		rx <- struct{}{}
		ry <- struct{}{}
		fmt.Print(s)
	}
}

func main() {
	r1 := make(chan struct{}, 1)
	r2 := make(chan struct{}, 1)
	r1 <- struct{}{}
	r2 <- struct{}{}
	go proc("|", r1, r2)
	go proc("-", r2, r1)
	<-make(chan struct{})
}
