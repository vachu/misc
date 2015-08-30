package main

import "fmt"

func main() {
	i, j := 5, 10
	fmt.Println(i, j)
	//i, j = swap(i, j)
	i, j = j, i
	fmt.Println(i, j)
}

func swap(i, j int) (int, int) {
	return j, i
}
