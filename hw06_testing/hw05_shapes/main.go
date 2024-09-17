package main

import (
	"errors"
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle with radius %.2f", c.Radius)
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle with width %.2f and height %.2f", r.Width, r.Height)
}

type Triangle struct {
	Base, Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) String() string {
	return fmt.Sprintf("Triangle with base %.2f and height %.2f", t.Base, t.Height)
}

func calculateArea(s any) (float64, error) {
	if shape, ok := s.(Shape); ok {
		return shape.Area(), nil
	}
	return 0, errors.New("invalid shape")
}

func PrintArea(s any) {
	area, err := calculateArea(s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Area of %s is %.2f\n", s, area)
	}
}

func main() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}
	triangle := Triangle{Base: 8, Height: 6}

	PrintArea(circle)
	PrintArea(rectangle)
	PrintArea(triangle)
	PrintArea(nil)
}
