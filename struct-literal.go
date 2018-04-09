package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 3}
	q := Vertex{X: 1}
	p := &Vertex{1, 2}

	fmt.Println(p, v, q)
}
