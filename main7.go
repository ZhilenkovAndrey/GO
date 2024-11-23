package main

import (
	"fmt"
	// "math"
	// "unsafe"
)

func main() {
	//var names [3]string
	// names := [3]string{"jordan", "John", "Kate"}
	// names[0] = "jordan"
	// names[2] = "Kate"
	// names[1] = "Emma"

	// fmt.Println(names)
	// fmt.Println(len(names))

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// marks := [5]float64{11, 9, 12, 8, 10}

	// var sum float64 = 0

	// for i := 0; i < len(marks); i++ {
	// 	sum += marks[i]
	// }
	// fmt.Println(sum)

	// var result float64 = sum / float64(len(marks))
	// fmt.Println(math.Round(result))

	// array := [...]int32{1, 2, 3}
	// const elemSize = unsafe.Sizeof(int32(0)) //4 bytes
	// pointer := unsafe.Pointer(&array)

	// first := *(*int32)(unsafe.Add(pointer, 0*elemSize))  //data[0]
	// second := *(*int32)(unsafe.Add(pointer, 1*elemSize)) //data[1]
	// third := *(*int32)(unsafe.Add(pointer, 2*elemSize))  //data[2]

	// dangerous1 := *(*int32)(unsafe.Add(pointer, -1))
	// dangerous2 := *(*int32)(unsafe.Add(pointer, 3))

	// fmt.Println("correct: ", first, second, third)
	// fmt.Println("dangerouse: ", dangerous1, dangerous2)

	values := [...]int{100, 200, 300}  // for idx,value :=range values
	for i := 0; i < len(values); i++ { // value += 50
		values[i] += 50 // изменит переменную, а не
	} // элемент массива
	fmt.Println(values)
}
