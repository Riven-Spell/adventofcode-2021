package util

import "math"

var deg2Rad = math.Pi / 180

func EnsureRotation360(rot int) int {
	for rot < 0 || rot >= 360 {
		if rot < 0 {
			rot += 360
		} else {
			rot -= 360
		}
	}

	return rot
}
