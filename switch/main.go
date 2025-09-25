package main

import (
	"fmt"
)

func main() {
	// switch time.Now().Weekday() {
	// case time.Friday:
	// 	fmt.Println("Today is Friday")
	// case time.Saturday:
	// 	fmt.Println("Today is Saturday")
	// case time.Sunday:
	// 	fmt.Println("Today is Sunday")
	// case time.Monday:
	// 	fmt.Println("Today is Monday")
	// case time.Tuesday:
	// 	fmt.Println("Today is Tuesday")
	// case time.Wednesday:
	// 	fmt.Println("Today is Wednesday")
	// case time.Thursday:
	// 	fmt.Println("Today is Thursday")
	// default:
	// 	fmt.Println("Today is a weekday")
	// }

	whAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}

	}
	whAmI(true)
	whAmI(1)
	whAmI("hey")
	whAmI(3.4)
}
