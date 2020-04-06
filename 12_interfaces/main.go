package main

import (
	"fmt"
	"math"
)

// Shape : interface for shape
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectange struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectange) area() float64 {
	return r.height * r.width
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle1 := Circle{10, 5, 7}
	rect1 := Rectange{10, 8}

	fmt.Println(circle1.area())
	fmt.Println(rect1.area())

	// or

	fmt.Println(getArea(circle1))
	fmt.Println(getArea(rect1))

}
