package array

import "testing"

func TestArrayUnique(t *testing.T) {
	resultArr := ArrayUnique([]string{"a", "c", "e", "a", "d", "a", "b", "c"})
	answer := []string{"a", "b", "c", "d", "e"}
	if !isArrayElementEqual(resultArr, answer) {
		t.Errorf("want %v ,got %v", answer, resultArr)
		return
	}
}
