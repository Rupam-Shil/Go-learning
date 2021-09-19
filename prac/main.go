package main

import "fmt"

func main() {
	rupam := student{"Rupam", "ENG19CS0265", 5}
	fmt.Println(rupam)
	rupam.printUsn()
	// fmt.Println("Please enter your name")
	// reader  := bufio.NewReader(os.Stdin)
	// name, _ := reader.ReadString('\n')
	// fmt.Println("Please enter your Gender(m/f)")
	// gender, _ := reader.ReadString('\n')
	// gender = strings.TrimSpace(gender)
	// if gender == "m" {
	// 	fmt.Println("Welcome Mr.", name)
	// }else{
	// 	fmt.Println("Welcome Ms.", name)
	// }

}

type student struct {
	name string
	usn  string
	year int
}

func (s student) printUsn() {
	fmt.Println("The usn is",s.name)
}