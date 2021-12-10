package solutions

import (
	"fmt"
	"github.com/Virepri/adventofcode-2021/util"
	"sort"
	"strconv"
	"strings"
)

type Day9Solution struct {
	heightmap [][]int
}

func (s *Day9Solution) Prepare(input string) {
	if s.heightmap != nil {
		return
	}

	s.heightmap = append(s.heightmap, []int{})
	for _, v := range strings.Split(input, "") {
		if v == "\n" {
			s.heightmap = append(s.heightmap, []int{})
			continue
		}

		height, err := strconv.ParseInt(v, 10, 64)
		util.PanicIfErr(err)
		s.heightmap[len(s.heightmap)-1] = append(s.heightmap[len(s.heightmap)-1], int(height))
	}
}

func (s *Day9Solution) Part1() string {
	totalRisk := 0

	for y := range s.heightmap {
		for x, v := range s.heightmap[y] {
			min := 10 // heights are no higher than 9
			pos := util.Vector2D{X: int64(x), Y: int64(y)}
			for _, v := range util.ImmediateAdjacentVector2D {
				toCheck := pos.Add(v)

				// do not check points outside of the map
				if toCheck.X < 0 || int(toCheck.X) >= len(s.heightmap[y]) || toCheck.Y < 0 || int(toCheck.Y) >= len(s.heightmap) {
					continue
				}

				if h := s.heightmap[toCheck.Y][toCheck.X]; h < min {
					min = h
				}
			}

			if v < min {
				totalRisk += 1 + v
			}
		}
	}

	return fmt.Sprint(totalRisk)
}

func (s *Day9Solution) Part2() string {
	basins := []int{}

	for y := range s.heightmap {
		for x, v := range s.heightmap[y] {
			isLowPoint := true
			pos := util.Vector2D{X: int64(x), Y: int64(y)}
			for _, adj := range util.ImmediateAdjacentVector2D {
				toCheck := pos.Add(adj)

				// do not check points outside of the map
				if toCheck.X < 0 || int(toCheck.X) >= len(s.heightmap[y]) || toCheck.Y < 0 || int(toCheck.Y) >= len(s.heightmap) {
					continue
				}

				if h := s.heightmap[toCheck.Y][toCheck.X]; v >= h {
					isLowPoint = false
					break
				}
			}

			if !isLowPoint {
				// do not process things that aren't low points.
				continue
			}

			type queueEntry struct {
				cHeight int
				pos     util.Vector2D // the next position to be checked.
			}

			seen := make(map[util.Vector2D]bool)
			pQueue := []queueEntry{{cHeight: v - 1, pos: pos}}
			size := 0

			for len(pQueue) > 0 {
				workItem := pQueue[0]
				pQueue = pQueue[1:]

				if _, ok := seen[workItem.pos]; ok {
					continue
				}

				if h := s.heightmap[workItem.pos.Y][workItem.pos.X]; h == 9 || h <= workItem.cHeight {
					continue
				}

				seen[workItem.pos] = true

				if workItem.cHeight < 9 {
					for _, v := range util.ImmediateAdjacentVector2D {
						newWI := queueEntry{pos: workItem.pos.Add(v), cHeight: workItem.cHeight + 1}

						// do not check points outside of the map
						if newWI.pos.X < 0 || int(newWI.pos.X) >= len(s.heightmap[y]) || newWI.pos.Y < 0 || int(newWI.pos.Y) >= len(s.heightmap) {
							continue
						}

						pQueue = append(pQueue, newWI)
					}
				}

				seen[workItem.pos] = true

				size++
			}

			// insert the basin
			basins = append(basins, size)
		}
	}

	sort.Ints(basins)

	return fmt.Sprint(basins[len(basins)-3] * basins[len(basins)-2] * basins[len(basins)-1])
}
