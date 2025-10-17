// package main

// import "fmt"

// func main() {
// 	fmt.Println("Enter your Roll: ")
// 	var roll string
// 	fmt.Scanln(&roll)
// 	fmt.Println("Enter your Password: ")
// 	var password string
// 	fmt.Scanln(&password)

// 	if roll == "admin" && password == "admin123" {
// 		fmt.Println("Welcome Admin")
// 	} else {
// 		for range 3 {
// 			fmt.Println("Invalid Credentials")
// 		}
// 	}
// }

package main

import "fmt"

func main() {
	fmt.Println("Enter your Roll: ")
	var roll int
	fmt.Scanln(&roll)

	switch roll {
	case 1:
		fmt.Println("Enter your Password: ")
		var password string
		fmt.Scan(&password)
		if password == "admin123" {
			fmt.Println("Welcome Admin")
		} else {
			for range 3 {
				fmt.Println("Invalid Credentials")
			}
		}
	case 2:
		fmt.Println("Enter your Password: ")
		var password string
		fmt.Scan(&password)
		if password == "user123" {
			fmt.Println("Welcome User")
		} else {
			for range 3 {
				fmt.Println("Invalid Credentials")
			}
		}
	default:
		fmt.Println("Invalid Roll")
	}
}
