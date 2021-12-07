package solutions

import (
	"fmt"
	"github.com/Virepri/adventofcode-2021/util"
	"math"
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

func (s *Day7Solution) Part1() string {
	var tested = map[int64]bool{}

	// the dumb solution is to just check every one
	min := int64(math.MaxInt64)
	for _, target := range s.crabs {
		if _, ok := tested[target]; ok {
			continue
		}
		tested[target] = true

		sum := int64(0)
		for _, crab := range s.crabs {
			sum += util.IntAbs(crab - target)
		}

		if sum < min {
			min = sum
		}
	}

	return fmt.Sprint(min)
}

func (s *Day7Solution) Solve(calculateFuel func(moveSize int64) int64) int64 {
	minPos := int64(math.MaxInt64)
	maxPos := int64(math.MinInt64)
	for _, v := range s.crabs {
		if v > maxPos {
			maxPos = v
		}
		if v < minPos {
			minPos = v
		}
	}

	getFullFuel := func(pos int64) int64 {
		sum := int64(0)
		for _, crab := range s.crabs {
			sum += calculateFuel(util.IntAbs(crab - pos))
		}

		return sum
	}

	// The output will be V-shaped. In other words, there's a minimum within the input. Let's start by finding the lowest point in the input.
	min := int64(math.MaxInt64)
	startPoint := int64(-1)
	for i := 0; i < len(s.crabs); i += len(s.crabs) / 20 { // look through every 5% of the output; we're bound to find something.
		sum := getFullFuel(s.crabs[i])
		if sum < min {
			min = sum
			startPoint = s.crabs[i]
		}
	}

	// Which direction towards the minimum?
	inc := util.TernaryInt64(getFullFuel(startPoint-1) < getFullFuel(startPoint+1), -1, 1)

	// Head towards the minimum.
	for {
		startPoint += inc

		if f := getFullFuel(startPoint); f < min {
			min = f
		} else {
			break
		}
	}

	return min
}

func (s *Day7Solution) Part2() string {
	calculateFuel := func(moveSize int64) int64 {
		sum := int64(0)
		for i := int64(1); i <= moveSize; i++ {
			sum += i
		}
		return sum
	}

	return fmt.Sprint(s.Solve(calculateFuel))
}
