package main

import "fmt"

const (
	first  = iota
	second = iota
)
const (
	third = iota
)

func main() {
	fmt.Println(first, second, third)
}
