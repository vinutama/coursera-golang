package main

import (
	"fmt"
	"slices"
)

func inputSortSlice(arr []int, userInput int) {
	counter := 0
	for {
		fmt.Println("\nPlease input your number here to insert slice or type 'X' to exit: ")

		// handle error
		_, err := fmt.Scanf("%d", &userInput)
		if err != nil || userInput == 'X' {
			return
		}

		// handle 3 initiate inputed values
		if counter < 3 {
			arr[counter] = userInput
		} else {
			arr = append(arr, userInput)
		}

		sortedArr := make([]int, len(arr))
		copy(sortedArr, arr)
		// sort slices
		slices.Sort(sortedArr)

		// just print top 3 ascending from the sortedArr
		top3 := sortedArr[:3]
		fmt.Printf("This is the result: %v", top3)
		counter++
	}
}

func main() {
	var userInput int

	//initialize length of arr slices 3
	arr := make([]int, 3)
	inputSortSlice(arr, userInput)
}
