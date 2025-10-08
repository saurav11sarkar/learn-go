package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	m := make(map[string]string)
	m["name"] = "Rahul"
	m["course"] = "Golang"
	m["year"] = "2023"
	m["city"] = "Delhi"
	fmt.Println(len(m))
	fmt.Println(m["name"])
	delete(m, "name")
	fmt.Println(m)
}
