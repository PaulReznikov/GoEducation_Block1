package main

import (
	"fmt"
	"strings"
)

/*
Нормализация строк и подсчёт слов:
Напишите функцию, которая принимает строку, удаляет из неё все знаки препинания,
переводит все буквы в нижний регистр, затем разбивает строку на слова и возвращает map,
где ключ — слово, а значение — количество его появлений.

*/

func CountingWords(words string) map[string]int {
	resultMap := make(map[string]int)
	for _, word := range strings.Fields(words) {
		word = strings.ToLower(strings.Trim(word, ".,:;!?"))
		if _, ok := resultMap[word]; !ok {
			resultMap[word] = 1
		} else {
			resultMap[word]++
		}
	}

	return resultMap
}

func main() {
	str := "...,?!;;"
	fmt.Println(CountingWords(str))
	//fmt.Println(strings.Trim(str, ".,:;!?"))

}
