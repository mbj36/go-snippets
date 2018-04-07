package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println(time.Now())
	fmt.Println("Hello there")
}
