package main

func processOptation(x int, y int, op func(a int, b int)) {
	op(x, y)
}

func add(a int, b int) {
	println("Add:", a+b)
}

func main() {
	println("Hello, World!")
	processOptation(10, 20, add)
}
