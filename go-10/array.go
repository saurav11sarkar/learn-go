package main

import "fmt"

var arr2 = [3]string{"a", "b", "c"}

func main() {
	var arr [2]int
	arr[1] = 2
	arr[0] = 1
	fmt.Println(arr)

	fmt.Println(arr2[1])
}
