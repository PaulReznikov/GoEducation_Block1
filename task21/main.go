package main

import "fmt"

/*
Модификация и фильтрация слайса:
Напишите функцию, которая принимает слайс целых чисел, умножает каждое число на его индекс,
затем оставляет только те числа, которые больше суммы всех чисел в оригинальном слайсе, и возвращает результат.
*/

func ModificationAndFiltering(slcStart []int) []int {
	resultSlc := make([]int, 0, len(slcStart))
	summ := 0

	for i, val := range slcStart {
		summ += val
		slcStart[i] = i * val
	}

	for _, val := range slcStart {
		if val > summ {
			resultSlc = append(resultSlc, val)
		}
	}

	return resultSlc
}

func main() {
	slc := []int{1, 2, 3, 4, 5, 6, 1}
	fmt.Println(ModificationAndFiltering(slc))
}
