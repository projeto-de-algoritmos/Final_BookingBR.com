package model

type Booking struct {
	Name  string
	Value float64
	Time  float64
	Stars float64
}

func (b *Booking) GetValue() float64 {
	return b.Value
}

func (b *Booking) GetTime() float64 {
	return b.Time
}
