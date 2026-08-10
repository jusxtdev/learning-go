package main

import "fmt"

type Shape interface {
	area() float64
	circumference() float64
}

type Rect struct {
	length  int
	breadth int
}

func (r Rect) area() float64 {
	return float64(r.length * r.breadth)
}
func (r Rect) circumference() float64 {
	return float64(r.length + r.breadth)
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * (c.radius * c.radius)
}
func (c Circle) circumference() float64 {
	return 2 * 3.14 * c.radius
}

func print_info(s Shape) {
	fmt.Printf("Area : %f | Circumf : %f \n", s.area(), s.circumference())
}

func main() {
	r := Rect{length: 12, breadth: 23}
	print_info(r)

	c := Circle{radius: 32.12}
	print_info(c)
}
