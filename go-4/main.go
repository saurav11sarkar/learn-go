package main

var a int = 20
var b int = 30

func add(x int, y int) {
	c := x + y
	printNumber(c)
}

func main() {
	println("Hello, Go 4!")
	add(a, b)
	unusedFunction()
}

func printNumber(num int) {
	println("Number:", num)
}

func unusedFunction() {
	println("This function is not used.")

	age := 25
	if age > 18 {
		a := "Adult"
		println(a)
	}
}
