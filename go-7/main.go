package main

import "fmt"

const a = 10

var p = 20

func outer() func() {
	mamory := 100
	age := 18

	fmt.Println("outer function called!", age)

	show := func() {
		mamory = mamory + a + p
		fmt.Println("memory:", mamory)
	}
	return show
}

func call() {
	show1 := outer()
	show1()
	show1()
	show1()

	show2 := outer()
	show2()
	show2()

}

func main() {
	call()
}
