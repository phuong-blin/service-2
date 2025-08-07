package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Service 2: stupid string reverser")
	fmt.Println("Usage: <op> string")
	fmt.Println("where <op> can be 'reverse', 'wordcount' and string is--you guessed it--a string")

	op := os.Args[1]
	str := os.Args[2]

	switch op {
	case "reverse": 
		ans := reverse(str)
		fmt.Printf("= %s", ans)
	case "wordcount": 
		ans := wordcount(str)
		fmt.Printf("= %d", ans)
	default: 
		fmt.Println("Illegal :)")
	}
}

func reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func wordcount(s string) int {
	words := strings.Fields(s)
	return len(words)
}