package main

import (
	"fmt"
	"math"
)

func main() {
	//num := 3
	//if num > 0 {
	//	fmt.Println("Number is greater then 0")
	//} else if num < 0 {
	//	fmt.Println("Number < 0")
	//} else if num == 3 {
	//	fmt.Println("number equals 3")
	//}
	var a, b, c float64

	fmt.Print("Enter a: ")
	fmt.Scan(&a)

	fmt.Print("Enter b: ")
	fmt.Scan(&b)

	fmt.Print("Enter c: ")
	fmt.Scan(&c)

	D := b*b - 4*a*c

	if D > 0 {

		var x1 float64
		var x2 float64

		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)

		fmt.Println("ваше уравнение имеет 2 корня\nD = " + fmt.Sprint(D))
		fmt.Println("X1: " + fmt.Sprint(x1) + "\nX2: " + fmt.Sprint(x2))

	} else if D == 0 {

		var x1 float64

		x1 = (-b) / (2 * a)

		fmt.Println("Уравнение имеет  1 корень\nD = " + fmt.Sprint(D))
		fmt.Println("X: " + fmt.Sprint(x1))

	} else if D < 0 {

		fmt.Println("Ваше уравнение не имеет корней\nD = " + fmt.Sprint(D))
	}

	fmt.Println("Для выход введите \"а\"")
	fmt.Scan(&a)
}
