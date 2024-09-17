package main

import (
	"fmt"
	"math"
	"testing"
)

func TestCircleArea(t *testing.T) {
	tests := []struct {
		radius       float64
		expectedArea float64
	}{
		{5, math.Pi * 5 * 5},
		{10, math.Pi * 10 * 10},
		{0, 0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Circle with radius %.2f", tt.radius), func(t *testing.T) {
			circle := Circle{Radius: tt.radius}
			actualArea := circle.Area()
			if actualArea != tt.expectedArea {
				t.Errorf("Expected area %.2f, got %.2f", tt.expectedArea, actualArea)
			}
		})
	}
}

func TestRectangleArea(t *testing.T) {
	tests := []struct {
		width, height float64
		expectedArea  float64
	}{
		{4, 6, 24},
		{5, 5, 25},
		{0, 10, 0},
		{10, 0, 0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Rectangle %.2fx%.2f", tt.width, tt.height), func(t *testing.T) {
			rectangle := Rectangle{Width: tt.width, Height: tt.height}
			actualArea := rectangle.Area()
			if actualArea != tt.expectedArea {
				t.Errorf("Expected area %.2f, got %.2f", tt.expectedArea, actualArea)
			}
		})
	}
}

func TestTriangleArea(t *testing.T) {
	tests := []struct {
		base, height float64
		expectedArea float64
	}{
		{8, 6, 24},
		{10, 5, 25},
		{0, 10, 0},
		{10, 0, 0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Triangle base %.2f, height %.2f", tt.base, tt.height), func(t *testing.T) {
			triangle := Triangle{Base: tt.base, Height: tt.height}
			actualArea := triangle.Area()
			if actualArea != tt.expectedArea {
				t.Errorf("Expected area %.2f, got %.2f", tt.expectedArea, actualArea)
			}
		})
	}
}

func TestCalculateArea(t *testing.T) {
	tests := []struct {
		name      string
		shape     any
		expectErr bool
		expected  float64
	}{
		{"Circle", Circle{Radius: 5}, false, math.Pi * 5 * 5},
		{"Rectangle", Rectangle{Width: 4, Height: 6}, false, 24},
		{"Triangle", Triangle{Base: 8, Height: 6}, false, 24},
		{"Invalid Shape", "not a shape", true, 0},
		{"Nil", nil, true, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := calculateArea(tt.shape)
			if (err != nil) != tt.expectErr || (err == nil && actual != tt.expected) {
				t.Errorf("Expected area: %.2f, got: %.2f, error: %v", tt.expected, actual, err)
			}
		})
	}
}
