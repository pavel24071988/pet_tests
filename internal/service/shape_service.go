package service

import (
	"fmt"
	"math"
)

type IShapeInterface interface {
	Area() float32
}

type Square struct {
	SideLenght float32
}

func (s Square) Area() float32 {
	return s.SideLenght * s.SideLenght
}

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return float32(math.Pi) * c.Radius
}

func PrintShapeArea(s IShapeInterface) {
	fmt.Printf("area is equal %.2f centimeters\n", s.Area())
}
