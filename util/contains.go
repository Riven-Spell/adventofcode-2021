package util

import "reflect"

func ArrayContains(arr interface{}, x interface{}) bool {
	arrVal := reflect.ValueOf(arr)

	for i := 0; i < arrVal.Len(); i++ {
		if arrVal.Index(i).Interface() == x {
			return true
		}
	}

	return false
}
