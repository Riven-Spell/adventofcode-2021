package solutions

import (
	"fmt"
	"strings"
)

type Day12Solution struct {
	// from mapped to to
	paths map[string][]string
}

func (s *Day12Solution) Prepare(input string) {
	if s.paths != nil {
		return
	}

	s.paths = make(map[string][]string)

	appendIfNecessary := func(key, value string) {
		paths := s.paths[key]

		for _, v := range paths {
			if v == value {
				return
			}
		}

		s.paths[key] = append(s.paths[key], value)
	}

	for _, v := range strings.Split(input, "\n") {
		fromTo := strings.Split(v, "-")

		appendIfNecessary(fromTo[0], fromTo[1])
		appendIfNecessary(fromTo[1], fromTo[0])
	}
}

type Day12TravelState struct {
	seenSmall  map[string]bool
	doubleSeen bool
	path       []string
}

func (t Day12TravelState) Clone() Day12TravelState {
	out := Day12TravelState{
		seenSmall:  make(map[string]bool),
		path:       make([]string, len(t.path)),
		doubleSeen: t.doubleSeen,
	}

	copy(out.path, t.path)

	for k, v := range t.seenSmall {
		out.seenSmall[k] = v
	}

	return out
}

func (s *Day12Solution) Part1() string {
	pathCount := 0
	states := []Day12TravelState{
		{
			path:      []string{"start"},
			seenSmall: make(map[string]bool),
		},
	}

	for len(states) > 0 {
		cState := states[0]
		states = states[1:]

		cLoc := cState.path[len(cState.path)-1]
		for _, v := range s.paths[cLoc] {
			if v == "start" {
				// No need to go back.
				continue
			}

			lower := strings.ToLower(v)
			if lower == v && cState.seenSmall[v] {
				// No need to go back.
				continue
			}

			if v == "end" {
				// This is a valid path.
				pathCount++
			} else {
				// We need to go deeper.
				newState := cState.Clone()
				newState.path = append(newState.path, v)
				if lower == v {
					newState.seenSmall[v] = true
				}

				states = append(states, newState)
			}
		}
	}

	return fmt.Sprint(pathCount)
}

func (s *Day12Solution) Part2() string {
	pathCount := 0
	states := []Day12TravelState{
		{
			path:      []string{"start"},
			seenSmall: make(map[string]bool),
		},
	}

	for len(states) > 0 {
		cState := states[0]
		states = states[1:]

		cLoc := cState.path[len(cState.path)-1]
		for _, v := range s.paths[cLoc] {
			if v == "start" {
				// No need to go back.
				continue
			}

			lower := strings.ToLower(v)
			if lower == v && cState.seenSmall[v] {
				// Have we seen a small one twice yet?
				if !cState.doubleSeen {
					newState := cState.Clone()
					newState.path = append(newState.path, v)
					newState.doubleSeen = true

					states = append(states, newState)
				}

				continue
			}

			if v == "end" {
				// This is a valid path.
				pathCount++
			} else {
				// We need to go deeper.
				newState := cState.Clone()
				newState.path = append(newState.path, v)
				if lower == v {
					newState.seenSmall[v] = true
				}

				states = append(states, newState)
			}
		}
	}

	return fmt.Sprint(pathCount)
}
