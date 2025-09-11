package main

import "fmt"

var a int = 10
var b int = 20

func add(x int, y int) {
	z := x + y
	fmt.Println("Inside add function, z =", z)
}

func main() {
	p := 30
	q := 40

	add(p, q)
	add(a, b)

	add(a, p)

}
