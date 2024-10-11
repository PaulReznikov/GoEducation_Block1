package main

import "fmt"

func remove(slc []int, idx int) []int {
	return append(slc[:idx], slc[idx+1:]...)
}

func FilteringAndSumm(slc []int) int {
	summ := 0

	for i := 0; i < len(slc); i++ {
		if slc[i] < 0 {
			slc = remove(slc, i)
			i--
			continue
		}

		summ += slc[i] * slc[i]
	}

	return summ

}

func main() {
	slc := []int{1, -5, -3, 2, -10, 3, -5}
	fmt.Println(FilteringAndSumm(slc))
}
