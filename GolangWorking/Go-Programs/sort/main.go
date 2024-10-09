package main

import (
	"fmt"
	"sort"
)

// aLower := strings.ToLower(a) // Convert to lowercase

func BubbleSort(values []int) []int {

	for i := 0; i < len(values)-1; i++ {
		for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
			}
		}

	}
	return values
}

func BubbleSortString(values string) string {

	runeValues := []rune(values)
	for i := 0; i < len(values)-1; i++ {
		for j := 0; j < len(values)-i-1; j++ {
			if runeValues[j] > runeValues[j+1] {
				runeValues[j], runeValues[j+1] = runeValues[j+1], runeValues[j]
			}
		}

	}
	return string(runeValues)
}

func BuiltInSort(values []int) {
	sort.Ints(values)
	fmt.Printf("built in sort: %v \n", values)

}

func BuiltInSortString(values string) {
	runes := []rune(values)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	fmt.Printf("built in sortString: %s \n", string(runes))

}

func BuiltInSortArrayString(charArray []string) {
	sort.Strings(charArray)

	// Print the sorted slice
	fmt.Printf("Sorted string array:%v \n", charArray)

}

func BuiltInSortDescString(values []string) {

	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})
	fmt.Printf("built in sortString: %v \n", values)

}

func SortMix(values []string) {

	for i := 0; i < len(values)-1; i++ {
		for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
			}
		}

	}
	fmt.Println(values)
}

func main() {

	//BubbelSort
	intValues := []int{10, 18, 15, 20, 16}
	stringValues := "abdegjvbtj"
	charArray := []string{"a", "d", "b", "f", "c"}
	fmt.Println(BubbleSort(intValues))
	BuiltInSort(intValues)
	fmt.Println(BubbleSortString(stringValues))
	BuiltInSortString(stringValues)
	BuiltInSortArrayString(charArray)
	BuiltInSortDescString(charArray)
	charMixArray := []string{"a", "d", "b", "f", "c", "A", "B", "C"}
	SortMix(charMixArray)

}
