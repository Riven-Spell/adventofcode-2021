package util

import (
	"reflect"
)

// Unionizer changed with the 2021 boilerplate-- Now, it can accept either side, at any time.
type Unionizer struct {
	// A & B are uniques, Union is the joined union
	A, B, Union map[interface{}]bool
}

func NewUnion() Unionizer {
	return Unionizer{
		A: map[interface{}]bool{},
		B: map[interface{}]bool{},
		Union: map[interface{}]bool{},
	}
}

func (u *Unionizer) Len() int {
	return len(u.A) + len(u.B) + len(u.Union)
}

func (u *Unionizer) Join(u2 Unionizer) Unionizer {
	out := NewUnion()
	
	out.AddItemsA(u.GetUnion())
	out.AddItemsB(u2.GetUnion())
	
	return out
}

func (u *Unionizer) Contains(i interface{}) bool {
	_, okA := u.A[i]
	_, okB := u.B[i]
	_, okU := u.Union[i]
	
	return okA || okB || okU
}

func (u *Unionizer) ForEachUnique(each func(i interface{}) bool) {
	for k, _ := range u.A {
		if !each(k) {
			break
		}
	}
	
	for k, _ := range u.B {
		if !each(k) {
			break
		}
	}
}

func (u *Unionizer) ForEachUnion(each func(i interface{}) bool) {
	for v := range u.Union {
		if !each(v) {
			break
		}
	}
}

func (u *Unionizer) ForEach(each func(i interface{}) bool) {
	u.ForEachUnique(each)
	u.ForEachUnion(each)
}

func (u *Unionizer) GetUnion() []interface{} {
	out := make([]interface{}, 0)
	
	for k := range u.Union {
		out = append(out, k)
	}

	return out
}

func (u *Unionizer) RemoveItems(itemList interface{}) {
	rList := reflect.ValueOf(itemList)
	rLen := rList.Len()

	for i := 0; i < rLen; i++ {
		val := rList.Index(i).Interface()

		_, okA := u.A[val]

		if okA {
			delete(u.A, val)
			continue
		}
		
		_, okB := u.B[val]
		
		if okB {
			delete(u.B, val)
		}
		
		_, okU := u.Union[val]
		
		if okU {
			delete(u.Union, val)
		}
	}
}

func (u *Unionizer) AddItemsA(itemList interface{}) {
	rList := reflect.ValueOf(itemList)
	rLen := rList.Len()

	for i := 0; i < rLen; i++ {
		val := rList.Index(i).Interface()
		
		_,okB := u.B[val]
		
		if okB {
			delete(u.B, val)
			u.Union[val] = true
		} else {
			u.A[val] = true
		}
	}
}

func (u *Unionizer) AddItemsB(itemList interface{}) {
	rList := reflect.ValueOf(itemList)
	rLen := rList.Len()
	
	for i := 0; i < rLen; i++ {
		val := rList.Index(i).Interface()
		
		_,okA := u.A[val]
		
		if okA {
			delete(u.A, val)
			u.Union[val] = true
		} else {
			u.B[val] = true
		}
	}
}
