package main

import (
	"fmt"
	"sort"
)

/*
Оптимизация слайса целых чисел:
Напишите функцию, которая принимает слайс целых чисел,
удаляет все элементы с нечетными индексами,
затем удваивает оставшиеся элементы и возвращает новый слайс, отсортированный по убыванию.
*/

func SliceOptimization(slc []int) []int {
	resultSlc := make([]int, 0, len(slc))
	for i, val := range slc {
		if i%2 == 0 {
			resultSlc = append(resultSlc, val*2)
		}
	}

	sort.Slice(resultSlc, func(i, j int) bool {
		return resultSlc[i] > resultSlc[j]
	})

	return resultSlc
}

func main() {
	slc := []int{1, 25, 30, 5, 60, 35}
	fmt.Println(SliceOptimization(slc))
}
