package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

/*
Работа с анограммами:
Напишите функцию, которая принимает строку и возвращает map, где ключом является слово,
а значением — список всех анаграмм этого слова, найденных в строке.
*/

func WorkingWithAnagrams(startStr string) map[string][]string {

	resultMap := make(map[string][]string)
	ValuesUnrepeat := make(map[string]struct{})
	words := strings.Fields(startStr)

	for _, word := range words {

		//checking for non-letter characters
		if ok, err := regexp.MatchString(`^[a-zA-z]+$`, word); !ok {
			if err != nil {
				panic("something wrong")
			}
			continue
		}

		word = strings.ToLower(word)

		//Checking for duplicate characters
		if _, ok := ValuesUnrepeat[word]; ok {
			continue
		} else {
			ValuesUnrepeat[word] = struct{}{}
		}

		sortWord := []rune(word)
		sort.Slice(sortWord, func(i, j int) bool { return sortWord[i] < sortWord[j] })

		resultMap[string(sortWord)] = append(resultMap[string(sortWord)], word)
	}

	return resultMap

}

func main() {
	str := " listen silEnt enlist inlets google! gogole    inlets"
	fmt.Println(WorkingWithAnagrams(str))
}
