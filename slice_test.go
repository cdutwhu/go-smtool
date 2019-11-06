package smtool

import (
	"testing"
)

var (
	arrInt = []int{11, 22, 33, 44, 55}
	arrStr = []string{"s1", "s2", "s3", "s4", "s5"}
)

func TestIdx(t *testing.T) {
	t.Log(Idx('a', arrInt))
}

func TestIsIn(t *testing.T) {
	t.Log(IsIn('a', arrInt))
}

func TestSearch(t *testing.T) {
	t.Log(Search(arrInt, func(i int) bool {
		switch {
		case i == 11:
			return true
		case arrInt[i] == 222:
			return true
		}
		return false
	}))

	_, iDel, _ := Search(arrInt, func(i int) bool {
		switch {
		case i == 10:
			return true
		case arrInt[i] == 222:
			return true
		}
		return false
	})
	t.Log(Delete(arrInt, iDel))

}

func TestSearchAll(t *testing.T) {
	t.Log(SearchAll(arrInt, func(i int) bool {
		switch {
		case i == 0:
			return true
		case arrInt[i]+1 == 56:
			return true
		}
		return false
	}))

	_, iDel, _ := SearchAll(arrInt, func(i int) bool {
		switch {
		case i == 0:
			return true
		case arrInt[i]+1 == 56:
			return true
		}
		return false
	})

	t.Log(Delete(arrInt, iDel...))
}

func TestSetDel(t *testing.T) {
	strs := []string{}
	for i := 0; i < 1000000; i++ {
		strs = append(strs, fSf("%d", i))
	}

	SetDel(strs, 1)
	// Delete(strs, 1)
	fPln(strs[:10])
}
