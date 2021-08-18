// This is example code written based on Junmin Lee video tutorial
// All credit goes to her. I am just trying to understand datastructures and
// implemention in go

package main

import (
	"fmt"
)

// MaxHeap
	type MaxHeap struct {
		array []int
	}

// Insert
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array)-1)

}

// Extract
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]

	if len(h.array) == 0 {
		fmt.Println("heap is empty")
		return -1
	}

	l := len(h.array) -1
	h.array[0] = h.array[l-1]
	h.array = h.array[:l]
	h.maxHeapifyDown(0)

	return extracted
}

// maxHeapifyUp
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifyDown
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) -1
	l, r := left(index), right(index)
	childToCompare := 0
	// loop while index has at least one child
	for l <= lastIndex  {
		// when left child is the only child
		// when left child is larger
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return 
		}
		// when right child is larger
	}


}
// swap
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

// parent
func parent(index int) int {
	return (index-1)/2	
}
// left
func left(index int) int {
	return (index * 2) + 1
}
// rigth
func right (index int) int {
	return (index * 2) + 2
}

func main() {
	h := &MaxHeap{}
	fmt.Println(h)
	
	buildHeap := []int{10, 20, 30, 5, 7, 9 , 11, 13, 15, 17}
	for _, v := range buildHeap {
		h.Insert(v)
		fmt.Println(h)
	}

	for i := 0 ; i < 5 ; i++{
		h.Extract()
		fmt.Println(h)
	}
}