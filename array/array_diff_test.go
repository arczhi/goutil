package array

import (
	"testing"
)

var testTable = [][][]string{
	[][]string{[]string{"a", "b", "c"}, []string{"a", "b"}, []string{"c"}},
	[][]string{[]string{"a", "a", "b", "c"}, []string{"a", "b"}, []string{"c"}},
	[][]string{[]string{"a", "b", "c"}, []string{"b"}, []string{"a", "c"}},
	[][]string{[]string{"a", "b", "c"}, []string{"a", "b", "c"}, []string{}},
	[][]string{[]string{}, []string{"a"}, []string{"a"}},
	[][]string{[]string{"a"}, []string{}, []string{"a"}},
	[][]string{[]string{}, []string{}, []string{}},
}

func TestArrayDiff(t *testing.T) {
	for _, v := range testTable {
		resultArr := ArrayDIff(v[0], v[1])
		answer := v[2]
		if !isArrayElementEqual(resultArr, answer) {
			t.Errorf("want %v ,got %v", answer, resultArr)
			return
		}
		// fmt.Println(answer, resultArr)
	}
}
