package main

import "fmt"

type name string
func (n name) hello() {
	fmt.Println("Hello World")
} 

type add func(int, int) int

func (a add) minus(x, b int) int {
	return x - b
}

func sum (a, b int) int {
	return a + b
}

func main() {
	x := name(" ")
	x.hello()
	test := add(sum)
	test.minus(3, 2)
}