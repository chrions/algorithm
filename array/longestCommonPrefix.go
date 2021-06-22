package main

func longestCommonPrefix(strs []string) string {
	m := map[string]int{}
	for _, str := range strs {
		for i, _ := range str {
			m[str[0:i+1]] += 1
		}
	}
	count := 0
	ret := ""
	for k, v := range m {
		if v >= len(strs) && len(k) > count {
			count = len(k)
			ret = k
		}
	}
	return ret
}
