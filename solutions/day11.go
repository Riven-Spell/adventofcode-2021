package solutions

import (
	"fmt"
	"github.com/Virepri/adventofcode-2021/util"
	"strconv"
	"strings"
)

type Day11Solution struct {
	octopi [][]int
}

func (s *Day11Solution) Prepare(input string) {
	if s.octopi != nil {
		return
	}

	s.octopi = append(s.octopi, []int{})
	for _, v := range strings.Split(input, "") {
		if v == "\n" {
			s.octopi = append(s.octopi, []int{})
			continue
		}

		height, err := strconv.ParseInt(v, 10, 64)
		util.PanicIfErr(err)
		s.octopi[len(s.octopi)-1] = append(s.octopi[len(s.octopi)-1], int(height))
	}
}

func (s *Day11Solution) CloneInput() [][]int {
	out := make([][]int, len(s.octopi))

	for k, v := range s.octopi {
		out[k] = make([]int, len(v))
		copy(out[k], v)
	}

	return out
}

func (s *Day11Solution) PrintBoard(state [][]int) {
	for _, row := range state {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Println()
	}
	fmt.Println()
}

func (s *Day11Solution) Part1() string {
	state := s.CloneInput()

	totalFlashes := 0
	for i := 0; i < 100; i++ { // one-hundred steps
		//s.PrintBoard(state)

		flashing := make([]util.Vector2D, 0)
		// octopi can only flash once per step
		flashed := map[util.Vector2D]bool{}

		for y, row := range state { // in the very beginning, we only increment every state.
			for x := range row {
				state[y][x]++

				if state[y][x] > 9 {
					toFlash := util.Vector2D{X: int64(x), Y: int64(y)}
					flashing = append(flashing, toFlash)
					flashed[toFlash] = true
				}
			}
		}

		for len(flashing) > 0 {
			// Pop the queue
			val := flashing[0]
			flashing = flashing[1:]
			// Count the flash
			totalFlashes++
			state[val.Y][val.X] = 0
			// Increase the energy of surroundings, flash those if needbe.
			for _, v := range util.AdjacentVector2D {
				target := val.Add(v)

				// do not target adjacents outside the map
				if target.X < 0 || int(target.X) >= len(s.octopi[0]) || target.Y < 0 || int(target.Y) >= len(s.octopi) {
					continue
				}

				if _, ok := flashed[target]; ok {
					continue // do not increase energy, do not flash.
				}

				state[target.Y][target.X]++

				if state[target.Y][target.X] > 9 { // if it can flash, add it to the queue & mark it as flashed
					flashing = append(flashing, target)
					flashed[target] = true
				}
			}
		}
	}

	return fmt.Sprint(totalFlashes)
}

func (s *Day11Solution) Part2() string {
	state := s.CloneInput()

	i := 0
	for { // Go until they all flash at once.
		//s.PrintBoard(state)

		flashing := make([]util.Vector2D, 0)
		// octopi can only flash once per step
		flashed := map[util.Vector2D]bool{}

		for y, row := range state { // in the very beginning, we only increment every state.
			for x := range row {
				state[y][x]++

				if state[y][x] > 9 {
					toFlash := util.Vector2D{X: int64(x), Y: int64(y)}
					flashing = append(flashing, toFlash)
					flashed[toFlash] = true
				}
			}
		}

		totalFlashes := 0
		for len(flashing) > 0 {
			// Pop the queue
			val := flashing[0]
			flashing = flashing[1:]
			// Count the flash
			totalFlashes++
			state[val.Y][val.X] = 0
			// Increase the energy of surroundings, flash those if needbe.
			for _, v := range util.AdjacentVector2D {
				target := val.Add(v)

				// do not target adjacents outside the map
				if target.X < 0 || int(target.X) >= len(s.octopi[0]) || target.Y < 0 || int(target.Y) >= len(s.octopi) {
					continue
				}

				if _, ok := flashed[target]; ok {
					continue // do not increase energy, do not flash.
				}

				state[target.Y][target.X]++

				if state[target.Y][target.X] > 9 { // if it can flash, add it to the queue & mark it as flashed
					flashing = append(flashing, target)
					flashed[target] = true
				}
			}
		}

		i++

		if totalFlashes == len(s.octopi)*len(s.octopi[0]) {
			return fmt.Sprint(i)
		}
	}
}
