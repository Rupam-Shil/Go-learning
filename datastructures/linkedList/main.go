package main

import (
	"fmt"
	"strings"
)

type node struct {
	value int
	next *node
}

type linkedList struct {
	head *node
	len int
}

func (n node)  String() string{
	return fmt.Sprintf("%d",n.value)
}

func (l *linkedList) add(value int){
	newNode := new(node)
	newNode.value = value
	// newNode.next = nil

	if l.head == nil {
		l.head = newNode
	} else {
		iterator := l.head
		for ; iterator.next != nil; iterator = iterator.next {
		}
		iterator.next = newNode
	}
}

func (l *linkedList) remove(value int){
	var previous *node
	for current := l.head; current != nil; current = current.next{
		if(current.value == value){
			previous.next = current.next
			return
		}

		previous = current
	}
}

func (l linkedList) get(value int) *node {
	for iterator := l.head;iterator != nil ; iterator = iterator.next {
		if(iterator.value == value){
			return iterator
		}
	}

	return nil
}

func (l linkedList) String() string {
	sb := strings.Builder{}

	for iterator := l.head;iterator != nil; iterator = iterator.next {
		sb.WriteString(fmt.Sprintf("%s",iterator))
	}

	return sb.String()
}

func main() {
	fmt.Println("Welcome to linked list")

	l := linkedList{}
	l.add(5)
	l.add(50)
	l.add(53)
	l.add(45)
	l.add(90)
	fmt.Println(l)
	fmt.Println(l.get(90))
}