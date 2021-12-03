package solutions

import (
	"github.com/Virepri/adventofcode-2021/util"
	"strconv"
	"strings"
)

type Day3Solution struct {
	input []string
}

func (s *Day3Solution) Prepare(input string) {
	if s.input == nil {
		s.input = strings.Split(input, "\n")
	}
}

type averageBit struct {
	count0 uint
	count1 uint
}

func (a averageBit) MostCommon() byte {
	if a.count0 > a.count1 {
		return '0'
	}

	return '1'
}

func (a averageBit) LeastCommon() byte {
	if a.count0 <= a.count1 {
		return '0'
	}

	return '1'
}

func (s *Day3Solution) GetAverages(input []string) []averageBit {
	out := make([]averageBit, len(s.input[0]))

	for _, v := range input {
		for k, b := range v {
			switch b {
			case '0':
				out[k].count0++
			case '1':
				out[k].count1++
			default:
				panic("wtf")
			}
		}
	}

	return out
}

func (s *Day3Solution) Part1() string {
	out := s.GetAverages(s.input)

	outGamma := ""
	outEpsilon := ""

	for _, v := range out {
		outGamma += string(v.MostCommon())
		outEpsilon += string(v.LeastCommon())
	}

	oG, err := strconv.ParseInt(outGamma, 2, 16)
	util.PanicIfErr(err)
	oE, err := strconv.ParseInt(outEpsilon, 2, 16)
	util.PanicIfErr(err)

	return strconv.FormatInt(oG*oE, 10)
}

func (s *Day3Solution) Part2() string {
	oxy, co2 := make([]string, len(s.input)), make([]string, len(s.input))
	copy(oxy, s.input)
	copy(co2, s.input)

	for k := range make([]bool, len(s.input[0])) {
		commons := s.GetAverages(oxy)
		v := commons[k]
		if len(oxy) > 1 {
			for i := 0; i < len(oxy); i++ {
				if len(oxy) == 1 {
					break
				}

				s := oxy[i]
				if v.MostCommon() != s[k] {
					// kill it
					oxy = append(oxy[:i], oxy[i+1:]...)
					i--
				}
			}
		}

		commons = s.GetAverages(co2)
		v = commons[k]
		if len(co2) > 1 {
			for i := 0; i < len(co2); i++ {
				if len(co2) == 1 {
					break
				}

				s := co2[i]
				if v.LeastCommon() != s[k] {
					// kill it
					co2 = append(co2[:i], co2[i+1:]...)
					i--
				}
			}
		}
	}

	oG, err := strconv.ParseInt(oxy[0], 2, 16)
	util.PanicIfErr(err)
	oE, err := strconv.ParseInt(co2[0], 2, 16)
	util.PanicIfErr(err)

	return strconv.FormatInt(oG*oE, 10)
}
