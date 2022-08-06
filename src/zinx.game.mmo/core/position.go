package core

import "math/rand"

const (
	HorizontalBase   = 150
	HorizontalOffset = 10
	VerticalBase     = 180
	VerticalOffset   = 20
)

type Position struct {
	X float32
	Y float32
	Z float32
	V float32
}

func NewPosition() *Position {
	return &Position{
		X: float32(HorizontalBase + rand.Intn(HorizontalOffset)),
		Y: 0,
		Z: float32(VerticalBase + rand.Intn(VerticalOffset)),
		V: 0,
	}
}
