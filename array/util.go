package array

func isArrayElementEqual(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	mp := map[string]struct{}{}
	for _, v := range arr1 {
		mp[v] = struct{}{}
	}
	for _, v := range arr2 {
		if _, exist := mp[v]; !exist {
			return false
		}
	}
	return true
}
