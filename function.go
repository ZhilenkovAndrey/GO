package main

import "fmt"

func main() {
	Greet()
	PersonGreet("Vasya")
	FioGreet("John", "Smith")
	ResultPrinter1(Sum(3, 4))
	ResultPrinter2(SumAndMyltyply(3, 4))
	ResultPrinter2(namedSumAndMyltiply(4, 5))
}

func Greet() {
	fmt.Println("Hello guys")
}

func PersonGreet(name string) {
	fmt.Printf("Zdarova %s\n", name)
}

func FioGreet(name, surname string) {
	fmt.Printf("Hi %s %s\n", name, surname)
}

func Sum(first, second int) int {
	return first + second
}

func ResultPrinter1(result int) {
	fmt.Printf("Result = %d\n", result)
}

func ResultPrinter2(result1, result2 int) {
	fmt.Printf("Result1 = %d, Result2 = %d\n", result1, result2)
}

func SumAndMyltyply(first, second int) (int, int) {
	return first + second, first * second
}

func namedSumAndMyltiply(first, second int64) (sum, multiply int) {
	sum = int(first + second)
	multiply = int(first) * int(second)
	return
}
