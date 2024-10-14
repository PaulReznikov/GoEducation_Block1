package main

import (
	"fmt"
	"strings"
)

/*
Удаление и преобразование элементов в слайсе:
Напишите функцию, которая принимает слайс строк, удаляет все строки, содержащие менее 5 символов,
затем преобразует оставшиеся строки в верхний регистр и возвращает их в обратном порядке.
*/

func DeletingAndConvertingElements(slcStr []string) []string {
	var resultSlc []string

	for i := len(slcStr) - 1; i >= 0; i-- {
		if len([]rune(slcStr[i])) < 5 {
			continue
		} else {
			resultSlc = append(resultSlc, strings.ToUpper(slcStr[i]))
		}
	}

	return resultSlc
}

func main() {
	slcStr := []string{"aaa", "bbbbbbbb", "c", "ddddd", "eeeeeeeeeeeeeeeee", "ffff"}
	fmt.Println(DeletingAndConvertingElements(slcStr))
}
