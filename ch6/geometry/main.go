package main

import "math"

type Point struct{ x, y float64 }
type Path []Point

func (p Point) Distance(q Point) float64 {
	//manhattan distance
	return math.Abs(p.x-q.x) + math.Abs(p.y-q.y)
}
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range len(path) {
		if i > 0 {
			sum += path[i].Distance(path[i-1])
		}
	}
	return sum
}
