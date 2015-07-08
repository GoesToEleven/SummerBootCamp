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

func (c *Circle) circleArea() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) rectangleArea() float64 {
	return r.l * r.w
}

func main() {
	c := Circle{x: 0, y: 0, r: 5}
	r := Rectangle{l: 10, w: 4}

	fmt.Println("c", c)
	fmt.Println("r", r)
	fmt.Println(c.x, c.y, c.r)
	fmt.Println(r.l, r.w)

	fmt.Println(c.circleArea())
	fmt.Println(r.rectangleArea())
}
