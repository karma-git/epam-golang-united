package golang_united_school_homework

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestBox_AddShape(t *testing.T) {
	circle := &Circle{}
	box := NewBox(1)

	actualErr := box.AddShape(circle)

	if len(box.shapes) != 1 {
		t.Errorf("the expected length is 1, but actual %d", len(box.shapes))
	}
	if actualErr != nil {
		t.Errorf("could not add shape to the box %v", actualErr)
	}
}

func TestBox_AddShape_ErrAddMoreThanMax(t *testing.T) {
	circle := &Circle{}
	triangle := &Triangle{}
	rectangle := &Rectangle{}
	box := NewBox(2)

	_ = box.AddShape(circle)
	_ = box.AddShape(triangle)
	actualErr := box.AddShape(rectangle)

	if actualErr == nil {
		t.Errorf("added more than max elements, but error has not been received")
	}
}

func TestBox_GetByIndex(t *testing.T) {
	circle := &Circle{Radius: 20}
	triangle := &Triangle{Side: 30}
	rectangle := &Rectangle{Height: 10, Weight: 20}
	box := NewBox(3)

	_ = box.AddShape(circle)
	_ = box.AddShape(triangle)
	_ = box.AddShape(rectangle)
	actualShape, actualErr := box.GetByIndex(1)

	if len(box.shapes) != 3 {
		t.Errorf("the expected length is 3, but actual %d", len(box.shapes))
	}
	if actualErr != nil {
		t.Errorf("received unexpected err %v", actualErr)
	}
	if !reflect.DeepEqual(actualShape, triangle) {
		t.Errorf("expected to get %v, but actual is %v", triangle, actualShape)
	}
}

func TestBox_GetByIndex_ErrShapeIsNotFound(t *testing.T) {
	circle := &Circle{Radius: 20}
	triangle := &Triangle{Side: 30}
	rectangle := &Rectangle{Height: 10, Weight: 20}
	box := NewBox(3)

	_ = box.AddShape(circle)
	_ = box.AddShape(triangle)
	_ = box.AddShape(rectangle)
	actualShape, actualErr := box.GetByIndex(3)

	if len(box.shapes) != 3 {
		t.Errorf("the expected length is 3, but actual %d", len(box.shapes))
	}
	if actualErr == nil {
		t.Errorf("expected to get an error, but error has not been received")
	}
	if actualShape != nil {
		t.Errorf("expected to get nil, but actual is %v", actualShape)
	}
}

func TestBox_ExtractByNumber(t *testing.T) {
	circle := &Circle{Radius: 20}
	triangle := &Triangle{Side: 30}
	rectangle := &Rectangle{Height: 10, Weight: 20}
	box := NewBox(3)

	_ = box.AddShape(circle)
	_ = box.AddShape(triangle)
	_ = box.AddShape(rectangle)
	actualShape, actualErr := box.ExtractByIndex(1)

	if len(box.shapes) != 2 {
		t.Errorf("the expected length is 2, but actual %d", len(box.shapes))
	}
	if !reflect.DeepEqual(actualShape, triangle) {
		t.Errorf("expected to get %v, but actual is %v", triangle, actualShape)
	}
	if actualErr != nil {
		t.Errorf("received unexpected err %v", actualErr)
	}
}

func TestBox_ExtractByNumber_ErrOutOfRange(t *testing.T) {
	circle := &Circle{Radius: 20}
	triangle := &Triangle{Side: 30}
	rectangle := &Rectangle{Height: 10, Weight: 20}
	box := NewBox(3)

	_ = box.AddShape(circle)
	_ = box.AddShape(triangle)
	_ = box.AddShape(rectangle)
	actualShape, actualErr := box.ExtractByIndex(3)

	if len(box.shapes) != 3 {
		t.Errorf("the expected length is 3, but actual %d", len(box.shapes))
	}
	if actualShape != nil {
		t.Errorf("expected to get nil, but actual is %v", actualShape)
	}
	if actualErr == nil {
		t.Errorf("expected to get an error")
	}
}

