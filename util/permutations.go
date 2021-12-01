package util

import (
	"math"
	"reflect"
)

func subPerms(vals []interface{}, permSize int) [][]interface{} {
	out := make([][]interface{}, int(math.Pow(float64(len(vals)), float64(permSize))))

	ot := 0
	ccmb := make([]int, permSize)

	for ccmb[permSize-1] != len(vals) - 1 {
		out[ot] = make([]interface{}, permSize)

		for k, v := range ccmb {
			out[ot][k] = v
		}

		ccmb[0]++
		for i := 0; ccmb[i] >= len(vals); i++ {
			ccmb[i] = 0
			if i + 1 < permSize {
				ccmb[i+1]++
			}
		}
	}

	return out
}

// takes in a list, hands back permutations of the list
func ListPerms(m interface{}, permSize int) [][]interface{} {
	mval := reflect.ValueOf(m)
	vals := make([]interface{}, mval.Len())

	for k := 0; k < mval.Len(); k++ {
		vals[k] = mval.Index(k).Interface()
	}

	return subPerms(vals, permSize)
}

// takes in a map, hands back key permutations
func KeyPerms(m interface{}, permSize int) [][]interface{} {
	mval := reflect.ValueOf(m)
	keys := make([]interface{}, mval.Len())

	for k,v := range mval.MapKeys() {
		keys[k] = v.Interface()
	}

	return subPerms(keys, permSize)
}
