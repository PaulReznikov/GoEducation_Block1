package main

import (
	"fmt"
	"sort"
)

func remove(slcStr []string, idx int) []string {
	return append(slcStr[:idx], slcStr[idx+1:]...)
}

func SortStrings(slcStr []string) string {
	m := make(map[string]struct{})
	resStr := ""
	for i := len(slcStr) - 1; i >= 0; i-- {
		if _, ok := m[slcStr[i]]; !ok {
			m[slcStr[i]] = struct{}{}
			continue
		} else {
			slcStr = remove(slcStr, i)
		}
	}

	sort.Slice(slcStr, func(i, j int) bool { return len([]rune(slcStr[i])) < len([]rune(slcStr[j])) })
	for i, val := range slcStr {
		if i == len(slcStr)-1 {
			resStr += val
			continue
		}
		resStr += val + ", "
	}

	return resStr
}

func main() {
	slc := []string{"ccc", "aaa", "bb", "dddd", "aaa", "c", "ffffffffff", "aaa", "c", "bbb"}
	fmt.Println(SortStrings(slc))
}
