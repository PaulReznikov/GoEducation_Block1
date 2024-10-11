package main

import (
	"fmt"
	"sort"
)

//Пересечение и объединение слайсов:
//Напишите функцию, которая принимает два слайса строк, находит их пересечение, затем объединяет уникальные элементы из обоих слайсов и возвращает новый слайс, отсортированный по алфавиту.

func MergeSlicesStrings(strSlc1, strSlc2 []string) []string {
	m := make(map[string]struct{})
	var resSlc []string

	for _, val1 := range strSlc1 {
		m[val1] = struct{}{}
	}

	for _, val2 := range strSlc2 {
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
	slc1 := []string{"aaa", "B", "advR", "bb", "AdvR"}
	slc2 := []string{"XX", "aaa", "cc", "bb"}
	fmt.Println(MergeSlicesStrings(slc1, slc2))
}
