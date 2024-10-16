package main

import (
	"fmt"
	"strings"
)

/*
Анализ строк и частотный словарь:
Напишите функцию, которая принимает строку, разбивает её на слова, удаляет все слова короче 3 символов,
затем возвращает map, где ключом является слово, а значением — количество его появлений в строке.
*/

func remove(strSlc []string, idx int) []string {
	return append(strSlc[:idx], strSlc[idx+1:]...)
}

func PrivateDictionary(s string) map[string]int {
	m := make(map[string]int)

	strSlc := strings.Fields(s)
	for i := 0; i < len(strSlc); i++ {
		if len([]rune(strSlc[i])) < 3 {
			strSlc = remove(strSlc, i)
			i--
		}
	}

	for _, val := range strSlc {
		if _, ok := m[val]; !ok {
			m[val] = 1
		} else {
			m[val]++
		}
	}

	return m
}

func main() {
	fmt.Println(PrivateDictionary("aaa bbbb ccc dd  a eeeeeeeeee fff dddd l ccc ccc aaa"))
}
