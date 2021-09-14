package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	name := "Welcome homies"

	fmt.Println(len(name))
	fmt.Println(name[:6])

	name += "!"
	fmt.Println(name)

	abc := "abc"
	b := []byte(abc)
	fmt.Printf("%s %s",abc, b)

	//string library
	fmt.Println(strings.ToUpper("rupam"))
	fmt.Println(strings.Count("sdfsdfsdfsdfsdfdsf","sdfsdf"))
	fmt.Println(strings.Replace("there is a dog","dog","cat",1))

	var sb strings.Builder;
	sb.WriteString("Hello")
	fmt.Println(sb.String())
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())

	x := 12345678

	y := strconv.Itoa(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}