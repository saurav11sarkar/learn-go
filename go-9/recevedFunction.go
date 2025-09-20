package main

import "fmt"

type User struct {
	name string
	age  int
}

// Receiver function1
func recevedFunction(x User) {
	fmt.Println("Inside recevedFunction, x =", x.name)
	fmt.Println("Inside recevedFunction, x =", x.age)
}

// Receiver function2
func (x User) printDetails() {
	fmt.Println("Inside printDetails, x =", x.name)
	fmt.Println("Inside printDetails, x =", x.age)
}

// Receiver function3
func (x User) printDetails2(a int) {
	fmt.Println("Inside printDetails2, x =", x.name)
	fmt.Println("Inside printDetails2, x =", x.age)
	fmt.Println("Inside printDetails2, a =", a*x.age)
}

func main() {
	user1 := User{"Alice", 20}
	user1.printDetails()
	// recevedFunction(user1)
	user2 := User{"Bob", 30}
	recevedFunction(user2)
	user2.printDetails2(10)
}
