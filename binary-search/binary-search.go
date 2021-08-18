package main

import "fmt"

func binarySearch(needle int, haystack []int) (postition int) {
	low := 0
	high := len(haystack) -1

	for low <= high {
		median := low + high / 2
		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	if low == len(haystack) || haystack[low] != needle {
		return -1 
	}

	return  low
}

func main() {
	array := []int{ 2, 5, 7, 18, 24, 28, 34, 45, 89, 100}
	position := binarySearch(18, array)
	fmt.Println(position)
}