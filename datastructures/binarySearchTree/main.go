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
	b.inOrderTraversal(&sb)
	return sb.String()
	
}

func (b bst) inOrderTraversal(sb *strings.Builder) {
	b.inOrderTraversalByNode(sb, b.root)
}

func (b bst) inOrderTraversalByNode(sb *strings.Builder, root *node) {
	if root == nil {
		return
	}
}

func main() {
	fmt.Println("Welcome to binary search tree")
}