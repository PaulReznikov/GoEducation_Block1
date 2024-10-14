package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Группировка и сортировка слов по длине:
Напишите функцию, которая принимает строку, разбивает её на слова,
группирует слова по их длине в map (ключ — длина слова, значение — список слов),
а затем возвращает слайс этих групп, отсортированных по длине.

*/

func GrouppingSort(str string) [][]string {
	mLenWords := make(map[int][]string)
	slcWords := strings.Fields(str)
	slcKey := make([]int, 0)

	for _, val := range slcWords {
		if _, ok := mLenWords[len([]rune(val))]; !ok {
			slcKey = append(slcKey, len([]rune(val)))
		}
		mLenWords[len([]rune(val))] = append(mLenWords[len([]rune(val))], val)
	}

	sort.Slice(slcKey, func(i, j int) bool { return slcKey[i] < slcKey[j] })

	resSlc := make([][]string, len(mLenWords))

	for i, val := range slcKey {
		resSlc[i] = mLenWords[val]
	}

	return resSlc
}

func main() {
	str := "    A bb ccc bb dddd ccc ccc eeeee "

	fmt.Println(GrouppingSort(str))
}
