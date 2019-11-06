package smtool

// Idx : first element position in slice if found, otherwise, get -1
func Idx(e, slice interface{}) int {
	v := rValueOf(slice)
	L, I := v.Len(), -1
	for i := 0; i < L; i++ {
		if v.Index(i).Interface() == e {
			I = i
			break
		}
	}
	return I
}

// IsIn : whether element is in slice
func IsIn(e, slice interface{}) bool {
	return Idx(e, slice) >= 0
}

// Search : search the first chk-qualified element index from slice
func Search(slice interface{}, chk func(int) bool) (bool, int, interface{}) {
	v := rValueOf(slice)
	L, I, rt := v.Len(), -1, rNew(v.Type())
	for i := 0; i < L; i++ {
		if chk(i) {
			I, rt = i, v.Index(i)
			break
		}
	}
	return I != -1, I, rt.Interface()
}

// SearchAll :
func SearchAll(slice interface{}, chk func(int) bool) (bool, []int, interface{}) {
	v := rValueOf(slice)
	L, fdIdx := v.Len(), []int{}
	for i := 0; i < L; i++ {
		if chk(i) {
			fdIdx = append(fdIdx, i)
		}
	}
	rt := rMakeSlice(v.Type(), len(fdIdx), len(fdIdx))
	for i := 0; i < len(fdIdx); i++ {
		rt.Index(i).Set(v.Index(fdIdx[i]))
	}
	return len(fdIdx) != 0, fdIdx, rt.Interface()
}

// Delete :
func Delete(slice interface{}, lsPos ...int) (bool, interface{}) {
	v := rValueOf(slice)
	L, rt := v.Len(), rMakeSlice(v.Type(), 0, 1)
	for i := 0; i < L; i++ {
		if IsIn(i, lsPos) {
			continue
		}
		rt = rAppend(rt, v.Index(i))
	}
	return rt.Len() == 0, rt.Interface()
}

// Replace :
func Replace(slice interface{}, mPosVal map[int]interface{}) (bool, interface{}) {

	lsPos := Keys(mPosVal).([]int)

	v := rValueOf(slice)
	L, rt := v.Len(), rMakeSlice(v.Type(), 0, 1)
	for i := 0; i < L; i++ {
		if IsIn(i, lsPos) {
			continue
		}
		rt = rAppend(rt, v.Index(i))
	}
	return rt.Len() != 0, rt.Interface()

}

// // InsertBefore :
// func InsertBefore(slice interface{}, mPosVal map[int]interface{}) (bool, interface{}) {

// }

// // InsertAfter :
// func InsertAfter(slice interface{}, mPosVal map[int]interface{}) (bool, interface{}) {

// }

// ToSet :
func ToSet(slice interface{}) interface{} {
	v := rValueOf(slice)
	L, rt := v.Len(), rMakeSlice(v.Type(), 0, 1)
NEXT:
	for i := 0; i < L; i++ {
		av := v.Index(i)
		for j := 0; j < rt.Len(); j++ {
			if av.Interface() == rt.Index(j).Interface() {
				continue NEXT
			}
		}
		rt = rAppend(rt, av)
	}
	return rt.Interface()
}

// Contains :
func Contains(slice interface{}, lsE ...interface{}) bool {
	for _, e := range lsE {
		if !IsIn(e, slice) {
			return false
		}
	}
	return true
}

// SeqContains : ONLY apply to Distinct Element Slice
func SeqContains(slice interface{}, lsE ...interface{}) bool {
	ps := []int{}
	for _, e := range lsE {
		if i := Idx(e, slice); i >= 0 {
			ps = append(ps, i)
		} else {
			return false
		}
	}
	return IntsAreSorted(ps)
}

// ElesAreSame :
func ElesAreSame(slice interface{}) bool {
	return rValueOf(ToSet(slice)).Len() == 1
}

// // Intersect :
// func Intersect(s1, s2 interface{}) interface{} {
// 	v1, v2 := ValueOf(s1), ValueOf(s2)
// 	pc(v1.Kind() != rfl.Slice, fEf("s1 is NOT a SLICE!"))
// 	pc(v2.Kind() != rfl.Slice, fEf("s2 is NOT a SLICE!"))
// 	pc(v1.Type() != v2.Type(), fEf("s1 and s2 must be same type!"))
// 	l1, l2 := v1.Len(), v2.Len()
// 	rt := rfl.MakeSlice(v1.Type(), 0, 1)
// 	for i := 0; i < l1; i++ {

// 	}
// }

// // Union :
// func Union(s1, s2 interface{}) interface{} {
// 	v1, v2 := valueOf(s1), valueOf(s2)
// 	pc(v1.Kind() != rfl.Slice, fEf("s1 is NOT a SLICE!"))
// 	pc(v2.Kind() != rfl.Slice, fEf("s2 is NOT a SLICE!"))
// 	pc(v1.Type() != v2.Type(), fEf("s1 and s2 must be same type!"))
// 	l1, l2 := v1.Len(), v2.Len()
// }

// **********************

// Attach :
func Attach(slice1, slice2 interface{}, pos int) interface{} {
	v1, v2 := rValueOf(slice1), rValueOf(slice2)
	// pc(v1.Kind() != rfl.Slice, fEf("s1 is NOT a SLICE!"))
	// pc(v2.Kind() != rfl.Slice, fEf("s2 is NOT a SLICE!"))
	// pc(v1.Type() != v2.Type(), fEf("s1 and s2 must be same type!"))
	l1, l2 := v1.Len(), v2.Len()
	if l1 > 0 && l2 > 0 {
		if pos > l1 {
			return slice1
		}
		lm := int(Max(float64(l1), float64(l2+pos)))
		v := rAppendSlice(v1.Slice(0, pos), v2)
		return v.Slice(0, lm).Interface()
	}
	if l1 > 0 && l2 == 0 {
		return slice1
	}
	if l1 == 0 && l2 > 0 {
		return slice2
	}
	return slice1
}

// Cover :
func Cover(lsSlice ...interface{}) interface{} {
	if len(lsSlice) == 0 {
		return nil
	}
	attached := lsSlice[0]
	for i, s := range lsSlice {
		if i >= 1 {
			attached = Attach(attached, s, 0)
		}
	}
	return attached
}

// SetDel : disregard set order, set is used for "set"
func SetDel(set interface{}, p int) (bool, interface{}) {
	v := rValueOf(set)
	L := v.Len()
	if L > 0 && p >= 0 && p < L {
		v.Index(p).Set(v.Index(L - 1))
		return true, v.Slice(0, L-1).Interface()
	}
	return false, set
}
