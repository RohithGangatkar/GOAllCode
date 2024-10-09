package main

import "fmt"

type StackList struct {
	Container []int
}

func (s *StackList) AddFront(add int) {
	s.Container = append(s.Container, add)
	fmt.Printf("Push : %v\n", s.Container)
}
func (s *StackList) Popup() {

	if len(s.Container) > 0 {
		fmt.Printf("popuedup : %d\n", s.Container[len(s.Container)-1])
		s.Container = s.Container[:len(s.Container)-1]
		return
	}
	fmt.Println("Stack is empty")

}

func (s *StackList) Empty() bool {
	return len(s.Container) == 0
}

func (s *StackList) FrontPeek() {
	if len(s.Container) > 0 {
		fmt.Printf("Front peek %d\n", s.Container[len(s.Container)-1])
	} else {
		fmt.Println("Stack is empty to FrontPeek")
	}

}

func (s *StackList) SearchStack(value int) int {

	if len(s.Container) > 0 {
		for i, eachChar := range s.Container {
			if eachChar == value {
				fmt.Println("Value exist in Queue for " + string(value))
				return i
			}
		}
	} else {
		fmt.Println("value exist in Queue as Queue is empty")
	}
	return -1
}

func (q *StackList) DeleteValue(value int) bool {

	index := q.SearchStack(value)
	if index != -1 {
		q.Container = append(q.Container[:index], q.Container[index+1:]...)
		return true
	}
	return false

}

func main() {

	stackList := StackList{
		Container: make([]int, 0),
	}
	stackList.AddFront(2)
	stackList.AddFront(4)
	stackList.AddFront(6)
	stackList.FrontPeek() // 6
	stackList.Popup()     // 6
	stackList.FrontPeek() //4
	stackList.AddFront(8) // [2,4,8]
	for len(stackList.Container) > 0 {
		stackList.Popup() // [8] [4] [2]
	}
	stackList.Popup()     //error stack empty
	stackList.AddFront(5) // [5]
	stackList.AddFront(3) // [5,3]
	stackList.SearchStack(5)
	stackList.FrontPeek() //3
	fmt.Println(stackList.DeleteValue(3))
	fmt.Printf("Stack : %v\n", stackList.Container)
	fmt.Printf("Is empty: %t ", stackList.Empty())

}
