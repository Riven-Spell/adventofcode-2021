package util

import (
	"math"
)

var AdjacentVector2D = []Vector2D{
	{-1, 0}, // X -1
	{-1, 1},
	{-1, -1},
	{1, 0}, // X 1
	{1, 1},
	{1, -1},
	{0, -1}, // Y -1
	{0, 1},  // Y 1
}

var ImmediateAdjacentVector2D = []Vector2D{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

type Vector2D struct {
	X, Y int64
}

func (p Vector2D) Add(p2 Vector2D) Vector2D {
	return Vector2D{X: p2.X + p.X, Y: p2.Y + p.Y}
}

func (p Vector2D) Sub(p2 Vector2D) Vector2D {
	return Vector2D{X: p.X - p2.X, Y: p.Y - p2.Y}
}

func (p Vector2D) Mul(c int64) Vector2D {
	return Vector2D{X: p.X * c, Y: p.Y * c}
}

func (p Vector2D) Rot(deg float64) Vector2D {
	rads := deg * deg2Rad

	return Vector2D{
		X: int64(math.Round(float64(p.X)*math.Cos(rads) - float64(p.Y)*math.Sin(rads))),
		Y: int64(math.Round(float64(p.Y)*math.Cos(rads) + float64(p.X)*math.Sin(rads))),
	}
}

func Manhattan(p1, p2 Vector2D) int64 {
	return IntAbs(p2.X-p1.X) + IntAbs(p2.Y-p1.Y)
}

func ProcessEachPoint(p1, p2 Vector2D, procFunc func(p Vector2D)) {
	slope := Slope(p1, p2)
	// rise over run
	move := Vector2D{slope.Denominator, slope.Numerator}
	pLoc := p1

	// todo: is there a p1/p2 combo that will cause a endless loop?
	for pLoc != p2 {
		procFunc(pLoc)
		pLoc = pLoc.Add(move)
	}

	procFunc(p2)
}

func HasSlope(p1, p2 Vector2D) bool {
	// if X or Y is equal, it is horizontal or diagonal.
	return p1.X != p2.X && p1.Y != p2.Y
}

func Slope(p1, p2 Vector2D) Fraction {
	if !HasSlope(p1, p2) {
		if p1.X == p2.X && p2.Y == p1.Y {
			return Fraction{0, 0}
		} else if p1.X == p2.X {
			// It'll be processed as rise over run.
			// Thus, a 0-denominated fraction is sensible.
			sign := (p2.Y - p1.Y) / IntAbs(p2.Y-p1.Y)
			return Fraction{1 * sign, 0}
		} else {
			sign := (p2.X - p1.X) / IntAbs(p2.X-p1.X)
			return Fraction{0, 1 * sign}
		}
	}

	f := Fraction{Numerator: p2.Y - p1.Y, Denominator: p2.X - p1.X}.Simplify()

	if f.Numerator == 0 {
		f.Denominator = TernaryInt64(f.Denominator > 0, math.MaxInt64, math.MinInt64)
	}

	if f.Denominator == 0 {
		f.Numerator = TernaryInt64(f.Numerator > 0, math.MaxInt64, math.MinInt64)
	}

	return f
}

type DistSorter struct {
	Center Vector2D
	List   []Vector2D
}

func (s *DistSorter) Len() int {
	return len(s.List)
}

func (s *DistSorter) Less(i, j int) bool {
	return Manhattan(s.Center, s.List[i]) < Manhattan(s.Center, s.List[j])
}

func (s *DistSorter) Swap(i, j int) {
	s.List[i], s.List[j] = s.List[j], s.List[i]
}

func IntAbs(x int64) int64 {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
