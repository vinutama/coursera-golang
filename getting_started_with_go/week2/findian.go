package main

import (
	"bufio"
	"fmt"
	"os"
	str "strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input strings here")

	var userInput string
	var rulesCount int

	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	firstChar := str.ToLower(string(userInput[0]))
	containsALetter := str.Contains(str.ToLower(userInput), "a")
	lastChar := str.ToLower(string(userInput[len(userInput)-2]))

	// check all the rules were met
	if firstChar == "i" {
		rulesCount++
	}
	if containsALetter {
		rulesCount++
	}
	if lastChar == "n" {
		rulesCount++
	}

	if rulesCount == 3 {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
