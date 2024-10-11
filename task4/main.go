package main

import "fmt"

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
