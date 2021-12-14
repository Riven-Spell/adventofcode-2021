package solutions

import (
	"fmt"
	"github.com/Virepri/adventofcode-2021/util"
	"math"
	"strings"
)

type Day14Rule struct {
	Match  string
	Insert string
}

type Day14Solution struct {
	Template string
	Rules    []Day14Rule
}

func (s *Day14Solution) Prepare(input string) {
	if s.Template != "" {
		return
	}

	for k, v := range strings.Split(input, "\n") {
		if k == 0 {
			s.Template = v
			continue
		} else if k == 1 {
			continue
		}

		rule := Day14Rule{}

		_, err := fmt.Sscanf(v, "%s -> %s", &rule.Match, &rule.Insert)
		util.PanicIfErr(err)
		s.Rules = append(s.Rules, rule)
	}
}

func (s *Day14Solution) Solve(count uint) int {
	// simplify rule lookup
	ruleTable := map[string]string{}
	for _, v := range s.Rules {
		ruleTable[v.Match] = v.Insert
	}

	counts := map[uint8]int{}

	// establish existing pairs
	pairs := map[string]int{}
	for i := 0; i < len(s.Template)-1; i++ {
		pairs[string(s.Template[i])+string(s.Template[i+1])]++
		counts[s.Template[i]]++ // add counts for initial characters
	}
	counts[s.Template[len(s.Template)-1]]++ // capture the last element

	// handle each pair
	for i := uint(0); i < count; i++ {
		newPairs := map[string]int{}
		for k, v := range pairs {
			if rule, ok := ruleTable[k]; ok { // if we match a rule, execute it.
				newPairs[string(k[0])+rule] += v
				newPairs[rule+string(k[1])] += v
				counts[rule[0]] += v // every time we add a character, note the count.
			} else { // If we match no rules, do not execute it.
				newPairs[k] = v
			}
		}
		pairs = newPairs
	}

	min, max := math.MaxInt, math.MinInt
	for _, v := range counts {
		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	return max - min
}

func (s *Day14Solution) Part1() string {
	return fmt.Sprint(s.Solve(10))
}

func (s *Day14Solution) Part2() string {
	return fmt.Sprint(s.Solve(40))
	//return ""
}
