package array

func ArrayDIff(arr1, arr2 []string) []string {
	if len(arr1) >= len(arr2) {
		return arrayDiff(arr1, arr2)
	} else {
		return arrayDiff(arr2, arr1)
	}
}

func arrayDiff(longArr, shortArr []string) []string {
	var resultArr []string
	mp := map[string]struct{}{}
	for _, v := range shortArr {
		mp[v] = struct{}{}
	}
	for _, v := range longArr {
		if _, exist := mp[v]; !exist {
			resultArr = append(resultArr, v)
		}
	}
	return resultArr
}
