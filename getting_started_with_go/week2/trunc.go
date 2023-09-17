package main

import (
	"fmt"
)

func main() {
	var userInput float64

	fmt.Printf("Please input your floating numbers here: ")

	_, err := fmt.Scan(&userInput)

	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	fmt.Println(int(userInput))
}
