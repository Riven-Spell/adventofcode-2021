package solutions

import (
	"fmt"
	"github.com/Virepri/adventofcode-2021/util"
	"sort"
	"strconv"
	"strings"
)

type Day7Solution struct {
	crabs []int64
}

func (s *Day7Solution) Len() int {
	return len(s.crabs)
}

func (s *Day7Solution) Less(i, j int) bool {
	return s.crabs[i] < s.crabs[j]
}

func (s *Day7Solution) Swap(i, j int) {
	s.crabs[i], s.crabs[j] = s.crabs[j], s.crabs[i]
}

func (s *Day7Solution) Prepare(input string) {
	if s.crabs != nil {
		return
	}

	for _, v := range strings.Split(input, ",") {
		f, err := strconv.ParseInt(v, 10, 64)
		util.PanicIfErr(err)
		s.crabs = append(s.crabs, f)
	}

	// sort the crabs
	sort.Sort(s)
}

func (s *Day7Solution) Solve(calculateFuel func(moveSize int64) int64) int64 {
	minPos := s.crabs[0]
	maxPos := s.crabs[len(s.crabs)-1]

	getFullFuel := func(pos int64) int64 {
		sum := int64(0)
		for _, crab := range s.crabs {
			sum += calculateFuel(util.IntAbs(crab - pos))
		}

		return sum
	}

	// Effectively binary-search our way through this.
	startPoint := minPos + ((maxPos - minPos) / 2)
	min := getFullFuel(startPoint)

	// Which direction towards the minimum?
	inc := util.TernaryInt64(getFullFuel(startPoint-1) < getFullFuel(startPoint+1), -1, 1) * ((maxPos - minPos) / 2)

	// Head towards the minimum.
	for {
		startPoint += inc
		startPoint = util.ClampInt64(startPoint, minPos, maxPos)

		if f := getFullFuel(startPoint); f < min {
			min = f
		} else {
			if abs := util.IntAbs(inc); abs != 1 { // We overshot
				if abs > 1 {
					inc /= -2
				}
				min = f
				continue
			}

			break
		}
	}

	return min
}

func (s *Day7Solution) Part1() string {
	return fmt.Sprint(s.Solve(func(moveSize int64) int64 {
		return moveSize
	}))
}

func (s *Day7Solution) Part2() string {
	target := s.crabs[len(s.crabs)-1] + 1
	cache := make([]int64, target)

	// precompute the maximum summation
	for i := int64(1); i < target; i++ {
		cache[i] = cache[i-1] + i
	}

	sum := func(move int64) int64 {
		return cache[move]
	}

	return fmt.Sprint(s.Solve(sum))
}
