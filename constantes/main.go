package main

import "fmt"

func main(){
	var a  int
	var b int8

	a = 11111
	b = 5


	//casting
	c := a + int(b)
	fmt.Println(c)

}