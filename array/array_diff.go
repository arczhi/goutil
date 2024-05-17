package array

func ArrayDIff(fromArr, againstArr []string) []string {
	var resultArr []string
	mp := map[string]struct{}{}
	for _, v := range againstArr {
		mp[v] = struct{}{}
	}
	for _, v := range fromArr {
		if _, exist := mp[v]; !exist {
			resultArr = append(resultArr, v)
		}
	}
	return resultArr
}
