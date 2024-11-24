package main

import (
	"fmt"
	"math/rand"
)

type Kilometrs uint
type Days uint
type MillionDollars uint
type KilometersSec uint

const DistanceToMars = 62_100_000
const SecInDay = 86_400

type Ticket struct {
	SpaceLineName string
	Duration      Days
	TripType      string
	Price         MillionDollars
}

var TripeType = [2]string{"RoundTripe", "OneWay"}

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

func TripeTipeSet() string {
	return TripeType[rand.Intn(2)]
}

func (t *Ticket) TicketGeneration() []*Ticket {
	var a []Ticket
	for i := 0; i < 10; i++ {
		t.SpaceLineName = SpaceStation[SortStation()[i]]
		t.Duration = DaysSet()
		t.TripType = TripeTipeSet()
		if t.TripType == TripeType[0] {
		    t.Price = PriceSet(SpeedSet()) * 2}
		if t.TripType == TripeType[1] {
			t.Price = PriceSet(SpeedSet())}
	}
}

func PrintTickets

func main() {
	fmt.Println(DaysSet())
}
