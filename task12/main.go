package main

import (
	"fmt"
	"regexp"
	"strconv"
)

/*
Извлечение и подсчёт символов:
Напишите функцию, которая принимает строку, извлекает все цифры из строки,
затем преобразует их в целые числа, и возвращает их сумму.
*/

func ExtractingCoutingChar(str string) int {

	summ := 0

	for _, v := range str {
		if flag, _ := regexp.MatchString(`^[0-9]$`, string(v)); flag {

			num, err := strconv.Atoi(string(v))
			if err != nil {
				panic("something wrong")
			}

			summ += num
		}

	}

	return summ

}

func main() {
	str := "5sad5nsd2JHsdk2"

	fmt.Println(ExtractingCoutingChar(str))
}
