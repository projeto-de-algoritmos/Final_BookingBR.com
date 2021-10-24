package main

import (
	"flag"
	"fmt"

	"github.com/kyokomi/emoji/v2"

	"BookingBR.com/roomallocation"
	"BookingBR.com/searchhotels"
	"BookingBR.com/shortestpath"
	"BookingBR.com/utils/model"
	"BookingBR.com/utils/user"
)

func main() {

	const bookingDescription = `
							:hotel: BookingBR.com

	BookingBR.com se trata de um mecanismo onde os clientes podem reservar acomodações para férias ou viagens.✈️  🏖️
		
	`
	emoji.Println(bookingDescription)
	//fmt.Println(bookingDescription)

	const (
		defaultConstraint = 2880.0 //Limit of booking
		keyString         = "Time"
	)

	maxStars, days := user.ReadUser()
	maxDays := user.ConvertDaysToMinute(days)
	flag.Float64Var(&maxDays, "maxDays", maxDays, "Constraint value about Rooms")
	flag.Parse()

	Names := []string{
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

	Values := []float64{
		289.0,
		190.0,
		195.0,
		1300.0,
		130.0,
		279.0,
		350.0,
		110.0,
		1689.0,
		490.0,
		1905.0,
		515.0,
		1130.0,
		279.0,
		350.0,
		1310.0,
		1410.0,
	}

	Time := []float64{
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
		6760.,
		8640.,
		10080.,
		12160.,
		15160.,
	}

	Star := []float64{
		3.0,
		2.0,
		4.0,
		10.0,
		2.0,
		3.0,
		1.0,
		9.0,
		6.,
		8.,
		8.,
		2.,
		3.,
		8.,
		4.,
		8.,
		9.,
	}

	Destiny := []int{
		1, 3, 5, 4, 6, 7, 8, 10, 11, 2, 9, 12, 17, 15, 14, 16, 13,
	}

	Link := []model.Destiny{
		{1, 2, 1},
		{2, 3, 2},
		{3, 4, 3},
		{4, 16, 5},
		{16, 1, 6},
		{1, 4, 4},
		{1, 12, 7},
		{12, 13, 8},
		{13, 15, 4},
		{15, 5, 10},
		{5, 15, 11},
		{5, 4, 1},
		{15, 6, 12},
		{6, 7, 3},
		{13, 14, 2},
		{14, 7, 3},
		{14, 9, 1},
		{7, 8, 4},
		{8, 9, 5},
		{9, 10, 9},
		{10, 11, 7},
		{11, 1, 10},
		{4, 5, 2},
		{5, 6, 4},
	}

	var bookings []model.Booking

	bookings = roomallocation.BuildBookingOnline(Names, Values, Time, Star, Destiny)

	taken, timeFull := roomallocation.Glutton(bookings, maxDays, keyString)

	f := func(i, j int) bool {
		return bookings[i].Value > bookings[j].Value
	}

	searchhotels.KnapSack(taken, 3, f)

	v, s := searchhotels.BestCombination(taken, maxStars)

	fmt.Print(`
	O tempo total das acomodações foi de `, int64(timeFull/1440), ` dias`)

	fmt.Print(`
	O valor total das acomodações foi de R$ `, v)

	fmt.Println(`
	O Hoteis escolhidos foram os seguintes: `)

	for _, v := range s {
		fmt.Println(`		`, v.Name, v.Value)
	}

	answerPath := user.ReadPathAnswer()

	if answerPath {
		g := shortestpath.GraphS{[]shortestpath.Vertex{1, 2, 3, 4}, make(map[shortestpath.Vertex]map[shortestpath.Vertex]int)}

		for i := 0; i < len(Link); i++ {
			g.Edge(Link[i].Source, Link[i].Local, Link[i].Weight)
		}

		dist, next := shortestpath.FloydWarshall(g)
		fmt.Println("pair\tdist\tpath")
		for u, m := range dist {
			for v, d := range m {
				if u != v {
					fmt.Printf("%d -> %d\t%3d\t%s\n", u, v, d, g.Path(shortestpath.Path(u, v, next)))
				}
			}
		}
	}

}
