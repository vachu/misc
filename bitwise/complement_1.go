package main

import "fmt"

func main() {
	var num uint8 = 0b0111_1111
	fmt.Printf("%d, %08b\n", num, num)

	num = ^num
	fmt.Printf("%d, %08b\n", num, num)
}
