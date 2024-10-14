package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

/*
Напишите функцию GroupAnagrams(words []string) map[string][]string, которая принимает срез строк и группирует их в карту,
где каждый ключ — это отсортированная версия слова (в виде строки), а значение — срез всех анаграмм этого слова.
Пример: []string{"listen", "silent", "enlist", "inlets", "google", "gogole"}
должно вернуть map[string][]string{"eilnst": {"listen", "silent", "enlist", "inlets"}, "eggloo": {"google", "gogole"}}.
*/

func GroupAnagrams(words []string) map[string][]string {
	group := make(map[string][]string)
	//////////////////////////////
	valuesUnrepeat := make(map[string]struct{})
	//////////////////////////////

	for _, val := range words {
		if ok, err := regexp.MatchString(`^[a-zA-z]+$`, val); !ok {
			if err != nil {
				panic("something wrong")
			}
			continue
		}

		val = strings.ToLower(val)

		////////////////////////////////////////////////////////
		if _, ok := valuesUnrepeat[val]; !ok {
			valuesUnrepeat[val] = struct{}{}
		} else {
			continue
		}

		///////////////////////////////////////////////////////
		arr := []rune(val)
		sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

		group[string(arr)] = append(group[string(arr)], val)

	}

	return group
}

func main() {
	slc := []string{"listen", "siL11ent", "enlist", "inlets", "goOgle%", "gogole", "listen"}
	fmt.Println(GroupAnagrams(slc))

}
