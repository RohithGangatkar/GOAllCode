package main

import (
	"fmt"
	"strings"
)

func getDistinct(arr []string) {
	result := make(map[string]bool)
	resultArr := []string{}

	for _, eachCh := range arr {
		if !result[eachCh] {
			result[eachCh] = true
			resultArr = append(resultArr, eachCh)
		}
	}
	fmt.Println(resultArr)

}

func getDistinctString(arr string) {
	result := make(map[rune]bool)
	var resultString string

	for _, eachCh := range arr {
		if !result[eachCh] {
			result[eachCh] = true
			resultString += string(eachCh)
		}
	}
	fmt.Println(resultString)

}

func getDistinctInt(arr []int) {
	result := make(map[int]bool)
	resultArr := []int{}

	for _, eachCh := range arr {
		if !result[eachCh] {
			result[eachCh] = true
			resultArr = append(resultArr, eachCh)
		}
	}
	fmt.Println(resultArr)

}

func CharCount(input []string) {
	resultMap := make(map[string]int)
	for _, eachInput := range input {
		// _, ok := resultMap[eachInput]
		// if ok {
		resultMap[eachInput]++
		// } else {
		// 	resultMap[eachInput] = 1
		// }
	}
	fmt.Println(resultMap)
}

func CharCountRune(input string) {
	resultMap := make(map[rune]int)

	for _, eachInput := range input {
		// _, ok := resultMap[eachInput]
		// if ok {
		resultMap[eachInput]++
		// } else {
		// 	resultMap[eachInput] = 1
		// }
	}
	fmt.Println(resultMap)
}

func convertAndCount(a, b []rune) {

	countMap := make(map[string]int)

	char1Str := string(a)
	char2Str := string(b)

	words1 := strings.Fields(char1Str)
	words2 := strings.Fields(char2Str)

	for _, eachWords := range words1 {
		countMap[eachWords]++

	}
	for _, eachWords := range words2 {
		countMap[eachWords]++

	}
	fmt.Println(countMap)

}

func main() {

	charArray := []string{"a", "d", "b", "f", "c", "a", "d"}
	getDistinct(charArray)
	input := "programming"
	getDistinctString(input)
	inputInt := []int{1, 2, 3, 2, 4, 1, 5, 6, 4}
	getDistinctInt(inputInt)
	CharCount(charArray)
	CharCountRune(input)
	charArray1 := []rune{'t', 'h', 'e', ' ', 's', 'u', 'n', ' ', 'r', 'i', 's', 'e', 's', ' ', 'i', 'n', ' ', 't', 'h', 'e', ' ', 'e', 'a', 's', 't'}
	charArray2 := []rune{'t', 'h', 'e', ' ', 's', 'u', 'n', ' ', 'g', 'i', 'v', 'e', 's', ' ', 'u', 's', ' ', 'w', 'a', 'r', 'm', 't', 'h'}
	convertAndCount(charArray1, charArray2)

}
