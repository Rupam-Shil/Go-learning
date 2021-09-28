package main

import (
	"fmt"
	"strings"
)

type node struct {
	value int
	next *node
}

type linkedList struct{
	head *node
	length int
}

func (l *linkedList) add(value int){
	newNode := new(node)
	newNode.value = value

	if l.head == nil {
		l.head = newNode
	} else {
		iterator:= l.head
		for ; iterator.next != nil; iterator = iterator.next{
		}
		iterator.next = newNode
	}
}

func (l linkedList) String() string {
	sb := strings.Builder{}
	for iterator := l.head; iterator != nil; iterator= iterator.next{
		sb.WriteString(fmt.Sprintf("%d ",iterator.value))
	}
	return sb.String()
}

func main() {
	
	fmt.Println("Welcome to linkedlist")
	l := linkedList{}
	l.add(5)
	l.add(90)
	fmt.Println(l)
}
