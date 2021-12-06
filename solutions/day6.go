package solutions

import (
	"fmt"
	"github.com/Virepri/adventofcode-2021/util"
	"strconv"
	"strings"
)

type Day6Solution struct {
	fish []uint
}

func (s *Day6Solution) Prepare(input string) {
	if s.fish != nil {
		return
	}

	for _, v := range strings.Split(input, ",") {
		f, err := strconv.ParseUint(v, 10, 64)
		util.PanicIfErr(err)
		s.fish = append(s.fish, uint(f))
	}
}

func (s *Day6Solution) Solve(totalDays uint) string {
	fish := [9]uint{}

	// insert fish
	for _, v := range s.fish {
		fish[v]++
	}

	// start working
	d := uint(0)
	for d <= totalDays {
		// first, find how many days need to pass for the next birth.
		pass := uint(0)
		for k, v := range fish {
			if v > 0 {
				pass = uint(k)
				break
			}
		}

		// do we need to rotate the array?
		if !(d+pass+1 > totalDays) {
			// next, rotate the array to handle it.
			newBorns := fish[pass]
			for k, v := range fish[pass+1:] {
				// there's nothing prior to fish[pass], so this works.
				fish[k] = v
				fish[uint(k)+pass+1] = 0
			}

			fish[6] += newBorns // reset the cycle for birthing fish
			fish[8] += newBorns // new cycle for newborns
		}

		d += pass + 1
	}

	sum := uint(0)
	for _, v := range fish {
		sum += v
	}

	return fmt.Sprint(sum)
}

func (s *Day6Solution) Part1() string {
	return s.Solve(80)
}

func (s *Day6Solution) Part2() string {
	return s.Solve(256)
}
