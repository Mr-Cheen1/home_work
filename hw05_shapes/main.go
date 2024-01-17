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
	return fmt.Sprintf("Круг: радиус %.f", c.Radius)
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
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

func (t Triangle) String() string {
	return fmt.Sprintf("Треугольник: основание %.f, высота %.f", t.Base, t.Height)
}

func calculateArea(s any) (float64, error) {
	if shape, ok := s.(Shape); ok {
		return shape.Area(), nil
	}
	return 0, errors.New("переданный объект не является фигурой")
}

func PrintArea(s any) {
	area, err := calculateArea(s)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("%v\nПлощадь: %v\n", s, area)
	}
}

func main() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 10, Height: 5}
	triangle := Triangle{Base: 8, Height: 6}

	PrintArea(circle)
	PrintArea(rectangle)
	PrintArea(triangle)
	PrintArea(nil)
}
