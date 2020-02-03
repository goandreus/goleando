package main

import "fmt"

func main(){
	var a  int
	var b int8

	a = 11111
	b = 5
	n := "andres"
	p := "chavez"

	//casting
	c := a + int(b)
	fmt.Println(c)
	fmt.Printf("Hola %s %s %d como estas\n", n , p ,b)
	fmt.Printf("c es de tipo :%T\n", c )
	fmt.Printf("b es de tipo :%T\n", b)
	// prioridad aritmetica
	//d := 6 + 6*(6 -6)

	//fmt.Println(d)

	//valor cero
	var nombre  string
	var numero int
	var entidad bool
	fmt.Println(nombre, numero, entidad)

}