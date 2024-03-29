package smtool

import (
	"regexp"
	"sort"
)

// MapPrint : Key Sorted Print
func MapPrint(m interface{}) {
	re := regexp.MustCompile(`^[+-]?[0-9]*\.?[0-9]+:`)
	mstr := fSp(m)
	mstr = mstr[4 : len(mstr)-1]
	// fPln(mstr)
	I := 0
	rmIlist := []int{}
	ss := sSpl(mstr, " ")
	for i, s := range ss {
		if re.MatchString(s) {
			I = i
		} else {
			ss[I] += " " + s
			rmIlist = append(rmIlist, i) // to be deleted (i)
		}
	}
	for i, s := range ss {
		if !IsIn(i, rmIlist) {
			fPln(s)
		}
	}
}

// Keys :
func Keys(m interface{}) interface{} {
	v := rValueOf(m)
	// pc(v.Kind() != ref.Map, fEf("NOT A MAP!"))
	keys := v.MapKeys()
	if L := len(keys); L > 0 {
		firstKey := keys[0].Interface()
		kType := rTypeOf(firstKey)
		rtVal := rMakeSlice(rSliceOf(kType), L, L)
		for i, k := range keys {
			rtVal.Index(i).Set(rValueOf(k.Interface()))
		}
		// sort keys if keys are sortable
		rt := rtVal.Interface()
		switch firstKey.(type) {
		case int:
			sort.Ints(rt.([]int))
		case float64:
			sort.Float64s(rt.([]float64))
		case string:
			sort.Strings(rt.([]string))
		case int64, int32, int16, int8, float32:
			F64s := NumTypeConv(rt, rTypeOf(float64(0))).([]float64)
			sort.Float64s(F64s)
			rt = NumTypeConv(F64s, rTypeOf(firstKey))
		}
		return rt
	}
	return nil
}

// KVs :
func KVs(m interface{}) (interface{}, interface{}) {
	v := rValueOf(m)
	// pc(v.Kind() != ref.Map, fEf("NOT A MAP!"))
	keys := v.MapKeys()
	if L := len(keys); L > 0 {
		kType := rTypeOf(keys[0].Interface())
		rtKey := rMakeSlice(rSliceOf(kType), L, L)
		vType := rTypeOf(v.MapIndex(keys[0]).Interface())
		rtVal := rMakeSlice(rSliceOf(vType), L, L)
		for i, k := range keys {
			rtKey.Index(i).Set(rValueOf(k.Interface()))
			rtVal.Index(i).Set(rValueOf(v.MapIndex(k).Interface()))
		}
		return rtKey.Interface(), rtVal.Interface()
	}
	return nil, nil
}

// MapsJoin : overwrited by the 2nd params
func MapsJoin(m1, m2 interface{}) interface{} {
	v1, v2 := rValueOf(m1), rValueOf(m2)
	// pc(v1.Kind() != ref.Map, fEf("m1 is NOT A MAP!"))
	// pc(v2.Kind() != ref.Map, fEf("m2 is NOT A MAP!"))
	keys1, keys2 := v1.MapKeys(), v2.MapKeys()
	if len(keys1) > 0 && len(keys2) > 0 {
		k1, k2 := keys1[0], keys2[0]
		k1Type, _ := rTypeOf(k1.Interface()), rTypeOf(k2.Interface())
		v1Type, _ := rTypeOf(v1.MapIndex(k1).Interface()), rTypeOf(v2.MapIndex(k2).Interface())
		// pc(k1Type != k2Type, fEf("different maps' key type!"))
		// pc(v1Type != v2Type, fEf("different maps' value type!"))
		aMap := rMakeMap(rMapOf(k1Type, v1Type))
		for _, k := range keys1 {
			aMap.SetMapIndex(rValueOf(k.Interface()), rValueOf(v1.MapIndex(k).Interface()))
		}
		for _, k := range keys2 {
			aMap.SetMapIndex(rValueOf(k.Interface()), rValueOf(v2.MapIndex(k).Interface()))
		}
		return aMap.Interface()
	}
	if len(keys1) > 0 && len(keys2) == 0 {
		return m1
	}
	if len(keys1) == 0 && len(keys2) > 0 {
		return m2
	}
	return m1
}

// MapsMerge : overwrited by the later params
func MapsMerge(ms ...interface{}) interface{} {
	if len(ms) == 0 {
		return nil
	}
	mm := ms[0]
	for i, m := range ms {
		if i >= 1 {
			mm = MapsJoin(mm, m)
		}
	}
	return mm
}
