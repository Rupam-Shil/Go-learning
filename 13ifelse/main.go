package main

import (
	"fmt"
)

func main() {
	fmt.Println("If else in golang")
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter login count")
	// value, _ := reader.ReadString('\n')
	// loginCount,_ := strconv.Atoi( strings.TrimSpace(value))
	// var result string
	// if loginCount < 10 {
	// 	result = "regular users"
	// } else if loginCount >10 {
	// 	result = "not regular users"
	// } else{
	// 	result = "miracle"
	// }
	// fmt.Println(result)

	if num :=3 ; num <10{
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is not less the 10")
	}
}