package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (c Circle) CalcPerimeter() (result float64) {
	result = 2 * math.Pi * c.Radius
	return
}

func (c Circle) CalcArea() (result float64) {
	result = math.Pi * math.Pow(c.Radius, 2)
	return
}
