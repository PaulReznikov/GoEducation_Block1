package main

import (
	"sort"
	"strings"
)

/*
Создание и сортировка словаря с длиной слов:
Напишите функцию, которая принимает строку, разбивает её на слова, создаёт map, где ключ — слово,
а значение — его длина, затем сортирует map по длине слов и возвращает список ключей.
*/

import (
	"fmt"
)

func SortingDictionary(words string) []string {
	m := make(map[string]int)
	keys := make([]string, 0)

	for _, word := range strings.Fields(words) {
		if _, ok := m[word]; !ok {
			m[word] = len([]rune(word))
			keys = append(keys, word)
		}
	}

	sort.Slice(keys, func(i, j int) bool { return m[keys[i]] < m[keys[j]] })

	return keys
}

func main() {
	str := "AAAA KKKKKKKKKKKKKKKKKK BB CCC AAAA FFFFFFFFFFFFFFFFF DDDDD EEEEEEEEEEE BB "
	fmt.Println(SortingDictionary(str))
}
