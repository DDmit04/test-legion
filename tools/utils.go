package tools

func UniqueIntsList(intSlice []int) []int {
	uniqueMap := make(map[int]bool)
	var list []int
	for _, entry := range intSlice {
		uniqueMap[entry] = true
	}
	for key := range uniqueMap {
		list = append(list, key)
	}
	return list
}

func UniqueStringsList(intSlice []string) []string {
	uniqueMap := make(map[string]bool)
	var list []string
	for _, entry := range intSlice {
		uniqueMap[entry] = true
	}
	for key := range uniqueMap {
		list = append(list, key)
	}
	return list
}
