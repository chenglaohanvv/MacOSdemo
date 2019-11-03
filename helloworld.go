package main

import "fmt"

func main() {
	a := 3
	b := 2
	c := 5
	res5 := a > b && c%a == b
	fmt.Println(res5)

}
