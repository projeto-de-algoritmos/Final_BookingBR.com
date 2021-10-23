package roomallocation

import (
	"fmt"
	"sort"

	"BookingBR.com/utils/model"
)

func BuildBookingOnline(Names []string, Values []float64, Time []float64, Star []float64) []model.Booking {
	booking := []model.Booking{}
	for index, Value := range Values {
		booking = append(booking, model.Booking{Names[index], Value, Time[index], Star[index]})
	}
	return booking
}

func Glutton(items []model.Booking, maxTime float64, keyString string) ([]model.Booking, float64) {
	var bookingCopy = make([]model.Booking, len(items))
	copy(bookingCopy, items)

	switch keyString {
	case "Value":
		sort.Slice(bookingCopy, func(i, j int) bool {
			return bookingCopy[i].GetValue() > bookingCopy[j].GetValue()
		})

	case "Time":
		sort.Slice(bookingCopy, func(i, j int) bool {
			return (1 / bookingCopy[i].GetTime()) > (1 / bookingCopy[j].GetTime())
		})
	}

	result := []model.Booking{}
	var totalValue float64
	var totalTime float64
	totalValue, totalTime = 0.0, 0.0
	for i := 0; i < len(bookingCopy); i++ {
		if totalTime+bookingCopy[i].GetTime() <= maxTime {
			result = append(result, bookingCopy[i])
			totalTime += bookingCopy[i].GetTime()
			totalValue += bookingCopy[i].GetValue()
		}
	}

	return result, totalTime
}

func RunGlutton(items []model.Booking, constraint float64, keyString string) {
	taken, val := Glutton(items, constraint, keyString)
	fmt.Println("Time available for booking =", constraint, "minutes")
	fmt.Println("Total Time of all selected rooms =", val, "minutes")
	for idx := range taken {

		fmt.Printf("	%s <%d, %d>\n", taken[idx].Name, int32(taken[idx].Value), int32(taken[idx].Time))
	}
}

func RunGluttons(booking []model.Booking, maxUnits float64) {
	fmt.Printf("Use glutton by Time to allocate %d rooms\n\n", int32(len(booking)))
	RunGlutton(booking, maxUnits, "Time")
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
