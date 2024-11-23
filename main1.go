package main

import "fmt"

func main() {
	age := 34
	name := "Natasha"
	fmt.Println(name + " " + fmt.Sprint(age) + " года")
	fmt.Println("Введите ваше имя:")
	fmt.Scan(&name)
	fmt.Println("Hellow " + name + "!")

	//var a int8 = 2
	//var g int64 = int64(a)

	var a []int = make([]int, 3)
	fmt.Print(a[2])

	var m map[int]string
}
