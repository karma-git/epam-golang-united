package golang_united_school_homework

import (
	"errors"
)

const errSlice = "not enough space in box to add new shape"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) == b.shapesCapacity {
		return errors.New(errSlice)
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if len(b.shapes) > i {
		return b.shapes[i], nil
	}
	return nil, errors.New(errSlice)
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if len(b.shapes) > i {
		s := b.shapes[i]
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		return s, nil
	}
	return nil, errors.New(errSlice)
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if len(b.shapes) > i {
		s := b.shapes[i]    // extract old shape by value
		b.shapes[i] = shape // replace old shape by new
		return s, nil
	}
	return nil, errors.New("index is higher then slice length")
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	for _, s := range b.shapes {
		sum += s.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64
	for _, s := range b.shapes {
		sum += s.CalcArea()
	}
	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var shapes []Shape

	for _, shape := range b.shapes {
		switch shape.(type) {
		case *Circle:
			continue
		default:
			shapes = append(shapes, shape)
		}
	}

	if len(b.shapes) == len(shapes) {
		return errors.New("not found")
	}

	b.shapes = shapes

	return nil
}
