package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (r Rectangle) CalcPerimeter() (result float64) {
	result = 2 * (r.Height + r.Weight)
	return
}

func (r Rectangle) CalcArea() (result float64) {
	result = r.Height * r.Weight
	return
}
