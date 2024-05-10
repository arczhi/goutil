package array

import "testing"

func TestArrayMerge(t *testing.T) {
	resultArr := ArrayMerge([]string{"a", "b"}, []string{"c"})
	answer := []string{"a", "b", "c"}
	if !isArrayElementEqual(resultArr, answer) {
		t.Errorf("want %v ,got %v", answer, resultArr)
		return
	}
}
