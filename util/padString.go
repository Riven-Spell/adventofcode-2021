package util

import "strings"

func PadStringRight(toPad string, padWith rune, newSize int) string {
	return toPad + strings.Repeat(string(padWith), newSize - len(toPad))
}

func PadStringLeft(toPad string, padWith rune, newSize int) string {
	return strings.Repeat(string(padWith), newSize - len(toPad)) + toPad
}

