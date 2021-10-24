package main

import (
	"flag"
	"fmt"

	"github.com/kyokomi/emoji/v2"
	"github.com/lunux2008/xulu"

	"BookingBR.com/roomallocation"
	"BookingBR.com/searchhotels"
	"BookingBR.com/shortestpath"
	"BookingBR.com/utils/data"
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

	var bookings []model.Booking

	bookings = roomallocation.BuildBookingOnline(data.Names, data.Values, data.Time, data.Star, data.Destiny)

	taken, timeFull := roomallocation.Glutton(bookings, maxDays, keyString)

	f := func(i, j int) bool {
		return bookings[i].Value > bookings[j].Value
	}

	searchhotels.KnapSack(taken, 3, f)

	v, hotels := searchhotels.BestCombination(taken, maxStars)

	fmt.Print(`
	O tempo total das acomodações foi de `, int64(timeFull/1440), ` dias`)

	fmt.Print(`
	O valor total das acomodações foi de R$ `, v)

	fmt.Println(`
	O Hoteis escolhidos foram os seguintes: `)

	for _, v := range hotels {
		fmt.Println(`		`, v.Name, v.Value)
	}

	answerPath := user.ReadPathAnswer()

	bestPath := []shortestpath.Vertex{}

	if answerPath {

		for _, v := range hotels {
			bestPath = append(bestPath, shortestpath.Vertex(v.Dest))
		}

		g := shortestpath.GraphS{bestPath, make(map[shortestpath.Vertex]map[shortestpath.Vertex]int)}

		for i := 0; i < len(data.Link); i++ {
			g.Edge(data.Link[i].Source, data.Link[i].Local, data.Link[i].Weight)
		}

		dist, next := shortestpath.FloydWarshall(g)

		s := ""
		for u, m := range dist {
			for v, d := range m {
				if u != v {
					source, local := g.Path(shortestpath.Path(u, v, next))
					xulu.Use(local, d)
					for _, v := range hotels {
						if v.Dest == source {
							s = `	O melhor trajeto para começar é ir de ` + v.Name + ` para os demais hoteis !`
						}

					}
				}
			}
		}
		fmt.Printf(s)
	}

}
