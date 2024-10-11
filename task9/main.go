package main

import "fmt"

func MergeHashTable(m1, m2 map[string]int) map[string]int {
	for k, v := range m1 {
		if _, ok := m2[k]; !ok {
			m2[k] = v
		} else {
			if sum := m1[k] + m2[k]; sum > 10 {
				m2[k] = sum
			} else {
				delete(m2, k)
			}
		}

	}

	return m2
}

func main() {
	m1 := map[string]int{"a": 7, "gGg": 6, "b": 2, "BHD": 3, "d": 6}
	m2 := map[string]int{"YYYYYYYY": 1, "a": 4, "b": 3, "d": 20, "QQ": 5}
	fmt.Println(MergeHashTable(m1, m2))
}
