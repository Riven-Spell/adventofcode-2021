package solutions

import (
	"fmt"
	"github.com/Virepri/adventofcode-2021/util"
	"strconv"
	"strings"
)

type Day13Solution struct {
	Points []util.Vector2D
	Folds  []util.Vector2D
}

func (s *Day13Solution) Prepare(input string) {
	if s.Folds != nil {
		return
	}

	insertingPoints := true
	for _, v := range strings.Split(input, "\n") {
		if v == "" {
			insertingPoints = false
			continue
		}

		if insertingPoints {
			v2 := util.Vector2D{}

			_, err := fmt.Sscanf(v, "%d,%d", &v2.X, &v2.Y)
			util.PanicIfErr(err)

			s.Points = append(s.Points, v2)
		} else {
			match := strings.TrimPrefix(v, "fold along ")
			v2 := util.Vector2D{}

			target := &v2.X // select axis
			if match[0] == 'y' {
				target = &v2.Y
			}

			var err error
			*target, err = strconv.ParseInt(match[2:], 10, 64)
			util.PanicIfErr(err)

			s.Folds = append(s.Folds, v2)
		}
	}
}

func (s *Day13Solution) FoldPage(in map[util.Vector2D]bool, fold util.Vector2D) (out map[util.Vector2D]bool) {
	out = make(map[util.Vector2D]bool)

	for k, v := range in {
		tmp := k

		if (fold.X != 0 && tmp.X == fold.X) || (fold.Y != 0 && tmp.Y == fold.Y) { // omit the field
			continue
		}

		if (fold.X != 0 && tmp.X > fold.X) || (fold.Y != 0 && tmp.Y > fold.Y) { // mirror the point.
			// invert the unaffected field, since it's a full subtraction.
			if fold.X == 0 {
				tmp.X = -tmp.X
			} else if fold.Y == 0 {
				tmp.Y = -tmp.Y
			}

			// fold - (tmp - fold) does the job; tmp - fold removes the extra from fold, fold - mirrors
			tmp = fold.Sub(tmp.Sub(fold))
		}

		out[tmp] = v
	}

	return
}

func (s *Day13Solution) PrintMap(in map[util.Vector2D]bool, empty string) string {
	// find the bounds
	max := util.Vector2D{}
	image := ""

	for k := range in {
		if k.X > max.X {
			max.X = k.X
		}

		if k.Y > max.Y {
			max.Y = k.Y
		}
	}

	idx := util.Vector2D{}
	for idx.Y <= max.Y {
		for idx.X <= max.X {
			image += util.TernaryString(in[idx], "#", empty)

			idx.X++
		}
		image += "\n"

		idx.X = 0
		idx.Y++
	}

	return image[:len(image)-1]
}

func (s *Day13Solution) Part1() string {
	// first, plot the points.
	page := map[util.Vector2D]bool{}
	for _, v := range s.Points {
		page[v] = true
	}

	// fold the page
	folded := s.FoldPage(page, s.Folds[0])

	sum := 0
	for _, v := range folded {
		if v {
			sum++
		}
	}

	return fmt.Sprint(sum)
}

func (s *Day13Solution) Part2() string {
	// first, plot the points.
	page := map[util.Vector2D]bool{}
	for _, v := range s.Points {
		page[v] = true
	}

	// fold the page
	for _, v := range s.Folds {
		page = s.FoldPage(page, v)
	}

	return s.PrintMap(page, " ")
}
