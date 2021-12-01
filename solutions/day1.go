package solutions

import (
	"fmt"
	"strconv"
	"strings"
	
	"github.com/Virepri/adventofcode-2021/util"
)

type Day1Solution struct {
	input []int64
}

func (s *Day1Solution) Prepare(input string) {
	if s.input != nil {
		return
	}
	
	s.input = make([]int64, 0)
	
	for _,v := range strings.Split(input, "\n") {
		i, err := strconv.ParseInt(strings.TrimSpace(v), 10, 64)
		util.PanicIfErr(err)
		
		s.input = append(s.input, i)
	}
}

func (s *Day1Solution) Part1() string {
	count := 0
	last := int64(-1)
	
	for _,v := range s.input {
		if last == -1 {
			last = v
			continue
		}
		
		if v > last {
			count++
		}
		
		last = v
	}
	
	return fmt.Sprint(count)
}

func (s *Day1Solution) Part2() string {
	sum := s.input[0] + s.input[1] + s.input[2]
	last := sum
	count := 0
	
	for k, v := range s.input[3:] {
		sum -= s.input[k]
		sum += v
		
		if sum > last {
			count++
		}
		
		last = sum
	}
	
	return fmt.Sprint(count)
}
