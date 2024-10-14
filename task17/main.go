package main

import "fmt"

/*

Алгоритм с накоплением результата:
Напишите функцию, которая принимает слайс целых чисел, создаёт новый слайс,
где каждый элемент является суммой всех предыдущих элементов оригинального слайса,
затем удаляет все элементы меньше среднего арифметического этого слайса.

*/

func AccumulationOfResults(slcStart []int) []int {
	resultSlc := make([]int, 0, len(slcStart))

	summ, median := 0, 0

	for i, val := range slcStart {

		if i == 0 || i == 1 {
			summ += val
			median += val
			continue
		}

		slcStart[i] = summ
		median += summ
		summ += summ

	}

	for _, val := range slcStart {
		if val >= median/len(slcStart) {
			resultSlc = append(resultSlc, val)
		}
	}

	return resultSlc
}

func main() {
	slc := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(AccumulationOfResults(slc))

}
