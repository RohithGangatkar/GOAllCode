package main

import (
	"fmt"
	"sync"
)

func checkChan(unbuf chan int) {
	unbuf <- 2
}

func sendChan(unbuf chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Starting sendChan")
	unbuf <- "Sent value to sendChan and recived in reciveChan"
	fmt.Println("Existing sendChan")

}

func reciveChan(unbuf chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Starting reciveChan")
	value := <-unbuf
	fmt.Println(value)
	fmt.Println("Ending reciveChan")
}

func consume(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func main() {

	// var wg sync.WaitGroup

	// chaUnBuff := make(chan int)
	// go checkChan(chaUnBuff)
	// value := <-chaUnBuff
	// fmt.Println(value)

	// wg.Add(2)
	// chaUnBuff2 := make(chan string)
	// go sendChan(chaUnBuff2, &wg)
	// go reciveChan(chaUnBuff2, &wg)

	// wg.Wait()

	// bufferCh := make(chan int, 2)
	// bufferCh <- 1
	// bufferCh <- 2
	// bufferCh <- 3 // deadlock
	// fmt.Println(<-bufferCh)
	// //or and
	// fmt.Println(<-bufferCh)

	// conprod := make(chan int, 2)
	// go consume(conprod)
	// for ech := range conprod {
	// 	fmt.Println(ech)
	// }

	bufferCh := make(chan int, 2)
	bufferCh <- 1
	bufferCh <- 2
	close(bufferCh)
	n, c := <-bufferCh //close
	fmt.Printf("Received: %d, open: %t\n", n, c)

	fmt.Println(<-bufferCh)
	fmt.Println(<-bufferCh)

}
