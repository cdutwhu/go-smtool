package smtool

import rfl "reflect"

// F64s2Digits :
func F64s2Digits(in interface{}, typ rType) interface{} {
	inVal := rValueOf(in)
	L := inVal.Len()
	rtVal := rMakeSlice(rSliceOf(typ), L, L)
	for i := 0; i < L; i++ {
		iInVal := inVal.Index(i).Interface()
		iRtVal := rtVal.Index(i)
		switch typ.Kind() {
		case rfl.Int:
			iRtVal.Set(rValueOf(int(iInVal.(float64))))
		case rfl.Int8:
			iRtVal.Set(rValueOf(int8(iInVal.(float64))))
		case rfl.Int16:
			iRtVal.Set(rValueOf(int16(iInVal.(float64))))
		case rfl.Int32:
			iRtVal.Set(rValueOf(int32(iInVal.(float64))))
		case rfl.Int64:
			iRtVal.Set(rValueOf(int64(iInVal.(float64))))
		case rfl.Float32:
			iRtVal.Set(rValueOf(float32(iInVal.(float64))))
		case rfl.Float64:
			iRtVal.Set(rValueOf(float64(iInVal.(float64))))
		default:
			panic("input must be digits, OR add cases in F64s2Digits")
		}
	}
	return rtVal.Interface()
}

// Ints2Digits :
func Ints2Digits(in interface{}, typ rType) interface{} {
	inVal := rValueOf(in)
	L := inVal.Len()
	rtVal := rMakeSlice(rSliceOf(typ), L, L)
	for i := 0; i < L; i++ {
		iInVal := inVal.Index(i).Interface()
		iRtVal := rtVal.Index(i)
		switch typ.Kind() {
		case rfl.Int:
			iRtVal.Set(rValueOf(int(iInVal.(int))))
		case rfl.Int8:
			iRtVal.Set(rValueOf(int8(iInVal.(int))))
		case rfl.Int16:
			iRtVal.Set(rValueOf(int16(iInVal.(int))))
		case rfl.Int32:
			iRtVal.Set(rValueOf(int32(iInVal.(int))))
		case rfl.Int64:
			iRtVal.Set(rValueOf(int64(iInVal.(int))))
		case rfl.Float32:
			iRtVal.Set(rValueOf(float32(iInVal.(int))))
		case rfl.Float64:
			iRtVal.Set(rValueOf(float64(iInVal.(int))))
		default:
			panic("input must be digits, OR add cases in Ints2Digits")
		}
	}
	return rtVal.Interface()
}

// Conv2Ints : in order to use "go internal sort"
func Conv2Ints(slice interface{}) (rt []int) {
	valSlc := rValueOf(slice)
	L := valSlc.Len()
	switch valSlc.Interface().(type) {
	case []int:
		for i := 0; i < L; i++ {
			rt = append(rt, valSlc.Index(i).Interface().(int))
		}
	case []int8:
		for i := 0; i < L; i++ {
			rt = append(rt, int(valSlc.Index(i).Interface().(int8)))
		}
	case []int16:
		for i := 0; i < L; i++ {
			rt = append(rt, int(valSlc.Index(i).Interface().(int16)))
		}
	case []int32:
		for i := 0; i < L; i++ {
			rt = append(rt, int(valSlc.Index(i).Interface().(int32)))
		}
	case []int64:
		for i := 0; i < L; i++ {
			rt = append(rt, int(valSlc.Index(i).Interface().(int64)))
		}
	case []float32:
		for i := 0; i < L; i++ {
			rt = append(rt, int(valSlc.Index(i).Interface().(float32)))
		}
	case []float64:
		for i := 0; i < L; i++ {
			rt = append(rt, int(valSlc.Index(i).Interface().(float64)))
		}
	default:
		panic("")
	}
	return rt
}

// Conv2F64s : in order to use "go internal sort"
func Conv2F64s(slice interface{}) (rt []float64) {
	valSlc := rValueOf(slice)
	L := valSlc.Len()
	switch valSlc.Interface().(type) {
	case []int:
		for i := 0; i < L; i++ {
			rt = append(rt, float64(valSlc.Index(i).Interface().(int)))
		}
	case []int8:
		for i := 0; i < L; i++ {
			rt = append(rt, float64(valSlc.Index(i).Interface().(int8)))
		}
	case []int16:
		for i := 0; i < L; i++ {
			rt = append(rt, float64(valSlc.Index(i).Interface().(int16)))
		}
	case []int32:
		for i := 0; i < L; i++ {
			rt = append(rt, float64(valSlc.Index(i).Interface().(int32)))
		}
	case []int64:
		for i := 0; i < L; i++ {
			rt = append(rt, float64(valSlc.Index(i).Interface().(int64)))
		}
	case []float32:
		for i := 0; i < L; i++ {
			rt = append(rt, float64(valSlc.Index(i).Interface().(float32)))
		}
	case []float64:
		for i := 0; i < L; i++ {
			rt = append(rt, valSlc.Index(i).Interface().(float64))
		}
	default:
		panic("")
	}
	return rt
}
