package solutions

import (
	"fmt"
	"sort"
	"strings"
)

type Day10Solution struct {
	lines []string
}

func (s *Day10Solution) Prepare(input string) {
	s.lines = strings.Split(input, "\n")
}

func (s *Day10Solution) Part1() string {
	openers := map[rune]bool{
		'(': true,
		'[': true,
		'{': true,
		'<': true,
	}

	closers := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	scores := map[rune]uint{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	score := uint(0)
	for _, v := range s.lines {
		stack := []rune{}

		for _, char := range v {
			if _, ok := openers[char]; ok {
				// add a thing to the stack, continue.
				stack = append(stack, char)
				continue
			}

			// We must exit something, OR we have invalid syntax.
			closer := closers[stack[len(stack)-1]]
			if char != closer {
				score += scores[char]
				break
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return fmt.Sprint(score)
}

func (s *Day10Solution) Part2() string {
	openers := map[rune]bool{
		'(': true,
		'[': true,
		'{': true,
		'<': true,
	}

	closers := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	scores := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}

	scoreList := make([]int, 0)
	for _, v := range s.lines {
		stack := []rune{}
		processing := true

		score := 0

		for _, char := range v {
			if _, ok := openers[char]; ok {
				// add a thing to the stack, continue.
				stack = append(stack, char)
				continue
			}

			// We must exit something, OR we have invalid syntax.
			closer := closers[stack[len(stack)-1]]
			if char != closer {
				processing = false
				break // this line is not to be processed.
			} else {
				stack = stack[:len(stack)-1]
			}
		}

		if processing { // did the line clear out correctly? What's left?
			for i := len(stack) - 1; i >= 0; i-- {
				score *= 5
				score += scores[stack[i]]
			}

			scoreList = append(scoreList, score)
		}
	}

	sort.Ints(scoreList)

	return fmt.Sprint(scoreList[(len(scoreList) / 2)])
}
