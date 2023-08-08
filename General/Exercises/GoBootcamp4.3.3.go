package main

import (
	"fmt"
)

var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
	"Emma", "Isabella", "Emily", "Madison",
	"Ava", "Olivia", "Sophia", "Abigail",
	"Elizabeth", "Chloe", "Samantha",
	"Addison", "Natalie", "Mia", "Alexis"}

func main() {
	var maxLen int
	for _, name := range names {
		if len(name) > maxLen {
			maxLen = len(name)
		}
	}
	result := make([][]string, maxLen)
	for _, name := range names {
		result[len(name)-1] = append(result[len(name)-1], name)
	}
	fmt.Printf("%q", result)
}
