package main

import "fmt"

func main() {
	var n int = 234567
	a := (n & (n - 1)) ^ n

	fmt.Println(a)
}
