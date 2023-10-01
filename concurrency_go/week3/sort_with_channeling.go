package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func create_partition(nums []int, n_partitions int) [][]int {
	partitions := make([][]int, n_partitions)

	len_partition := len(nums) / n_partitions
	for i := range partitions {
		start_idx := i * len_partition
		end_idx := (i + 1) * len_partition

		if end_idx > len(nums) {
			end_idx = len(nums)
		}

		if i == 3 {
			partitions[i] = nums[3*len_partition:]
		} else {
			partitions[i] = nums[start_idx:end_idx]
		}
	}

	return partitions

}

func sort(partition_nums []int, c chan []int) {
	quick_sort(partition_nums)

	c <- partition_nums
}

func swap(nums []int, current int, next int) {
	nums[current], nums[next] = nums[next], nums[current]
}

func quick_sort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	left, right := 0, len(nums)-1
	pivot := rand.Int() % len(nums)

	swap(nums, pivot, right)

	for i, _ := range nums {
		if nums[i] < nums[right] {
			swap(nums, left, i)
			left++
		}
	}
	swap(nums, left, right)

	quick_sort(nums[:left])
	quick_sort(nums[left+1:])

	return nums
}

func read_values() []int {
	reader := bufio.NewReader(os.Stdin)
	var slices []int

	fmt.Println("Please input sequences of integer separated by space")
	fmt.Printf("\n>")
	if values, _, err := reader.ReadLine(); err != nil {
		panic(fmt.Sprintf("Error when input the values: %v", err))
	} else {
		values_arr := strings.Split(string(values), " ")

		for _, v := range values_arr {
			num, _ := strconv.Atoi(v)
			slices = append(slices, num)
		}
	}
	return slices
}

func merge_sorted_slices(slices ...[]int) []int {
	result := []int{}
	for _, slice := range slices {
		result = merge(result, slice)
	}
	return result
}

func merge(a, b []int) []int {
	merged := make([]int, len(a)+len(b))
	i, j, k := 0, 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			merged[k] = a[i]
			i++
		} else {
			merged[k] = b[j]
			j++
		}
		k++
	}

	for i < len(a) {
		merged[k] = a[i]
		i++
		k++
	}

	for j < len(b) {
		merged[k] = b[j]
		j++
		k++
	}

	return merged
}

func main() {
	var num_of_goroutines = 4
	c := make(chan []int, num_of_goroutines)

	nums := read_values()
	if len(nums) <= num_of_goroutines {
		quick_sort(nums)
		fmt.Printf("This is the final sorted result: %v\n", nums)
	} else {
		partitions := create_partition(nums, num_of_goroutines)
		for i := 0; i < num_of_goroutines; i++ {
			go sort(partitions[i], c)
		}

		result := merge_sorted_slices(partitions...)

		// receive all values stored in buffered channels
		stored_channels := <-c

		fmt.Printf("This is sorted slices stored on go routines: %v\n", stored_channels)
		fmt.Printf("This is the final sorted result: %v\n", result)
	}

}
