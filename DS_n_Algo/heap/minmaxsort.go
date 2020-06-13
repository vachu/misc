package heap

import "fmt"

// Type ...
type Type uint8

// HeapType ...
const (
	MinHeap Type = 0
	MaxHeap Type = 1
)

// Heapify ...
func Heapify(arr []int, heapType Type) {
	arrLen := len(arr)
	if arrLen >= 2 {
		mid := int(arrLen/2) - 1
		switch heapType {
		case MinHeap:
			minify(arr, mid)
		case MaxHeap:
			maxify(arr, mid)
		}
	}
}

func minify(arr []int, idxRoot int) {
	arrLen := len(arr)
	idxSmallest := idxRoot
	idxLeftChild := idxRoot*2 + 1
	idxRightChild := idxRoot*2 + 2

	if idxLeftChild < arrLen && arr[idxLeftChild] < arr[idxSmallest] {
		idxSmallest = idxLeftChild
	}
	if idxRightChild < arrLen && arr[idxRightChild] < arr[idxSmallest] {
		idxSmallest = idxRightChild
	}

	if idxSmallest != idxRoot {
		arr[idxRoot], arr[idxSmallest] = arr[idxSmallest], arr[idxRoot]
		minify(arr, idxSmallest)
	}
	if idxRoot > 0 {
		minify(arr, idxRoot-1)
	}
}

func maxify(arr []int, idxRoot int) {
	arrLen := len(arr)
	idxLargest := idxRoot
	idxLeftChild := idxRoot*2 + 1
	idxRightChild := idxRoot*2 + 2

	if idxLeftChild < arrLen && arr[idxLeftChild] > arr[idxLargest] {
		idxLargest = idxLeftChild
	}
	if idxRightChild < arrLen && arr[idxRightChild] > arr[idxLargest] {
		idxLargest = idxRightChild
	}

	if idxLargest != idxRoot {
		arr[idxRoot], arr[idxLargest] = arr[idxLargest], arr[idxRoot]
		maxify(arr, idxLargest)
	}
	if idxRoot > 0 {
		maxify(arr, idxRoot-1)
	}
}

// IsMaxHeap ...
func IsMaxHeap(arr []int) bool {
	for arrLen, idxRoot := len(arr), 0; idxRoot <= arrLen/2; idxRoot++ {
		idxLeftChild := idxRoot*2 + 1
		idxRightChild := idxRoot*2 + 2
		if idxLeftChild < arrLen && arr[idxLeftChild] > arr[idxRoot] {
			return false
		}
		if idxRightChild < arrLen && arr[idxRightChild] > arr[idxRoot] {
			return false
		}
	}
	return true
}

// IsMinHeap ...
func IsMinHeap(arr []int) bool {
	for arrLen, idxRoot := len(arr), 0; idxRoot <= arrLen/2; idxRoot++ {
		idxLeftChild := idxRoot*2 + 1
		idxRightChild := idxRoot*2 + 2
		if idxLeftChild < arrLen && arr[idxLeftChild] < arr[idxRoot] {
			return false
		}
		if idxRightChild < arrLen && arr[idxRightChild] < arr[idxRoot] {
			return false
		}
	}
	return true
}

// Sort ...
func Sort(arr []int) {
	origArr := arr
	for ; len(arr) >= 2; arr = arr[1:] {
		Heapify(arr, MinHeap)
		fmt.Printf("len=%d: %v\n", len(arr), origArr)
	}
}

// Run ...
func Run() {
	const SIZE = 20
	var arr [SIZE]int
	for i := 0; i < SIZE; i++ {
		arr[i] = SIZE - i // in descending
		// arr[i] = i + 1 // in ascending
	}
	fmt.Println("Original arr:", arr)
	// Heapify(arr, MinHeap)
	Sort(arr[:])
	fmt.Println("Heapified arr:", arr)
	fmt.Println("IsMaxHeap:", IsMaxHeap(arr[:]))
	fmt.Println("IsMinHeap:", IsMinHeap(arr[:]))
}
