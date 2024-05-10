package array

func ArrayUnique(arr []string) []string {
	var resultArr []string
	mp := map[string]struct{}{}
	for _, v := range arr {
		if _, exist := mp[v]; !exist {
			resultArr = append(resultArr, v)
			mp[v] = struct{}{}
		}
	}
	return resultArr
}
