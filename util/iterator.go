package util

type Iterator interface {
	Next() bool
	Value() interface{}
}
