package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	IsValidShape() bool
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) IsValidShape() bool {
	return true
}

func (c Circle) String() string {
	return fmt.Sprintf("Круг: радиус %.f", c.Radius)
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) IsValidShape() bool {
	return true
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Прямоугольник: ширина %.f, высота %.f", r.Width, r.Height)
}

type Triangle struct {
	Base, Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) IsValidShape() bool {
	return true
}

func (t Triangle) String() string {
	return fmt.Sprintf("Треугольник: основание %.f, высота %.f", t.Base, t.Height)
}

func calculateArea(s Shape) float64 {
	return s.Area()
}

func PrintArea(s Shape) (string, error) {
	if !s.IsValidShape() {
		return "", fmt.Errorf("the transferred object is not a valid shape")
	}
	area := calculateArea(s)
	return fmt.Sprintf("%v\nПлощадь: %v", s, area), nil
}

func main() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 10, Height: 5}
	triangle := Triangle{Base: 8, Height: 6}

	if areaStr, err := PrintArea(circle); err == nil {
		fmt.Println(areaStr)
	}
	if areaStr, err := PrintArea(rectangle); err == nil {
		fmt.Println(areaStr)
	}
	if areaStr, err := PrintArea(triangle); err == nil {
		fmt.Println(areaStr)
	}
}
