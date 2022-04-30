package main

import (
	"fmt"
	"math"
)

/*
	Разработать программу нахождения расстояния между двумя точками,
	которые представлены в виде структуры Point
	с инкапсулированными параметрами x, y и конструктором.
*/

type Numbers interface {
	byte | int | int64 | int32 | int8 | float32 | float64
}

type Point[T Numbers] struct {
	x, y T
}

func NewPoint[T Numbers](x, y T) *Point[T] {
	return &Point[T]{x, y}
}

func (p *Point[T]) Sub(p2 *Point[T]) *Point[T] {
	return NewPoint[T](p.x-p2.x, p.y-p2.y)
}

func (p *Point[T]) Distance(p2 *Point[T]) float64 {
	var n T
	t := p.Sub(p2)
	n = (t.x * t.x) + (t.y * t.y)
	return math.Sqrt(float64(n))
}

func main() {
	pFloat1 := NewPoint(10.0, 10.0)
	pFloat2 := NewPoint(15.0, 15.0)
	fmt.Println(pFloat1.Distance(pFloat2))

	pInt1 := NewPoint(15, 2)
	pInt2 := NewPoint(8, 13)
	fmt.Println(pInt1.Distance(pInt2))

	pByte1 := NewPoint[byte](50, 14)
	pByte2 := NewPoint[byte](28, 12)
	fmt.Println(pByte1.Distance(pByte2))

}
