package main

import (
	"fmt"
	"strings"
)

func main() {

	// name := "No 'x' in Nixon"
	// fmt.Println(fibonacciNumber(8))
	// fmt.Println(palindrom(name))
	// fmt.Println(fibo(6))
	s := []int{1, 2, 3, 5, 6, 8}
	fmt.Println(reverse(s))

}
func fibo(n int) int {
	fib := make([]int, n+1)
	fib[0] = 0
	fib[1] = 1

	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib[n]
}

func fibonacciNumber(num int) int {

	if num == 1 {
		return 0
	}
	if num == 1 {
		return 1
	}
	return fibonacciNumber(num-2) + fibonacciNumber(num-1)
}

func palindrom(name string) bool {

	var cName string
	name = strings.ToLower(name)

	for _, c := range name {
		if (c > 'a' && c < 'z') || (c > 0 && c < 9) {
			cName += string(c)
		}
	}

	lName := strings.ToLower(name)

	left := 0
	right := len(lName) - 1

	if lName[left] != lName[right] {
		return false
	}
	right--
	left++

	return true

}

func reverse(value []int) []int {

	// newValue := make([]int, len(value))

	for i, j := 0, len(value)-1; i < j; i, j = i+1, j-1 {
		value[i], value[j] = value[j], value[i]

	}
	return value

}
