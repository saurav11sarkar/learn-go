package main

import "fmt"

func main() {
	// fmt.Println("Hello, World!")
	// var a int = 10
	// b := 20
	// fmt.Println("a + b =", a+b)

	// age := 18
	// gender := "fimale"

	// if(age > 18) {
	// 	fmt.Println("You are an adult.")
	// }else if(age == 18){
	// 	fmt.Println("You are just an adult.")
	// }else{
	// 	fmt.Println("You are a minor.")
	// }

	// if(age > 18 && gender == "male") {
	// 	fmt.Println("You are and adult")
	// }else if(age == 18 && gender == "male") {
	// 	fmt.Println("You are just an adult")
	// }else if(age < 18 || gender == "fimale") {
	// 	fmt.Println("You are a minor")
	// }else{
	// 	fmt.Println("You are special")
	// }

	// isPresent := false

	// if(!isPresent) {
	// 	fmt.Println("You are absent")
	// }else{
	// 	fmt.Println("You are present")
	// }

	a := 300
	switch a {
	case 1:
		fmt.Println("One")
	case 2, 3:
		fmt.Println("Two or Three")
	case 4:
		fmt.Println("Three")
	default:
		fmt.Println("Other number")
	}
}