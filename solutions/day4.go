package solutions

import (
	"fmt"
	"github.com/Virepri/adventofcode-2021/util"
	"strconv"
	"strings"
)

type BingoBoard struct {
	Numbers [5][5]int
	Hits    [5][5]bool
	Points  map[int][]int
}

func (b *BingoBoard) Hit(d int) {
	ps := b.Points[d]
	for i := 0; i < len(ps); i += 2 {
		b.Hits[ps[i]][ps[i+1]] = true
	}
}

func (b *BingoBoard) MapPoints() {
	b.Points = make(map[int][]int)

	for r, row := range b.Numbers {
		for c, num := range row {
			b.Points[num] = append(b.Points[num], r, c)
		}
	}
}

func (b *BingoBoard) Clone() BingoBoard {
	out := BingoBoard{
		Numbers: b.Numbers,
		Hits:    [5][5]bool{},
	}

	out.MapPoints()

	return out
}

func (b *BingoBoard) IsWin() bool {
	colWins := []bool{true, true, true, true, true}
	// test if any rows won (and set colWins)
	for row := 0; row < 5; row++ {
		rowWin := true
		for c, item := range b.Hits[row] {
			rowWin = rowWin && item
			colWins[c] = colWins[c] && item // If it was already winning and this is still winning, keep winning.
		}

		if rowWin {
			return true
		}
	}

	// test if any column won
	return colWins[0] || colWins[1] || colWins[2] || colWins[3] || colWins[4]
}

type Day4Solution struct {
	Numbers []int
	Boards  []BingoBoard
}

func (s *Day4Solution) Prepare(input string) {
	if s.Numbers != nil {
		return
	}

	lines := strings.Split(input, "\n")

	// parse the numbers
	for _, v := range strings.Split(lines[0], ",") {
		out, err := strconv.ParseInt(v, 10, 64)
		util.PanicIfErr(err)
		s.Numbers = append(s.Numbers, int(out))
	}

	// parse the boards
	cBoard := BingoBoard{}
	lIdx := 0
	for _, v := range lines[2:] {
		if v == "" || lIdx == 5 {
			cBoard.MapPoints()
			s.Boards = append(s.Boards, cBoard)
			cBoard = BingoBoard{}
			lIdx = 0
			continue
		}

		_, err := fmt.Sscanf(v, "%d %d %d %d %d", &cBoard.Numbers[lIdx][0], &cBoard.Numbers[lIdx][1], &cBoard.Numbers[lIdx][2], &cBoard.Numbers[lIdx][3], &cBoard.Numbers[lIdx][4])
		util.PanicIfErr(err)
		lIdx++
	}

	cBoard.MapPoints()
	s.Boards = append(s.Boards, cBoard)
	cBoard = BingoBoard{}
	lIdx = 0
}

func (s *Day4Solution) CloneInput() []BingoBoard {
	out := make([]BingoBoard, len(s.Boards))

	for k, v := range s.Boards {
		out[k] = v.Clone()
	}

	return out
}

func (s *Day4Solution) Part1() string {
	input := s.CloneInput()
	winrar := -1
	winNum := -1

	for _, v := range s.Numbers {
		for bIdx := range input {
			input[bIdx].Hit(v)

			if input[bIdx].IsWin() {
				winrar = bIdx
				winNum = v
				goto calcScore
			}
		}
	}

calcScore:
	if winrar == -1 || winNum == -1 {
		panic("winrar is -1")
	}

	sum := 0
	for x := range input[winrar].Hits {
		for y := range input[winrar].Hits[x] {
			if !input[winrar].Hits[x][y] {
				sum += input[winrar].Numbers[x][y]
			}
		}
	}

	return fmt.Sprint(sum * winNum)
}

func (s *Day4Solution) Part2() string {
	input := s.CloneInput()
	winBoard := BingoBoard{}
	winrar := -1
	winNum := -1

	for _, v := range s.Numbers {
		for bIdx := 0; bIdx < len(input); bIdx++ {
			input[bIdx].Hit(v)

			if input[bIdx].IsWin() {
				winrar = bIdx
				winNum = v
				winBoard = input[bIdx]
				input = append(input[:bIdx], input[bIdx+1:]...)
				bIdx--
			}
		}
	}

	//calcScore:
	if winrar == -1 || winNum == -1 {
		panic("winrar is -1")
	}

	sum := 0
	for x := range winBoard.Hits {
		for y := range winBoard.Hits[x] {
			if !winBoard.Hits[x][y] {
				sum += winBoard.Numbers[x][y]
			}
		}
	}

	return fmt.Sprint(sum * winNum)
}
