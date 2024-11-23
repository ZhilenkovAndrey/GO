package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{3, 1, 2, 5, 7, 4}
	slice = append(slice, 0) // добаляет элемент в конец массива
	slice[0] = 11            //меняет значение элемента
	sort.Ints(slice)         //сортирует срез интов
	//sort.Str()               сортирует срез строк

	for idx, elem := range slice { //пербирает срез
		fmt.Printf("%d, %d\n", idx, elem)
	}

	fmt.Println(slice)
}
