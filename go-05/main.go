package main

import "fmt"

const a = 100

var b = 200

func call() {
	add := func(x, y int) {
		z := x + y
		fmt.Println("z =", z)
	}
	add(5, 6)
	add(a, b)
}

func main() {
	a := 500
	fmt.Println("Hello, Go!")
	call()
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}

func init() {
	fmt.Println("init()...")
}
