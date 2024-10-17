package main

import "fmt"

/*
Фильтрация, преобразование и суммирование:
Напишите функцию, которая принимает слайс целых чисел,
удаляет из него все отрицательные числа,
затем возводит каждое оставшееся число в квадрат и возвращает сумму всех полученных квадратов.
*/

func remove(slc []int, idx int) []int {
	return append(slc[:idx], slc[idx+1:]...)
}

func FilteringAndSumm(slc []int) int {
	//summ := 0
	//
	//for i := 0; i < len(slc); i++ {
	//	if slc[i] < 0 {
	//		slc = remove(slc, i)
	//		i--
	//		continue
	//	}
	//
	//	summ += slc[i] * slc[i]
	//}
	//
	//return summ

	summ := 0
	for _, num := range slc {
		if num >= 0 {
			summ += num * num
		}
	}
	return summ

}

func main() {
	slc := []int{1, -5, -3, 2, -10, 3, -5}
	fmt.Println(FilteringAndSumm(slc))
}
