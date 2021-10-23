package main

import (
	"flag"
	"fmt"

	"github.com/kyokomi/emoji/v2"

	"BookingBR.com/roomallocation"
	"BookingBR.com/searchhotels"
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
		5760.,
		8640.,
		10080.,
		12160.,
		12160.,
	}

	Star := []float64{
		7.0,
		2.0,
		4.0,
		1.0,
		5.0,
		8.0,
		10.0,
		2.0,
		1.,
		6.,
		1.,
		7.,
		5.,
		8.,
		10.,
		1.,
		1.,
	}

	var bookings []model.Booking

	bookings = roomallocation.BuildBookingOnline(Names, Values, Time, Star)

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
		fmt.Println(answerPath)
	}

	/*g := shortestpath.GraphS{[]shortestpath.Vertex{1, 2, 3, 4}, make(map[shortestpath.Vertex]map[shortestpath.Vertex]int)}

	g.Edge(1, 2, 2)
	g.Edge(2, 3, 6)
	g.Edge(3, 2, 7)
	g.Edge(4, 3, 1)
	g.Edge(4, 5, 3)
	g.Edge(5, 1, 1)
	g.Edge(5, 2, 4)
	g.Edge(3, 1, 3)
	g.Edge(4, 2, 1)

	dist, next := shortestpath.FloydWarshall(g)
	fmt.Println("pair\tdist\tpath")
	for u, m := range dist {
		for v, d := range m {
			if u != v {
				fmt.Printf("%d -> %d\t%3d\t%s\n", u, v, d, g.Path(shortestpath.Path(u, v, next)))
			}
		}
	}*/
}
