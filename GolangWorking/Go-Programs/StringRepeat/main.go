package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	valStr := "ab2d2c1"
	var result string

	for i := 0; i < len(valStr); i++ {
		eachCh := valStr[i]
		if i+1 < len(valStr) && unicode.IsDigit(rune(valStr[i+1])) {
			count, _ := strconv.Atoi(string(valStr[i+1]))
			result += strings.Repeat(string(eachCh), count)
			i++
		} else {
			result += string(eachCh)
		}
	}

	fmt.Println(result)

}
