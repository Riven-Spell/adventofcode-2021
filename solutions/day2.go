package solutions

import (
	"fmt"
	"strings"
	
	"github.com/Virepri/adventofcode-2021/util"
)

type Day2Solution struct {
	moves []util.Vector2D
}

func (s *Day2Solution) Prepare(input string) {
	if s.moves != nil {
		return
	}
	
	for _,v := range strings.Split(input, "\n") {
		move, size := "", int64(0)
		_, err := fmt.Sscanf(v, "%s %d", &move, &size)
		util.PanicIfErr(err)
		
		s.moves = append(s.moves, Movements[move].Mul(size))
	}
}

var Movements = map[string]util.Vector2D {
	"forward": {1, 0},
	"up": {0,-1},
	"down": {0,1},
}

func (s *Day2Solution) Part1() string {
	pos := util.Vector2D{}
	
	for _,v := range s.moves {
		pos = pos.Add(v)
	}
	
	return fmt.Sprint(pos.X * pos.Y)
}

func (s *Day2Solution) Part2() string {
	pos := util.Vector3D{} // X = horizontal, Y = aim, Z = depth
	
	reparse := func(move util.Vector2D) util.Vector3D {
		// remap
		out := util.Vector3D{
			X: move.X,
			Y: move.Y,
		}
		
		// depth changes with aim multiplied by X
		out.Z = pos.Y * out.X
		return out
	}
	
	for _,v := range s.moves {
		pos = pos.Add(reparse(v))
	}
	
	return fmt.Sprint(pos.X * pos.Z)
}
