package main

import (
	"fmt"
	"math"
)

type RoundHole struct {
	radius float64
}

func (rh *RoundHole) GetRadius() float64 {
	return rh.radius
}

func (rh *RoundHole) Fits(peg Peg) bool {
	return rh.GetRadius() >= peg.GetRadius()
}

type Peg interface {
	GetRadius() float64
}

type RoundPeg struct {
	radius float64
}

func (rp *RoundPeg) GetRadius() float64 {
	return rp.radius
}

type SquarePeg struct {
	width float64
}

func (sp *SquarePeg) GetWidth() float64 {
	return sp.width
}

type SquarePegAdapter struct {
	SquarePeg
}

func (spa *SquarePegAdapter) GetRadius() float64 {
	return spa.GetWidth() * math.Sqrt(2) / 2
}

func main() {
	rh := &RoundHole{radius: 5}
	rp := &RoundPeg{radius: 5}

	fmt.Println(rh.Fits(rp))

	sp1 := &SquarePeg{width: 5}
	sp2 := &SquarePeg{width: 10}

	sp1a := &SquarePegAdapter{*sp1}
	sp2a := &SquarePegAdapter{*sp2}

	fmt.Println(rh.Fits(sp1a))
	fmt.Println(rh.Fits(sp2a))
}
