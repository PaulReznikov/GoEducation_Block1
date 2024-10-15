package main

import (
	"fmt"
	"sort"
	"unicode"
)

/*
Фильтрация, сортировка и поиск строки:
Напишите функцию, которая принимает слайс строк, фильтрует строки, содержащие хотя бы одну цифру,
затем сортирует оставшиеся строки в порядке убывания длины и возвращает первую строку с минимальной длиной больше заданного значения.
*/

func remove(words []string, idx int) []string {
	return append(words[:idx], words[idx+1:]...)
}

func hasDigit(word string) bool {
	for _, r := range word {
		if unicode.IsDigit(r) {
			return true
		}
	}

	return false
}

func FilteringSortSearch(words []string, searchLen int) string {
	for i := len(words) - 1; i >= 0; i-- {
		if hasDigit(words[i]) {
			words = remove(words, i)
		}
	}

	sort.Slice(words, func(i, j int) bool {
		return len([]rune(words[i])) > len([]rune(words[j]))
	})

	for i := len(words) - 1; i >= 0; i-- {
		if len(words[i]) > searchLen {
			return words[i]
		}
	}

	return ""
}

func main() {
	slcStr := []string{"hhhh4", "a", "bbb", "c4", "cccc", "eeeeee", "fffffffffff"}
	fmt.Println(FilteringSortSearch(slcStr, 5))
}
