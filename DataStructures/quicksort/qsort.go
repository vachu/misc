package quicksort

import "fmt"

func qsort(arr []int) {
	lo, hi := 0, len(arr)-1
	if len(arr) > 1 {
		pivotIndex := partitionHoare(arr)
		qsort(arr[lo : pivotIndex+1])
		qsort(arr[pivotIndex+1 : hi+1])
	}
}

// var callCount uint

func partitionHoare(arr []int) (newPivotIndex int) {
	lo, hi := 0, len(arr)-1
	pi := int((lo + hi) / 2)
	pivot := arr[pi]
	i := lo - 1
	j := hi + 1
	// fmt.Println("partition(): input arr:", arr)
	// fmt.Printf("i=%d, j=%d, pivot=%d\n", i, j, pivot)
	// defer func() {
	// 	fmt.Println("re-partitioned arr:", arr)
	// 	fmt.Printf("i=%d, j=%d\n", i, j)
	// 	fmt.Println("newPivotIndex =", newPivotIndex)
	// 	fmt.Println("-------------------------------------------------------")
	// }()

	for {
		for ok := true; ok; ok = arr[i] < pivot {
			i++
		}
		for ok := true; ok; ok = arr[j] > pivot {
			j--
		}
		if i >= j {
			newPivotIndex = j // new pivot index
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	return
}

// Run ...
func Run() {
	arr := []int{9, 9, 8, 8, 7, 7, 6, 6, 0, 0, 5, 5, 4, 4, 3, 3, 2, 2, 1, 1}
	fmt.Println("original arr:", arr)
	qsort(arr)
	fmt.Println("  sorted arr:", arr)
}
