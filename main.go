package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8}
	fmt.Println(a)
	fmt.Println(b)
	a = b
	fmt.Println(a)
}
