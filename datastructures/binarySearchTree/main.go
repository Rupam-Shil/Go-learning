package main

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	value int
	left *node
	right *node
}

type bst struct {
	root *node
	length int
}

func (n node) String() string {
	return strconv.Itoa(n.value)
}

func (b bst) String() string {
	sb := strings.Builder{}
	
}

func main() {
	fmt.Println("Welcome to binary search tree")
}