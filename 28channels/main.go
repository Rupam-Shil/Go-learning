package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channel in go lang")
	myCh := make(chan int,2)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	//Receieve only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		close(ch)
		val, isChannelOpen := <-myCh
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		fmt.Println(<-ch)
		wg.Done()
	}(myCh, wg)
	//send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 0
		close(ch)
		// ch <- 5
		// ch <- 6
		wg.Done()

	}(myCh, wg)
	wg.Wait()
	// myCh <- 5 
	// fmt.Println(<-myCh)

}