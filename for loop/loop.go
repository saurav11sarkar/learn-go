package main

import "fmt"

func main() {
	// i := 1
	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	for i := range 5 {
		fmt.Println(i)
	}
}
