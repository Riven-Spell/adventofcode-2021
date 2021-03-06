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
	{
		Solution:        &Day8Solution{},
		StringInput:     &inputs.Day8,
		ExpectedOutputs: []string{"534", "1070188"},
	},
	{
		Solution:        &Day9Solution{},
		StringInput:     &inputs.Day9,
		ExpectedOutputs: []string{"506", "931200"},
	},
	{
		Solution:        &Day10Solution{},
		StringInput:     &inputs.Day10,
		ExpectedOutputs: []string{"464991", "3662008566"},
	},
	{
		Solution:        &Day11Solution{},
		StringInput:     &inputs.Day11,
		ExpectedOutputs: []string{"1732", "290"},
	},
	{
		Solution:        &Day12Solution{},
		StringInput:     &inputs.Day12,
		ExpectedOutputs: []string{"3708", "93858"},
	},
	{
		Solution:        &Day13Solution{},
		StringInput:     &inputs.Day13,
		ExpectedOutputs: []string{"710", Day13Part2Output},
	},
	{
		Solution:        &Day14Solution{},
		StringInput:     &inputs.Day14,
		ExpectedOutputs: []string{"", ""},
	},
}
