package main

import "fmt"

/*
Напишите функцию CommonCharacters(a, b string) []rune, которая принимает две строки и возвращает срез рун, которые встречаются в обеих строках.
Каждая руна должна появиться в результате только один раз, даже если она встречается несколько раз в обеих строках.
Пример: CommonCharacters("hello", "world") должно вернуть []rune{'o', 'l'}.
*/

func CommonCharacters(a, b string) []rune {
	var resSlc []rune
	m := make(map[rune]struct{})
	for _, valA := range a {
		for _, valB := range b {
			if valA == valB {
				m[valA] = struct{}{}
			}
		}
	}

	for k := range m {
		resSlc = append(resSlc, k)
		fmt.Println(string(k))
	}

	return resSlc
}

func main() {
	res := CommonCharacters("aaaaW", "world")
	fmt.Println(res)
}
