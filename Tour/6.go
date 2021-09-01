package main

import (
	"fmt"
	"math"
	"runtime"
)

func getPower(num, pow float64) float64{
	return math.Pow(num, pow)
}

func main() {
	fmt.Println(runtime.GOOS)
}