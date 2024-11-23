package main

import "fmt"

func main() {
	name := "John"

	switch name {
	case "Jordan":
		fmt.Println("hello Jordan")

	case "Kate":
		fmt.Println("hello Kate")

	case "John":
		fmt.Println("hellp John")

	default:
		fmt.Println("i dont know u")
	}

	number := 10

	switch {
	case number > 1:
		fmt.Println("Number is greater then 1")
		fallthrough
	case number < 11:
		fmt.Println("Number < 11")
	}

}
