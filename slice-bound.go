package main

import "fmt"

func main() {
	a := []int{2, 3, 4, 5}
	s := a[:]
	fmt.Println(s)
}
