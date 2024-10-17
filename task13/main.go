package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Пересечение и объединение слайсов:
Напишите функцию, которая принимает два слайса строк, находит их пересечение,
затем объединяет уникальные элементы из обоих слайсов и возвращает новый слайс, отсортированный по алфавиту.
*/

func MergeSlicesStrings(strSlc1, strSlc2 []string) []string {
	m := make(map[string]struct{})
	var resSlc []string

	for _, val1 := range strSlc1 {
		val1 = strings.ToLower(val1)
		m[strings.ToLower(val1)] = struct{}{}
	}

	for _, val2 := range strSlc2 {
		val2 = strings.ToLower(val2)
		if _, ok := m[val2]; !ok {
			m[val2] = struct{}{}
		} else {
			delete(m, val2)
		}
	}

	for k := range m {
		resSlc = append(resSlc, k)
	}

	sort.Strings(resSlc)

	return resSlc
}

func main() {
	slc1 := []string{"x", "y"}
	slc2 := []string{"x", "y"}
	fmt.Println(MergeSlicesStrings(slc1, slc2))
}
