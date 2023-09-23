package main

import (
	"fmt"
	"time"
)

func print_str(str string) {
	// this logical process is to print each character from inputted word within 2 goroutines
	for _, char := range str {
		fmt.Printf("Printed result: %v \n", string(char))
		// need to wait to prove one of each goroutines either 1 or 2 executed and print the result
		time.Sleep(20 * time.Millisecond)
	}
}

func main() {
	// create 2 goroutines with func go
	go print_str("ABCD")
	go print_str("WXYZ")

	// need to wait to prevent program is teminated immediately
	time.Sleep(1 * time.Second)
}
