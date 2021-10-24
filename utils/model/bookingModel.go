package model

import (
	"BookingBR.com/shortestpath"
)

type Destiny struct {
	Source shortestpath.Vertex
	Local  shortestpath.Vertex
	Weight int
}

type Booking struct {
	Name  string
	Value float64
	Time  float64
	Stars float64
	Dest  int
}

func (b *Booking) GetValue() float64 {
	return b.Value
}

func (b *Booking) GetTime() float64 {
	return b.Time
}
