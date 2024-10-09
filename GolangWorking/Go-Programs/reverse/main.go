package main

import "fmt"

func reverseInt(values []int) []int {

	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}
	return values
}

func reverseString(values string) string {

	reverse := []rune(values)
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		reverse[i], reverse[j] = reverse[j], reverse[i]
	}

	return string(reverse)
}

func reverseStringBasic(values string) (result string) {
	for _, eachValue := range values {
		result = string(eachValue) + result
	}
	return result
}

func main() {

	values := []int{10, 12, 14, 16, 18, 20}
	fmt.Println(reverseInt(values))
	values2 := []int{10, 12, 14, 16, 18}
	fmt.Println(reverseInt(values2))
	values3 := "slices"
	fmt.Println(reverseString(values3))
	values4 := "slices"
	fmt.Println(reverseStringBasic(values4))

}
