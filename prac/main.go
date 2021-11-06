package main

import (
	"fmt"
	"sync"
)

func main() {
	greeting("Start")

	count := []int{0}

	wg := &sync.WaitGroup{}

	mt :=  &sync.Mutex{}

	wg.Add(4)

	go func (wg *sync.WaitGroup, mt *sync.Mutex) {
		fmt.Println("In One")
		mt.Lock()
		count  = append(count,1)
		mt.Unlock()
		wg.Done()
	}(wg, mt)

	go func (wg *sync.WaitGroup, mt *sync.Mutex)  {
		fmt.Println("In Two")
		mt.Lock()
		count  = append(count,2)
		mt.Unlock()
		wg.Done()
	}(wg, mt)

	go func (wg *sync.WaitGroup, mt *sync.Mutex)  {
		fmt.Println("In Three")
		mt.Lock()
		count  = append(count,3)
		mt.Unlock()
		wg.Done()
	}(wg, mt)
	
	go func (wg *sync.WaitGroup, mt *sync.Mutex)  {
		fmt.Println("In Four")
		mt.Lock()
		count  = append(count,4)
		mt.Unlock()
		wg.Done()
	}(wg, mt)
	
	wg.Wait()
	fmt.Println(count)
	fmt.Println("Finished")
}

func greeting(s string){
	fmt.Println(s)
}