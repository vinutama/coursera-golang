package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const MAXLEN_NAME int = 20

type Person struct {
	Fname string
	Lname string
}

func inputFileName() string {
	reader := bufio.NewReader(os.Stdin)
	fname, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error when read filename input: %v", err)
	}

	return fname
}

func checkStringsMaxLen(str string) string {
	if len(str) > MAXLEN_NAME {
		return str[0:MAXLEN_NAME]
	}
	return str
}

func readFileAndParse(fname string) []Person {
	file, err := os.Open(fname)
	result := make([]Person, 0)
	if file, err := os.Open(fname); err != nil {
		fmt.Printf("Error when open the file: %v", err)
	} else {
		fileScanner := bufio.NewScanner(file)
		fileScanner.Split(bufio.ScanLines)

		for fileScanner.Scan() {
			line := checkStringsMaxLen(fileScanner.Text())
			fullname := strings.Split(line, " ")
			result = append(result, Person{Fname: fullname[0], Lname: fullname[1]})
		}
	}

	// close file
	if err = file.Close(); err != nil {
		fmt.Printf("Error occured when closing the file: %v", err)
	}

	return result
}

func main() {
	fmt.Println("Please input your raw file name here: ")
	fname := strings.Trim(inputFileName(), "\n")
	result := readFileAndParse(fname)

	// print final result
	printResult(result)

}

func printResult(result []Person) {
	for _, p := range result {
		fmt.Printf("First name: %v, Last name: %v \n", p.Fname, p.Lname)
	}
}
