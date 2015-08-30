package main

import "fmt"

func main() {
	buf := make([]byte, 0, 10)
	fmt.Println("buf =", buf)

	data := []byte{'a', 'b', 'c'}
	buf = append(buf, data...)
	fmt.Println("buf =", buf)
}
