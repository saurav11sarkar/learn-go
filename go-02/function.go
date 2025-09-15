// package main

// import "fmt"

// func add(num1 int, num2 int) int {
// 	return num1 + num2
// }

// func main() {
// 	a := 100
// 	b := 200

// 	sum := a + b
// 	fmt.Println("Sum =", sum)

// 	diff := a - b
// 	fmt.Println("Difference =", diff)

// 	result := add(20, 40)
// 	fmt.Println("Result from function call =", result)
// }

// package main

// import "fmt"

// func main() {
// 	fmt.Println("hello wellcome to application development")

// 	var name string
// 	fmt.Print("Enter your name: ")
// 	fmt.Scanln(&name)
// 	fmt.Println("Hello", name)

// }

package main

import "fmt"

func printWellcomeMessage() {
	fmt.Println("hello wellcome to application development")
}

func addTwoNumbers(num1 int, num2 int) int {
	return num1 + num2
}

func addTwoNumberScan() (int, int) {
	var num1, num2 int
	fmt.Print("Enter two First Number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter two secend Number: ")
	fmt.Scanln(&num2)
	return num1, num2
}


func main() {
	printWellcomeMessage()
	num1, num2 := addTwoNumberScan()
	sum := addTwoNumbers(num1, num2)
	fmt.Println("Add Two Number:", sum)
}
