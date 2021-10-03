package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const link string = "https://www.amazon.in/b?node=27014610031&pf_rd_r=6MMEC7PAP6VXHZVQRR1W&pf_rd_p=43e613ef-6e87-4b32-b1db-b9eedc5224b0&pd_rd_r=1bc44e11-538e-4c26-acdc-e62d6e1c9780&pd_rd_w=uhPzq&pd_rd_wg=C67rP&ref_=pd_gw_unk"


func main() {
	fmt.Println("Lets do some practice")
	// name := getName()
	// fmt.Println("Hello",name)
	getWebsite(link)

}

func checkError(err error) {
	if err != nil{
		panic(err)
	}
}

func getName() string{
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name")
	name, err := reader.ReadString('\n')
	checkError(err)
	name = strings.TrimSpace(name)
	return name
}

func getWebsite(s string){
	res , err := http.Get(s)
	checkError(err)
	fmt.Printf("The type of website is %T\n", res)
	fmt.Printf("%v\n",*res)
	html, err := ioutil.ReadAll(res.Body)
	checkError(err)
	fmt.Println(string(html))
	defer res.Body.Close()
}