package main

import "fmt"

func main() {
	a := 1

	if a > 1 {
		fmt.Println("a > 0")
	}

	if a >= 1 {
		fmt.Println("a > 0")
	}

	b := 3
	c := 10

	if b > 1 && c > 5 {
		fmt.Println("True!")
	}

	if b > 1 || c > 11 {
		fmt.Println("True!")
	}

	if b != 2 && c > 5 || b > 6 {
		fmt.Println("Yeds!")
	}
}
