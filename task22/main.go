package main

/*
Создание индексации слов:
Напишите функцию, которая принимает строку, разбивает её на слова, создаёт map, где ключом является слово,
а значением — список индексов, где это слово встречается в строке.
*/

import (
	"fmt"
	"strings"
)

func WordsIndexing(strStart string) map[string][]int {
	resultMap := make(map[string][]int)
	words := strings.Fields(strStart)

	for i, val := range words {
		resultMap[val] = append(resultMap[val], i)
	}

	return resultMap
}

func main() {
	str := " I word I bbb I word cc XX XX"
	fmt.Println(WordsIndexing(str))
}
