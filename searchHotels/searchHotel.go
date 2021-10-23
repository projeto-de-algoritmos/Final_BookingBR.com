package searchhotels

import (
	"math"
	"sort"

	"BookingBR.com/utils/model"
)

func KnapSack(bookings []model.Booking, maxStars float64, metric func(i, j int) bool) (r []model.Booking, r2 []model.Booking) {
	sort.Slice(bookings, metric)

	s := 0.

	for _, i := range bookings {
		if s+i.Stars <= maxStars {
			r = append(r, i)
			s += i.Stars
		}
	}
	//fmt.Println(r)
	return
}

func PossibleCombinations(bookings []model.Booking, ch chan []model.Booking) {
	defer close(ch)

	p := int(math.Pow(2., float64(len(bookings))))

	for i := 0; i < p; i++ {
		set := []model.Booking{}
		for j := 0; j < len(bookings); j++ {
			if (i>>uint(j))&1 == 1 {
				set = append(set, bookings[j])
			}
		}
		ch <- set
	}
}

func getSackStars(set []model.Booking) (r float64) {
	for _, i := range set {
		r += i.Stars
	}
	return
}

func getSackValue(set []model.Booking) (r float64) {
	for _, i := range set {
		r += i.Value
	}
	return
}

func BestCombination(bookings []model.Booking, maxStars float64) (float64, []model.Booking) {
	bestVal := 0.
	bestSack := []model.Booking{}

	ch := make(chan []model.Booking)
	go PossibleCombinations(bookings, ch)

	for sack := range ch {
		if getSackStars(sack) <= maxStars {
			v := getSackValue(sack)
			if v > bestVal {
				bestVal = v
				bestSack = sack
			}
		}
	}
	return bestVal, bestSack
}
