package main

import (
	"fmt"
	"sort"
)

func Sort2DSlice(slc2D [][]int) []int {
	resultSlc := make([]int, len(slc2D))

	for i := range slc2D {
		max := slc2D[i][0]
		for _, val := range slc2D[i] {
			if val > max {
				max = val
			}
		}
		resultSlc[i] = max
	}

	sort.Slice(resultSlc, func(i, j int) bool { return resultSlc[i] < resultSlc[j] })

	return resultSlc

}

func main() {
	fmt.Println(Sort2DSlice([][]int{[]int{3, 4, 200}, []int{7, 1, 1}, []int{27, 3, 4}}))
}
