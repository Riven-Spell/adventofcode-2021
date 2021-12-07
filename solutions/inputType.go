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

var RegisteredDays = []struct {
	Solution        Solution // sample solution can be found in ./sampleday.go
	StringInput     *string  // inputs should exist in ../inputs and be a single var in a single file. These are just default inputs.
	ExpectedOutputs []string // these determine pass/failure in case I come back to try and optimize a solution and fuck it up.
}{
	{
		Solution:        &Day1Solution{},
		StringInput:     &inputs.Day1,
		ExpectedOutputs: []string{"1226", "1252"},
	},
	{
		Solution:        &Day2Solution{},
		StringInput:     &inputs.Day2,
		ExpectedOutputs: []string{"1459206", "1320534480"},
	},
	{
		Solution:        &Day3Solution{},
		StringInput:     &inputs.Day3,
		ExpectedOutputs: []string{"2640986", "6822109"},
	},
	{
		Solution:        &Day4Solution{},
		StringInput:     &inputs.Day4,
		ExpectedOutputs: []string{"2745", "6594"},
	},
	{
		Solution:        &Day5Solution{},
		StringInput:     &inputs.Day5,
		ExpectedOutputs: []string{"5585", "17193"},
	},
	{
		Solution:        &Day6Solution{},
		StringInput:     &inputs.Day6,
		ExpectedOutputs: []string{"355386", "1613415325809"},
	},
	{
		Solution:        &Day7Solution{},
		StringInput:     &inputs.Day7,
		ExpectedOutputs: []string{"345035", "97038163"},
	},
}
