package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	l, w float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) Area() float64 {
	return r.l * r.w
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r *Rectangle) Perimeter() float64 {
	return (2 * r.l) + (2 * r.w)
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.Area()
	}
	return area
}

func main() {
	var c Shape
	var r Shape
	c = &Circle{x: 0, y: 0, r: 5}
	r = &Rectangle{l: 8, w: 3}

	fmt.Println("c", c)
	fmt.Println("r", r)
	// you can't access the fields b/c
	// the methods were defined for the shape, not the interface:
	//	fmt.Println(c.(Circle).x, c.y, c.r)
	//	fmt.Println(r.l, r.w)

	fmt.Println(c.Area())
	fmt.Println(r.Area())
	fmt.Println(c.Perimeter())
	fmt.Println(r.Perimeter())

	fmt.Println(totalArea(c, r))
}
