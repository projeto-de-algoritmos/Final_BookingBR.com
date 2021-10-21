package roomallocation

import (
	"flag"
	"fmt"
	"sort"
)

type Booking struct {
	name  string
	value float64
	time  float64
}

func (b *Booking) getValue() float64 {
	return b.value
}

func (b *Booking) getTime() float64 {
	return b.time
}

func BuildBookingOnline(names []string, values []float64, time []float64) []Booking {
	booking := []Booking{}
	for index, value := range values {
		booking = append(booking, Booking{names[index], value, time[index]})
	}
	return booking
}

func Glutton(items []Booking, maxTime float64, keyString string) ([]Booking, float64) {
	var bookingCopy = make([]Booking, len(items))
	copy(bookingCopy, items)

	switch keyString {
	case "value":
		sort.Slice(bookingCopy, func(i, j int) bool {
			return bookingCopy[i].getValue() > bookingCopy[j].getValue()
		})

	case "time":
		sort.Slice(bookingCopy, func(i, j int) bool {
			return (1 / bookingCopy[i].getTime()) > (1 / bookingCopy[j].getTime())
		})
	}

	result := []Booking{}
	var totalValue float64
	var totalTime float64
	totalValue, totalTime = 0.0, 0.0
	for i := 0; i < len(bookingCopy); i++ {
		if totalTime+bookingCopy[i].getTime() <= maxTime {
			result = append(result, bookingCopy[i])
			totalTime += bookingCopy[i].getTime()
			totalValue += bookingCopy[i].getValue()
		}
	}

	return result, totalTime
}

func RunGlutton(items []Booking, constraint float64, keyString string) {
	taken, val := Glutton(items, constraint, keyString)
	fmt.Println("Time available for booking =", constraint, "minutes")
	fmt.Println("Total time of all selected rooms =", val, "minutes")
	for idx := range taken {

		fmt.Printf("	%s <%d, %d>\n", taken[idx].name, int32(taken[idx].value), int32(taken[idx].time))
	}
}

func RunGluttons(booking []Booking, maxUnits float64) {
	fmt.Printf("Use glutton by time to allocate %d rooms\n\n", int32(len(booking)))
	RunGlutton(booking, maxUnits, "time")
}

/*
1 dia = 1440 min
1/2 dia = 2160 min
2 dias = 2880 min
3 dias = 4320 min
4 dias = 5760 min
5 dias = 7200 min
6 dias = 8640 min
7 dias = 10080 min
*/

func main() {
	const (
		defaultConstraint = 14060.0 //Limit of booking
	)

	var constraint float64
	flag.Float64Var(&constraint, "constraint", defaultConstraint, "Constraint value about Rooms")
	flag.Parse()

	names := []string{
		"Lolapaluza",
		"NotreDame",
		"Shaluna",
		"LasNoches",
		"Bienvenue",
		"Cielo",
		"Amigos",
		"Donatello",
		"SecondHand Hotel",
		"Spitfire Hotel",
		"Wicked Hotel",
		"Hotel Escolhas",
		"Foundry Hotel",
		"Quantum Hotel",
		"Oráculo Hotel",
		"Global Hotel",
		"Glorial Hotel",
	}

	values := []float64{
		289.0,
		190.0,
		195.0,
		300.0,
		130.0,
		279.0,
		350.0,
		110.0,
		689.0,
		490.0,
		905.0,
		515.0,
		130.0,
		279.0,
		350.0,
		1310.0,
		1410.0,
	}

	time := []float64{
		7200.0,
		2880.0,
		4320.0,
		1440.0,
		5760.0,
		8640.0,
		10080.0,
		2160.0,
		17200.,
		6580.,
		1320.,
		7840.,
		5760.,
		8640.,
		10080.,
		12160.,
		12160.,
	}

	bookings := BuildBookingOnline(names, values, time)

	RunGluttons(bookings, constraint)
}
