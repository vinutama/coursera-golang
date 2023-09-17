package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func swap(sli []int, idx int) {
	sli[idx], sli[idx-1] = sli[idx-1], sli[idx]
}

func swap2(sli []int, idx int) {
	sli[idx], sli[idx+1] = sli[idx+1], sli[idx]
}

func insertion_sort(slice []int) {
	// for i := 0; i < len(slice); i++ {
	// 	for j := 0; j < (len(slice))-1; j++ {
	// 		if slice[j] > slice[j+1] {
	// 			swap(slice, j)
	// 		}
	// 	}
	// }

	for i := 1; i < len(slice); i++ {
		j := i
		for j > 0 && slice[j] < slice[j-1] {
			swap(slice, j)
			j--
		}
	}

	// for i := len(slice) - 2; i >= 0; i-- {
	// 	j := i

	// 	for j <= len(slice)-2 && slice[j] > slice[j+1] {
	// 		swap2(slice, j)
	// 		j++
	// 	}
	// }

}

func main() {
	// change this caps
	caps := 10
	slic := read_values(caps)

	insertion_sort(slic)

	fmt.Println(slic)
}

func read_values(caps int) []int {
	reader := bufio.NewReader(os.Stdin)

	slice := make([]int, 0, caps)

	fmt.Println("Please input sequences of integer separated by space")
	if values, _, err := reader.ReadLine(); err != nil {
		fmt.Printf("Error when input the values: %v", err)
	} else {
		values_arr := strings.Split(string(values), " ")
		if len(values_arr) > caps {
			panic(fmt.Sprintf("Input cannot exceed %v items", caps))
		}

		for _, v := range values_arr {
			num, _ := strconv.Atoi(v)
			slice = append(slice, num)
		}
	}

	return slice
}
