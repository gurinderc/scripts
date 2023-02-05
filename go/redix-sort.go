package main

import (
	"fmt"
	"strconv"
)

func radixSort(arr []int) []int {
	// Find the maximum number to know the number of digits
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	// Do counting sort for every digit
	exp := 1
	for max/exp > 0 {
		bucket := make([]int, 10)
		output := make([]int, len(arr))
		for i := 0; i < len(arr); i++ {
			bucket[(arr[i]/exp)%10]++
		}
		for i := 1; i < 10; i++ {
			bucket[i] += bucket[i-1]
		}
		for i := len(arr) - 1; i >= 0; i-- {
			output[bucket[(arr[i]/exp)%10]-1] = arr[i]
			bucket[(arr[i]/exp)%10]--
		}
		for i := 0; i < len(arr); i++ {
			arr[i] = output[i]
		}
		exp *= 10
	}

	return arr
}

func main() {
	arr := []int{170, 45, 75, 90, 802, 24, 2, 66}
	fmt.Println("Sorted array is: ", radixSort(arr))
}
