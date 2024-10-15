package main

import (
	"fmt"
	"strings"
)

/*
Реверс слов в предложении:
Напишите функцию, которая принимает строку, разбивает её на слова,
затем реверсирует каждое слово и возвращает строку,
где слова расположены в обратном порядке по сравнению с исходной строкой
*/

func reverse(str string) string {
	slcRune := []rune(str)
	resultStr := ""
	for i := len(slcRune) - 1; i >= 0; i-- {
		resultStr += string(slcRune[i])
	}

	return resultStr
}

func ReverseWords(strStart string) string {
	words := strings.Split(strStart, " ")
	resultStr := ""

	for i := len(words) - 1; i >= 0; i-- {
		resultStr += reverse(words[i])
		if i != 0 {
			resultStr += " "
		}
	}

	return resultStr
}

func main() {
	str := "The        golang code is written"
	fmt.Println(ReverseWords(str))
}
