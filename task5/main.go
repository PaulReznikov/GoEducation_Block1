package main

import (
	"fmt"
	"strconv"
)

func CompressString(s string) string {
	str := ""
	counter := 1

	for i, val := range s {

		if i == len([]rune(s))-1 {
			str += string(val)
			if counter > 1 {
				str += strconv.Itoa(counter)
			}
			continue
		}

		if val == []rune(s)[i+1] {
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
	res := CompressString("aaabbcccca")
	fmt.Println(res)
}
