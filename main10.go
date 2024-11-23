package main

import "fmt"

func main() {
	var money map[string]int = map[string]int{
		"dollars": 1000,
		"euros":   2000, //мапа срзу сортируется
		"apples":  3,
	}

	//money := map[string]int можно так
	money["dollars"] = 5000 //изменили значение
	delete(money, "apples")
	fmt.Println(money["dollars"])

	a := 1
	c, a := 5, 3
	fmt.Println(a, c)
}
