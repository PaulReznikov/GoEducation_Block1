package main

import (
	"fmt"
	"sort"
)

/*

Работа с вложенными слайсами:
Напишите функцию, которая принимает слайс слайсов целых чисел.
Функция должна найти максимальное значение в каждом вложенном слайсе,
затем вернуть слайс, содержащий эти максимальные значения, отсортированный в порядке возрастания.
*/

func Sort2DSlice(slc2D [][]int) []int {
	resultSlc := make([]int, 0, len(slc2D))

	for i := range slc2D {
		if len(slc2D[i]) == 0 {
			continue
		}
		max := slc2D[i][0]
		for _, val := range slc2D[i] {
			if val > max {
				max = val
			}
		}
		resultSlc = append(resultSlc, max)
	}

	sort.Ints(resultSlc)

	return resultSlc

}

func main() {
	fmt.Println(Sort2DSlice([][]int{[]int{3, 4, 200}, []int{7, 1, 1}, []int{27, 3, 4}}))
}
