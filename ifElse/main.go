package main

import "fmt"

func main() {
	fmt.Println("Enter your Roll: ")
	var roll string
	fmt.Scanln(&roll)
	fmt.Println("Enter your Password: ")
	var password string
	fmt.Scanln(&password)

	if roll == "admin" && password == "admin123" {
		fmt.Println("Welcome Admin")
	} else {
		for range 3 {
			fmt.Println("Invalid Credentials")
		}
	}
}
