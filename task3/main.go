package main

import (
	"fmt"
	"regexp"
	"sort"
)

func GroupAnagrams(words []string) map[string][]string {
	group := make(map[string][]string)
	//////////////////////////////
	valuesUnrepeat := make(map[string]struct{})
	//////////////////////////////

	for _, val := range words {
		if ok, err := regexp.MatchString(`^[a-z]+$`, val); !ok {
			if err != nil {
				panic("something wrong")
			}
			continue
		}
		////////////////////////////////////////////////////////
		if _, ok := valuesUnrepeat[val]; !ok {
			valuesUnrepeat[val] = struct{}{}
		} else {
			continue
		}
		///////////////////////////////////////////////////////
		arr := []byte(val)
		sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

		group[string(arr)] = append(group[string(arr)], val)

	}

	return group
}

func main() {
	slc := []string{"listen", "silent", "enlist", "inlets", "google", "gogole", "listen"}
	fmt.Println(GroupAnagrams(slc))

}
