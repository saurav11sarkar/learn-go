package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {
	var a int
	var b User

	b = User{"Alice", 20}
	a = 10

	fmt.Println("Hello, World!")
	fmt.Println(a)
	fmt.Println(b)
}
