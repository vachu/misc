package quicksort

import (
	"math/rand"
	"testing"
	"time"
)

func Test_qsort(t *testing.T) {
	for i := 0; i < 10; i++ {
		arr := generateRandomArray(10)
		t.Log("input array:", arr)
		qsort(arr)

		ok := isSorted(arr)
		t.Log("sorted array:", arr, "sorted:", ok)
		if !ok {
			t.Fail()
		}
	}
}

func generateRandomArray(size uint) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, size)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Int() % (len(arr) * 10)
	}
	return arr
}

func isSorted(arr []int) bool {
	arrLen := len(arr)
	if arrLen > 1 {
		for i, j := 0, 1; j < arrLen; i, j = j, j+1 {
			if arr[i] > arr[j] {
				return false
			}
		}
	}
	return true
}
