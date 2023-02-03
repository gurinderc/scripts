package main

import (
	"fmt"
)

func heapify(numbers []int, n int, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && numbers[left] > numbers[largest] {
		largest = left
	}

	if right < n && numbers[right] > numbers[largest] {
		largest = right
	}

	if largest != i {
		numbers[i], numbers[largest] = numbers[largest], numbers[i]
		heapify(numbers, n, largest)
	}
}

func heapSort(numbers []int) {
	n := len(numbers)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(numbers, n, i)
	}

	for i := n - 1; i >= 0; i-- {
		numbers[0], numbers[i] = numbers[i], numbers[0]
		heapify(numbers, i, 0)
	}
}

func main() {
	numbers := []int{3, 2, 1, 4, 5}
	heapSort(numbers)
	fmt.Println(numbers)
	// Output: [1 2 3 4 5]
}