func TestBox_ReplaceByIndex(t *testing.T) {
	circle := &Circle{Radius: 20}
	triangle := &Triangle{Side: 30}
	rectangle := &Rectangle{Height: 10, Weight: 20}
	rectangle2 := &Rectangle{}
	box := NewBox(3)

	_ = box.AddShape(circle)
	_ = box.AddShape(triangle)
	_ = box.AddShape(rectangle)
	actualShape, actualErr := box.ReplaceByIndex(2, rectangle2)

	if len(box.shapes) != 3 {
		t.Errorf("the expected length is 3, but actual %d", len(box.shapes))
	}
	if !reflect.DeepEqual(actualShape, rectangle) {
		t.Errorf("expected to get %v, but actual is %v", triangle, actualShape)
	}
	if actualErr != nil {
		t.Errorf("received unexpected err %v", actualErr)
	}
}

func TestBox_ReplaceByIndex_ErrOutOfRange(t *testing.T) {
	circle := &Circle{Radius: 20}
	triangle := &Triangle{Side: 30}
	rectangle := &Rectangle{Height: 10, Weight: 20}
	rectangle2 := &Rectangle{Height: 10, Weight: 20}
	box := NewBox(3)

	_ = box.AddShape(circle)
	_ = box.AddShape(triangle)
	_ = box.AddShape(rectangle)
	_ = box.AddShape(rectangle2)
	actualShape, actualErr := box.ReplaceByIndex(3, rectangle2)

	if len(box.shapes) != 3 {
		t.Errorf("the expected length is 3, but actual %d", len(box.shapes))
	}
	if actualShape != nil {
		t.Errorf("expected to get nil, but actual is %v", actualShape)
	}
	if actualErr == nil {
		t.Errorf("expected to get an error")
	}
}

func TestBox_SumPerimeter(t *testing.T) {
	circle := &Circle{Radius: 20}
	triangle := &Triangle{Side: 30}
	rectangle := &Rectangle{Height: 10, Weight: 20}
	box := NewBox(3)
	_ = box.AddShape(circle)
	_ = box.AddShape(triangle)
	_ = box.AddShape(rectangle)
	expectedSumPerimeter := (math.Pi * 40) + 90 + 60

	actualSumPerimeter := box.SumPerimeter()

	if expectedSumPerimeter != actualSumPerimeter {
		t.Errorf("expected to get %f, but actual is: %f", expectedSumPerimeter, actualSumPerimeter)
	}
}

func TestBox_SumArea(t *testing.T) {
	circle := &Circle{Radius: 20}
	triangle := &Triangle{Side: 30}
	rectangle := &Rectangle{Height: 10, Weight: 20}
	box := NewBox(3)
	_ = box.AddShape(circle)
	_ = box.AddShape(triangle)
	_ = box.AddShape(rectangle)
	var expectedSumArea = (math.Pi * 400) + 200 + (math.Sqrt(3) / 4 * 900)

	actualSumArea := box.SumArea()

	if expectedSumArea != actualSumArea {
		t.Errorf("expected to get %f, but actual is: %f", expectedSumArea, actualSumArea)
	}
}

func TestBox_RemoveAllCircles(t *testing.T) {
	circle1 := &Circle{Radius: 20}
	circle2 := &Circle{Radius: 30}
	circle3 := &Circle{Radius: 5}
	circle4 := &Circle{Radius: 2}
	triangle := &Triangle{Side: 30}
	rectangle := &Rectangle{Height: 10, Weight: 20}
	box := NewBox(6)
	_ = box.AddShape(circle1)
	_ = box.AddShape(circle2)
	_ = box.AddShape(circle3)
	_ = box.AddShape(circle4)
	_ = box.AddShape(triangle)
	_ = box.AddShape(rectangle)

	actualErr := box.RemoveAllCircles()

	if len(box.shapes) != 2 {
		fmt.Println(box.shapes)
		t.Errorf("expected length is 2, but received %d", len(box.shapes))
	}
	if actualErr != nil {
		t.Errorf("received unexpected err %v", actualErr)
	}
}

func TestBox_RemoveAllCircles_ErrCirclesDoNotExist(t *testing.T) {
	triangle := &Triangle{}
	rectangle := &Rectangle{}
	box := NewBox(2)
	_ = box.AddShape(triangle)
	_ = box.AddShape(rectangle)

	actualErr := box.RemoveAllCircles()

	if len(box.shapes) != 2 {
		t.Errorf("expected length is 2, but received %d", len(box.shapes))
	}
	if actualErr == nil {
		t.Errorf("expected to get an error, but actual is nil %v", actualErr)
	}
}
