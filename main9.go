package main

import "fmt"

func main() {
	age := 30
	name := "John"
	value := 1000.435
	a := false

	fmt.Println("My age is " + fmt.Sprint(age))
	fmt.Printf("My age is %d. My name is %s. My money: %f, u - %t.",
		age, name, value, a)
}
