package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (t Triangle) CalcPerimeter() (result float64) {
	result = 3 * t.Side
	return
}

func (t Triangle) CalcArea() (result float64) {
	result = math.Sqrt(3) / 4 * (t.Side * t.Side)
	return
}
