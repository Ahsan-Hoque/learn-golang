package main

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Widht  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Widht * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (c Triangle) Area() float64 {
	return (c.Base * c.Height) * 0.5
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Widht + rectangle.Height)
}
