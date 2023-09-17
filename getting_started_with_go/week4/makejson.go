package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	type Person map[string]string
	rMap := make(Person)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please input your name")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Type error when input name: %v", err)
	}

	fmt.Println("Please input your address")
	addr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Type error when input address: %v", err)
	}

	rMap["name"] = strings.Trim(name, "\n")
	rMap["address"] = strings.Trim(addr, "\n")

	res, err := json.Marshal(rMap)
	if err != nil {
		fmt.Printf("An error occured when serialize JSON: %v", err)
	}
	fmt.Printf("This is the JSON Object result: %v", string(res))
}
