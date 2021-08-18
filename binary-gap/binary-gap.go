package main

import (
	"strconv"
	"fmt"
)

func main() {
	N := 1041
	binary := strconv.FormatInt(int64(N), 2)
	fmt.Println(binary)
	largestCap := 0
	currentCap := 0
	startCount := false
	for _, r := range binary {
		if startCount {
			currentCap++
		}
		if string(r) == "1" {
			startCount = true
			if largestCap < currentCap {
				largestCap = currentCap
			}
			currentCap = 0
		}
	}
	fmt.Printf("result:%d\n", largestCap)
}

