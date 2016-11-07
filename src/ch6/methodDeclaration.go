package main

import (
	"math"
	"fmt"
)

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))

	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	t := Point{3, 4}
	pptr := &t
	pptr.ScaleBy(2)
	fmt.Println(t)

	l := Point{5, 6}
	(&l).ScaleBy(2)
	fmt.Println(l)

	fmt.Println(Point{1, 2}.Distance(q))
}
