package smtool

// NumTypeConv :
func NumTypeConv(in interface{}, typ rType) interface{} {
	inVal := rValueOf(in)
	L := inVal.Len()
	rtVal := rMakeSlice(rSliceOf(typ), L, L)
	for i := 0; i < L; i++ {
		iInVal := inVal.Index(i)
		IiInVal := iInVal.Interface()
		iRtVal := rtVal.Index(i)
		if rTypeOf(IiInVal).ConvertibleTo(typ) {
			iRtVal.Set(iInVal.Convert(typ))
		}
	}
	return rtVal.Interface()
}
