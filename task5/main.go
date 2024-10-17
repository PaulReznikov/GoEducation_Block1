package main

import (
	"fmt"
	"strconv"
)

/*
Напишите функцию CompressString(s string) string, которая сжимает строку таким образом,
что последовательные повторяющиеся символы заменяются на символ, за которым следует количество повторений.
Если сжатая строка не короче исходной, вернуть исходную строку.
Пример: "aaabbcccc" должно вернуть "a3b2c4", но "abcd" должно вернуть "abcd".
*/

func CompressString(s string) string {
	str := ""
	counter := 1
	runes := []rune(s)

	fmt.Println(len(s))
	for i, val := range runes {

		if i == len(runes)-1 {
			str += string(val)
			if counter > 1 {
				str += strconv.Itoa(counter)
			}
			continue
		}

		if val == runes[i+1] { //??? for i, val := range str
			counter++
		} else {
			str += string(val)
			if counter > 1 {
				str += strconv.Itoa(counter)
				counter = 1
			}
		}

	}

	return str

}

func main() {
	res := CompressString("hééllo")
	fmt.Println(res)
}
