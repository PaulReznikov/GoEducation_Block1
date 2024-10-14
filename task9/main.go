package main

import "fmt"

/*
Слияние map с условием:
Напишите функцию, которая принимает два map (оба с типом ключей — строка, значений — целое число),
объединяет их в один map, складывая значения одинаковых ключей, но только если сумма больше 10.
Если меньше, ключ должен быть удалён.
*/

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
