package main

import "fmt"

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	ex := s.start.x + int(s.a)
	ey := s.start.y + int(s.a)
	return Point{x: ex, y: ey}
}

func (s Square) Area() uint {
	return s.a * s.a
}

func (s Square) Perimeter() uint {
	return s.a * 4
}

func main() {
	p := Point{x: 5, y: 2}
	s := Square{start: p, a: 3}
	area, perimeter := s.Area(), s.Perimeter()
	fmt.Println("Square: ", s)
	fmt.Printf("Area=<%v>,Perimetr=<%v>", area, perimeter)
}
