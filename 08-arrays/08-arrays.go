package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println(len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", c)

	c = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", c)

	var two [2][3]int
	for i := range 2 {
		for j := range 3 {
			two[i][j] = i + j
		}
	}
	fmt.Println("2d: ", two)

	two_a := [...][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2d: ", two_a)
}
