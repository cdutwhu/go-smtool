package smtool

import "testing"

func TestNumTypeConv(t *testing.T) {
	I8s := []int8{1, 2, 3, 4, 5}
	Output := NumTypeConv(I8s, TypFloat64).([]float64)
	fPln(Output)
}
