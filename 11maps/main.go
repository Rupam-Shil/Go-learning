package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"

	for short, name := range languages {
		fmt.Println(short, name)
	}

	delete(languages, "rb")
	fmt.Println(languages)
}