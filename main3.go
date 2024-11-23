package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Printf(" Hello %d\n", i)
		if i%2 == 0 {
			//fmt.Println(i)
			continue
		}
	}

	for i := 1; i >= -10; i-- {
		if i == -5 {
			break
		}
	}

	i := 0
	for i < 5 { // аналог while
		fmt.Println(i)
		i++
	}

	nums := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	for index, element := range nums {
		fmt.Printf("Index: %d Element: %d\n", index, element)
	}

	for _, element := range nums { // "_" - сообщает среде о отсутствии не
		fmt.Printf("Element: %d\n", index, element) //обходимости инициализировать
	} //значение.

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for _, row := range matrix {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}
}
