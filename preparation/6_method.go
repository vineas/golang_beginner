package main

import "fmt"

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

func measureShape(s shape) {
	fmt.Printf("Shape has an area if %f\n", s.getArea())
	fmt.Printf("Shape has an perimeter if %f\n", s.getPerimeter())
}

func (r rectangle) getPerimeter() float64 {
	return 2*r.width + 2*r.height
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

	measureShape(myRectangle)

}
