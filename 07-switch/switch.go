package main

import (
	"fmt"
	"time"
)

func main() {
	i := 3
	fmt.Print(i, " is ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("not one, two, three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is weekend")
	default:
		fmt.Println("It is weekday")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// if t.Hour() < 12 {
	// 	fmt.Println("It's before noon")
	// } else {
	// 	fmt.Println("It's after noon")
	// }

	f := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	f(true)
	f(4)
	f("test")
}
