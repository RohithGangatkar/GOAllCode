package anonymous

import "fmt"

type add func(a, b int) int

var a add = func(a, b int) int {
	return a + b
}

func AnonymousUserDefined() {

	fmt.Printf("Sum of number :%d \n", a(2, 4))

}

func AnonymousArgument(f func(a, b string) string) {

	fmt.Println(f("jon ", "paul"))

}

func AnonymousReturnFunc() func(a, b int) int {

	f := func(a, b int) int {
		return a * b
	}

	return f

}

func AnonymousClosure(t string) {
	var str string

	str = func(a string) string {
		str = "Helo " + a
		return str
	}(t)

	fmt.Println(str)

}

func ClosureReturn() func(s string) string {
	var str string

	fun := func(b string) string {
		str = str + "Heloo " + b + " "
		return str
	}

	return fun

}
