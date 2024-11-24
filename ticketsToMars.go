package main

import (
	"fmt"
	"math/rand"
)

type Kilometrs uint
type Days uint
type TripeType bool
type MillionDollars uint
type KilometersSec uint

const DistanceToMars = 62_100_000
const SecInDay = 86_400

type Ticket struct {
	SpaceLineName string
	Durations     Days
	TripType      TripeType
	Speed         KilometersSec
	Price         MillionDollars
}

var SpaceStation = [5]string{
	"SpaceAdventures",
	"SpaceX",
	"VirginGalaxy",
	"Baykhonure",
	"DeadSpace"}

func SortStation() []int {
	var a []int = rand.Perm(5)
	var b []int = rand.Perm(5)
	for _, c := range b {
		a = append(a, c)
	}
	return a
}

func SpeedSet() KilometersSec {
	return KilometersSec(rand.Intn(14) + 16)
}

func PriceSet(a KilometersSec) MillionDollars {
	return MillionDollars(int(a) + 22)
}

func DaysSet() Days {
	return Days(DistanceToMars / uint(SpeedSet()*24*60*60))
}

func TripTipeSet() int

func main() {
	fmt.Println(DaysSet())
}
