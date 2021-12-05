package solutions

import (
	"fmt"
	"github.com/Virepri/adventofcode-2021/util"
	"strings"
)

type Day5Solution struct {
	Lines [][2]util.Vector2D
}

func (s *Day5Solution) Prepare(input string) {
	if s.Lines != nil {
		return
	}

	for _, v := range strings.Split(input, "\n") {
		out := [2]util.Vector2D{}

		_, err := fmt.Sscanf(v, "%d,%d -> %d,%d", &out[0].X, &out[0].Y, &out[1].X, &out[1].Y)
		util.PanicIfErr(err)

		s.Lines = append(s.Lines, out)
	}
}

func (s *Day5Solution) Solve(needsDiags bool) string {
	out := map[util.Vector2D]uint{}

	for _, v := range s.Lines {
		// Only considering non-sloped lines.
		if !needsDiags && util.HasSlope(v[0], v[1]) {
			continue
		}

		util.ProcessEachPoint(v[0], v[1], func(p util.Vector2D) {
			out[p]++
		})
	}

	sum := 0
	for _, v := range out {
		if v >= 2 {
			sum++
		}
	}

	return fmt.Sprint(sum)
}

func (s *Day5Solution) Part1() string {
	return s.Solve(false)
}

func (s *Day5Solution) Part2() string {
	return s.Solve(true)
}
