package main

import (
	"fmt"
	"strings"
)

/*

Извлечение уникальных символов и их реверс:
Напишите функцию, которая принимает строку, извлекает все уникальные символы,
затем реверсирует строку и возвращает результат как новую строку.

*/

func ExtractionRevers(str string) string {
	resultStr := ""
	SlcRune := []rune(str)

	for i := len(SlcRune) - 1; i >= 0; i-- {
		if strings.Count(str, string(SlcRune[i])) == 1 {
			continue
		} else {
			resultStr += string(SlcRune[i])
		}
	}

	return resultStr
}

func main() {
	str := "IBCCCDD DFFFDE"
	fmt.Println(ExtractionRevers(str))
}
