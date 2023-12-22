package main

import "fmt"

func Excercise1() {
	fmt.Println(`Excercise 1: Write a program where you declare an integer varaible called i with the value 20.
Assign i to a floating point variable named f. Print out i and f.`)
	var i int = 20
	var f float32 = float32(i)
	//var f float32 = 20.0
	fmt.Printf("%d\n", i)
	fmt.Printf("%f\n", f)
}

func main() {
	Excercise1()
}
