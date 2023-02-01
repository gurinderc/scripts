package main

import "fmt"

func mergeSort(array []int) []int {
    if len(array) <= 1 {
        return array
    }

    middle := len(array) / 2
    leftArray := array[:middle]
    rightArray := array[middle:]

    leftArray = mergeSort(leftArray)
    rightArray = mergeSort(rightArray)

    return merge(leftArray, rightArray)
}

func merge(leftArray, rightArray []int) []int {
    sortedArray := make([]int, len(leftArray) + len(rightArray))
    i, j, k := 0, 0, 0

    for i < len(leftArray) && j < len(rightArray) {
        if leftArray[i] < rightArray[j] {
            sortedArray[k] = leftArray[i]
            i++
        } else {
            sortedArray[k] = rightArray[j]
            j++
        }
        k++
    }

    for i < len(leftArray) {
        sortedArray[k] = leftArray[i]
        i++
        k++
    }

    for j < len(rightArray) {
        sortedArray[k] = rightArray[j]
        j++
        k++
    }

    return sortedArray
}

func main() {
    unsortedArray := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
    fmt.Println("Unsorted Array:", unsortedArray)
    fmt.Println("Sorted Array:", mergeSort(unsortedArray))
}
