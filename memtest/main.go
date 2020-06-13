package main

import "fmt"

func main() {
	const SIZE = 1_00_000_000

	//var arr [SIZE]uint
	arr := make(map[int]uint)
	for i := 0; i < SIZE; i++ {
		arr[i] = uint(0xffff_ffff_ffff_ffff)
	}

	fmt.Print("Press <Enter> to quit ... ")
	
	dummyInput := ""
	fmt.Scanln(&dummyInput)
}