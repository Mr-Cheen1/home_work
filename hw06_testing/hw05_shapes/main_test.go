package main

import (
	"fmt"
	"testing"
)

type DummyShape struct{}

func (ds DummyShape) Area() float64 {
	return 0
}

func (ds DummyShape) IsValidShape() bool {
	return false
}

func TestCalculateArea(t *testing.T) {
	tests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{
			name:     "Circle area",
			shape:    Circle{Radius: 5},
			expected: 78.53981633974483,
		},
		{
			name:     "Rectangle area",
			shape:    Rectangle{Width: 10, Height: 5},
			expected: 50,
		},
		{
			name:     "Triangle area",
			shape:    Triangle{Base: 8, Height: 6},
			expected: 24,
		},
		{
			name:     "Invalid shape area",
			shape:    DummyShape{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateArea(tt.shape); got != tt.expected {
				t.Errorf("calculateArea() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestIsValidShape(t *testing.T) {
	tests := []struct {
		shape Shape
		valid bool
	}{
		{Circle{Radius: 5}, true},
		{Rectangle{Width: 10, Height: 5}, true},
		{Triangle{Base: 8, Height: 6}, true},
		{DummyShape{}, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%T valid", tt.shape), func(t *testing.T) {
			if got := tt.shape.IsValidShape(); got != tt.valid {
				t.Errorf("%T.IsValidShape() = %v, want %v", tt.shape, got, tt.valid)
			}
		})
	}
}

func TestPrintArea(t *testing.T) {
	tests := []struct {
		name    string
		shape   Shape
		want    string
		wantErr bool
	}{
		{"Circle", Circle{Radius: 5}, "Круг: радиус 5\nПлощадь: 78.53981633974483", false},
		{"Rectangle", Rectangle{Width: 10, Height: 5}, "Прямоугольник: ширина 10, высота 5\nПлощадь: 50", false},
		{"Triangle", Triangle{Base: 8, Height: 6}, "Треугольник: основание 8, высота 6\nПлощадь: 24", false},
		{"NotAShape", DummyShape{}, "the transferred object is not a valid shape", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrintArea(tt.shape)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrintArea() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("PrintArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
