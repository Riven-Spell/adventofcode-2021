package solutions

import (
	"github.com/Virepri/adventofcode-2021/inputs"
)

// Solution is an interface that exposes a simple structure:
// The runner should call Prepare() on it to prepare the input
// Then call the part 1 and part 2 functions if wanted.
type Solution interface {
	Prepare(input string)
	Part1() string
	Part2() string
}

var RegisteredDays = []struct{
	Solution        Solution // sample solution can be found in ./sampleday.go
	StringInput     *string // inputs should exist in ../inputs and be a single var in a single file. These are just default inputs.
	ExpectedOutputs []string // these determine pass/failure in case I come back to try and optimize a solution and fuck it up.
}{
	{
		Solution: &Day1Solution{},
		StringInput: &inputs.Day1,
		ExpectedOutputs: []string{"1226", "1252"},
	},
}