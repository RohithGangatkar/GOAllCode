package main

import "fmt"

type Conatiner struct {
	Queue []string
}

func (q *Conatiner) PushBack(a string) {
	q.Queue = append(q.Queue, a)

}

func (q *Conatiner) RemoveFront() {
	if len(q.Queue) > 0 {
		q.Queue = q.Queue[1:]
		return
	} else {
		fmt.Println("Queue is empty")
	}

}

func (q *Conatiner) PeekFront() {
	if len(q.Queue) > 0 {
		fmt.Printf("Front element: %s \n", q.Queue[0])
	} else {
		fmt.Println("No peek front as queue is empty ")
	}

}

func (q *Conatiner) EmptyQueue() bool {

	// if s.IsEmpty() we an use to check empty
	return len(q.Queue) == 0
}

func (q *Conatiner) SearchQueue(value string) {

	if len(q.Queue) > 0 {
		for _, eachChar := range q.Queue {
			if eachChar == value {
				fmt.Println("Value exist in Queue for " + value)
				return
			}
		}
	} else {
		fmt.Println("value exist in Queue as Queue is empty")
	}
}

func main() {

	queue := Conatiner{
		Queue: make([]string, 0),
	}

	queue.PushBack("a")
	queue.PushBack("b")
	queue.PushBack("c")
	fmt.Println(queue.Queue)
	queue.PeekFront()
	fmt.Printf("Empty Queue: %t\n", queue.EmptyQueue())
	for len(queue.Queue) > 0 {
		queue.RemoveFront()
	}
	fmt.Printf("Queue Element: %s\n", queue.Queue)
	fmt.Printf("Empty Queue: %t\n", queue.EmptyQueue())
	queue.PushBack("e")
	queue.PushBack("f")
	queue.RemoveFront()
	fmt.Printf("Queue Element: %s\n", queue.Queue)
	queue.SearchQueue("f")

}
