package solutions

import (
	"fmt"
	"github.com/Virepri/adventofcode-2021/util"
	"strconv"
	"strings"
)

type Day8Solution struct {
	Input [][2][]string
}

type SevenSegment struct {
	/*
		 0000
		1    2
		1    2
		 3333
		4    5
		4    5
		 6666
	*/

	lights [7]bool
}

func (s *Day8Solution) Prepare(input string) {
	if s.Input != nil {
		return
	}

	for _, v := range strings.Split(input, "\n") {
		out := [2][]string{}
		in := strings.Split(v, " | ")

		out[0] = strings.Split(in[0], " ")
		out[1] = strings.Split(in[1], " ")

		s.Input = append(s.Input, out)
	}
}

func (s *Day8Solution) Part1() string {
	// unique values:
	/*
		1 = 2
		4 = 4
		7 = 3
		8 = 7
	*/

	uniqueCounts := map[int]int{
		2: 1,
		4: 4,
		3: 7,
		7: 8,
	}

	count := 0
	for _, v := range s.Input {
		for _, dig := range v[1] {
			if _, ok := uniqueCounts[len(dig)]; ok {
				count++
			}
		}
	}

	return fmt.Sprint(count)
}

func (s *Day8Solution) Part2() string {
	numberPins := map[int][]int{
		0: {0, 1, 2, 4, 5, 6},
		1: {2, 5},
		2: {0, 2, 3, 4, 6},
		3: {0, 2, 3, 5, 6},
		4: {1, 2, 3, 5},
		5: {0, 1, 3, 5, 6},
		6: {0, 1, 3, 4, 5, 6},
		7: {0, 2, 5},
		8: {0, 1, 2, 3, 4, 5, 6},
		9: {0, 1, 2, 3, 5, 6},
	}

	var digits = map[[7]bool]string{}

	for k, v := range numberPins {
		key := [7]bool{}

		for _, b := range v {
			key[b] = true
		}

		digits[key] = fmt.Sprint(k)
	}

	sum := int64(0)
	for _, v := range s.Input {
		// Map values to counts.
		counts := map[int][]string{}
		for _, item := range append(append([]string{}, v[0]...), v[1]...) {
			counts[len(item)] = append(counts[len(item)], item)
		}

		pins := map[byte]int{}

		// find 0 for sure.
		zero := util.GetUniqueChars(counts[2], counts[3])
		if len(zero) != 1 {
			panic(fmt.Sprint("only one value should be present! ", zero))
		}
		pins[zero[0]] = 0

		// find 1, 2, 4, 5
		fiveLen := util.GetUniqueChars(counts[5])

		// find 2, 3, 4
		sixLen := util.GetUniqueChars(counts[6])

		// find 3
		three := ""
		for k := range util.QuickUnion([]byte(fiveLen), []byte(sixLen)).B {
			if three != "" {
				panic("3 should only have one value present, also found " + string(k.(byte)))
			}

			pins[k.(byte)] = 3
			three += string(k.(byte))
		}

		// find 5
		five := ""
		for k := range util.QuickUnion([]byte(sixLen), []byte(counts[2][0])).B {
			if five != "" {
				panic("5 should only have one value present, also found " + string(k.(byte)))
			}

			pins[k.(byte)] = 5
			five += string(k.(byte))
		}

		// find 2
		two := ""
		for k := range util.QuickUnion([]byte(five), []byte(counts[2][0])).B {
			if two != "" {
				panic("2 should only have one value present, also found " + string(k.(byte)))
			}

			pins[k.(byte)] = 2
			two += string(k.(byte))
		}

		// find 4
		four := ""
		for k := range util.QuickUnion([]byte(two+three), []byte(sixLen)).B {
			if four != "" {
				panic("4 should only have one value present, also found " + string(k.(byte)))
			}

			pins[k.(byte)] = 4
			four += string(k.(byte))
		}

		// find 1
		one := ""
		for k := range util.QuickUnion([]byte(two+four+five), []byte(fiveLen)).B {
			if one != "" {
				panic("1 should only have one value present, also found " + string(k.(byte)))
			}

			pins[k.(byte)] = 1
			one += string(k.(byte))
		}

		// find 6
		six := ""
		for k := range util.QuickUnion([]byte(zero+one+two+three+four+five), []byte(counts[6][0])).B { // 6-lengths are guaranteed to have a 6.
			if six != "" {
				panic("6 should only have one value present, also found " + string(k.(byte)))
			}

			pins[k.(byte)] = 6
			six += string(k.(byte))
		}

		// now that we have all 7 pins for certain, we can translate!
		num := ""
		for _, encoded := range v[1] {
			// first, let's convert to our light sets.
			lights := [7]bool{}
			for _, v := range []byte(encoded) {
				lights[pins[v]] = true
			}

			// now, let's jump from our representation to a number.
			dig, ok := digits[lights]

			if !ok {
				panic(fmt.Sprint(lights))
			}

			num += dig
		}

		o, err := strconv.ParseInt(num, 10, 64)
		if err != nil {
			panic(err)
		}

		sum += o
	}

	return fmt.Sprint(sum)
}
