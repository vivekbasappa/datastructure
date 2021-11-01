
package main

import "fmt"

func bico(n, k int) int {
	if k == 0 || k == n {
		return 1
	}
	return bico(n-1, k-1) + bico(n-1, k)
}

func main() {
	n := 5
	k := 2
	res := bico(n, k)
	fmt.Println("result:", res)
}