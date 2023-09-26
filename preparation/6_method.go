package main

import (
	"fmt"
	"math"
)

type CentiMeters float64

func (c CentiMeters) IsTooLong() {
	if c > 100 {
		fmt.Printf("Too long~ \n")
	} else {
		fmt.Printf("Just Right~ \n")
	}
}

type shape interface {
	getArea() float64
	getPerimeter() float64
}

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) getArea() float64 {
	return r.width * r.height
}

func (r rectangle) getPerimeter() float64 {
	return 2*r.width + 2*r.height
}

func measureShape(s shape) {
	fmt.Printf("Shape has an area if %f\n", s.getArea())
	fmt.Printf("Shape has an perimeter if %f\n", s.getPerimeter())
}

type circle struct {
	radius float64
}

func (c circle) getArea() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) getPerimeter() float64 {
	return 2 * math.Pi * c.radius
}

type square struct {
	width float64
}

func (s square) getArea() float64 {
	return s.width * s.width
}

func (s square) getPerimeter() float64 {
	return 4 * s.width
}

func main() {
	myVar := CentiMeters(101)
	fmt.Printf("Type: %T, Value %v\n", myVar, myVar)
	// myVar.IsTooLong()

	myRectangle := rectangle{
		width:  30,
		height: 23,
	}
	fmt.Printf("Type: %T, Value %+v\n", myRectangle, myRectangle)

	myCircle := circle{
		radius: 5,
	}
	mySquare := square{
		width: 13,
	}

	measureShape(myRectangle)
	measureShape(myCircle)
	measureShape(mySquare)

}
