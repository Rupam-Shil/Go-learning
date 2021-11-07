package main

import (
	"fmt"
	"sync"
)


func main() {
	wg := &sync.WaitGroup{}
	mt := &sync.Mutex{}
	score := []int{}
	wg.Add(3)

	go includeScore(&score, 5,wg,mt)
	go includeScore(&score, 20,wg,mt)
	go includeScore(&score, 50,wg,mt)
	wg.Wait()
	fmt.Println(score)
}

func includeScore(score *[]int, val int, wg *sync.WaitGroup, mt *sync.Mutex) {
	for i := 0; i <10 ; i++ {
		mt.Lock()
		*score = append(*score, val+i)
		mt.Unlock()
	}
	wg.Done()
}