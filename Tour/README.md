# 1.go

## Packages

Every Go program is made up of packages.

Programs start running in package main.

By convention, the package name is the same as the last element of the import path.

```
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```

# 2.go

## Imports

This code groups the imports into a parenthesized, "factored" import statement.

You can also write multiple import statements, like:

import "fmt"
import "math"
But it is good style to use the factored import statement.
In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.

```
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}
```

# 3.go

A function can take zero or more arguments.

In this example, add takes two parameters of type int.

Notice that the type comes after the variable name.

```
package main

import "fmt"

func sum(a,b int) int{
	return a + b
}

func main() {
	fmt.Println(sum(9,4))
}
```

# 4.go

Go has only one looping construct, the for loop. Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.

```
package main

import "fmt"

func main(){
	oneName := "rupam"
	var anotherName string
	anotherName = "Subhadeep"

	for i:=0;i<10;i++{

		fmt.Printf("%v and %v \n",oneName,anotherName)
	}

}
```
