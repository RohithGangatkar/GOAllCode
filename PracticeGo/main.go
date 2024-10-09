package main

import (
	anonymous "PracticeGo/Anonymous"
	anonymousFilter "PracticeGo/AnonymousFilter"
	"fmt"
)

func main() {

	//assigns a function to a variable.
	a := func(name string) {
		fmt.Printf("user name: %s \n", name)
	}
	fmt.Printf("%T\n", a)
	a("jon")

	//User defined function types
	anonymous.AnonymousUserDefined()

	//Passing functions as arguments to other functions
	f := func(st1, str2 string) string {
		return fmt.Sprintln(st1 + str2)
	}
	anonymous.AnonymousArgument(f)

	//Returning functions from other functions
	f2 := anonymous.AnonymousReturnFunc()
	fmt.Printf("Result of %d*%d = %d \n", 2, 3, f2(2, 3))

	//function closure
	anonymous.AnonymousClosure("All")
	anonymous.AnonymousClosure("Everyone")

	//function closure return function
	funA := anonymous.ClosureReturn()
	funB := anonymous.ClosureReturn()
	fmt.Printf("user greeting: %s \n", funA("all"))
	fmt.Printf("user greeting: %s \n", funA("!!Good"))
	fmt.Printf("user greeting: %s \n", funB("EveryOne"))
	fmt.Printf("user greeting: %s \n", funB("!ok"))

	//filter
	s1 := anonymousFilter.Student{
		FirstName: "Naveen",
		LastName:  "Ramanathan",
		Grade:     "A",
		Country:   "India",
	}
	s2 := anonymousFilter.Student{
		FirstName: "Samuel",
		LastName:  "Johnson",
		Grade:     "B",
		Country:   "USA",
	}
	s := []anonymousFilter.Student{s1, s2}

	SFilter := anonymousFilter.Filter(s, func(s anonymousFilter.Student) bool {
		return s.Country == "India"

	})
	fmt.Printf("Filtered result : %+v \n", SFilter)

}
