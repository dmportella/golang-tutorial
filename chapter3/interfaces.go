package chapter3

import (
	"fmt"
	"math"
	"reflect"
)

// Interface for a geometry shape
type geometry interface {
	area() float64
	perim() float64
}

// The representation of a rectangle
type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// The representation of a circle
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// method using a interface
func measure(g geometry) {
	s := reflect.TypeOf(g)

	fmt.Println(s, g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func Interfaces() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
