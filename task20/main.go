package main

import (
	"fmt"
	"strings"
)

/*
Реализация частотного анализа и фильтрация:
Напишите функцию, которая принимает строку, разбивает её на слова, создаёт частотный словарь,
затем удаляет все слова, которые встречаются только один раз,
и возвращает новый map с оставшимися словами.

*/

func FrequencyFiltering(words string) map[string]int {

	resultMap := make(map[string]int)

	for _, word := range strings.Fields(words) {
		if count := strings.Count(words, word); count > 1 {
			resultMap[word] = count
		}
	}

	return resultMap
}

func main() {
	str := "AA BBB  CCC CCC BBB d d D D E FF FF CCC BBB d d"
	fmt.Println(FrequencyFiltering(str))
}
