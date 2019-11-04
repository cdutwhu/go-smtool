package smtool

import "testing"

func TestMapPrint(t *testing.T) {
	m := map[float32]string{-100: "A B", 2.01: "D", 3: "F", -1: "AAA"}
	MapPrint(m)
}

func TestKeys(t *testing.T) {
	m := map[int8]string{-3: "neg", -100: "A B", 2: "D", 3: "F", -1: "AAA"}
	m1 := map[float32]string{-100: "A B", 22.22: "D", 3: "F", -1: "AAA", 11.001: "11"}
	fPln(Keys(m))
	fPln(Keys(m1))
}

func TestMap(t *testing.T) {
	m := map[string]string{"zabc": "b", "f": "d", "e": "f"}
	fPln(Keys(m).([]string))
	m1 := map[int]string{11: "B", 2: "D", 3: "F", -1: "-1"}
	fPln(Keys(m1).([]int))

	fPln(" ---------------------------------------------- ")

	k, v := KVs(m)
	K, V := k.([]string), v.([]string)
	fPln(K)
	fPln(V)

	fPln(" ---------------------------------------------- ")

	k, v = KVs(m1)
	K1, V1 := k.([]int), v.([]string)
	fPln(K1)
	fPln(V1)

	fPln(" ---------------------------------------------- ")

	// m2 := map[string]string{"aa": "bb", "cc": "dd", "ee": "ff"}
	m3 := map[int]string{2: "B B   B C", -5: "DD", -1: "FF"}
	m13 := MapsJoin(m1, m3).(map[int]string)
	fPln(m13)

	fPln(" ********************************************** ")
	MapPrint(m13)

	m02 := MapsJoin(m3, m1).(map[int]string)
	fPln(m02)

	m4 := map[int]string{7: "BBB", 8: "DDD", 1: "FFF"}
	mm := MapsMerge(m1, m3, m4)
	fPln(mm)

	m4 = map[int]string{7: "BBB"}
	mm = MapsMerge(m4)
	fPln(mm)

	EE := MapsMerge()
	fPln(EE)
}
