package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (rect Rectangle) Area() float64 {
	return rect.Height * rect.Width
}

type Circle struct {
	Radius float64
}

func (circle Circle) Area() float64 {
	return circle.Radius * circle.Radius * math.Pi
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * (t.Height / 2)
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}
