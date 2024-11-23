package main

import (
	"fmt"
	"math"
)

func main() {
	var a float32 = 3
	var b float32 = 5
	c := 20
	d := 6
	var e float64 = 144
	var f float64 = 25.52322

	sum := a - b

	div := a / b

	rem := c % d

	sqrt := math.Sqrt(e) //корень числаБ переменная должна быть float64

	round1 := math.Ceil(f) // округление в большую сторону

	round2 := math.Floor(f) //округление в меньшую сторону

	round3 := math.Round(f) //округление к ближайшему

	fmt.Println(sum)
	fmt.Println(div)
	fmt.Println(rem)
	fmt.Println(sqrt)
	fmt.Println(round1)
	fmt.Println(round2)
	fmt.Println(round3)
}
