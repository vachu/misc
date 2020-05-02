package main

import (
	"os"

	"github.com/vachu/avlbst"
)

//
func main() {
	bst := avlbst.NewBinarySearchTree(false, func(d1, d2 interface{}) int {
		v1, v2 := d1.(int), d2.(int)
		if v1 < v2 {
			return -1
		} else if v1 == v2 {
			return 0
		} else {
			return 1
		}
	})

	for _, n := range []int{1, 2, 3, 4, 5, 5, 5, 5} {
		bst.Add(n)
	}
	bst.ToXML(os.Stdout)
}
