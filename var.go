package smtool

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strings"
)

var (
	rMapOf        = reflect.MapOf
	rMakeMap      = reflect.MakeMap
	rValueOf      = reflect.ValueOf
	rTypeOf       = reflect.TypeOf
	rMakeSlice    = reflect.MakeSlice
	rNew          = reflect.New
	rAppend       = reflect.Append
	rAppendSlice  = reflect.AppendSlice
	rSliceOf      = reflect.SliceOf
	Max           = math.Max
	fPln          = fmt.Println
	fPf           = fmt.Printf
	fSf           = fmt.Sprintf
	fSp           = fmt.Sprint
	IntsAreSorted = sort.IntsAreSorted
	sSpl          = strings.Split
)

type (
	rType        = reflect.Type
	rSliceHeader = reflect.SliceHeader
)
