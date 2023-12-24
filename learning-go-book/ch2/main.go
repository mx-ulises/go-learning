package main

import "fmt"

func PrintLine() {
	fmt.Println("------------------------------------------------------------")
}

func Excercise1() {
	fmt.Println(`Excercise 1: Write a program where you declare an integer varaible called i with the value 20.
Assign i to a floating point variable named f. Print out i and f.`)
	var i int = 20
	var f float32 = float32(i)
	//var f float32 = 20.0
	fmt.Printf("%d\n", i)
	fmt.Printf("%f\n", f)
}

func Excercise2() {
	fmt.Println(`Excercise 2: Write a pogram where you declare a constant called value that can be assigned to both an integer and a floating point variable.
Assign it to an integer called i and a floating point variable called f. Print out i and f.`)
	const value = 10
	var i int = value
	var f float32 = value
	fmt.Printf("%d\n", i)
	fmt.Printf("%f\n", f)
}

func Excercise3() {
	fmt.Println(`Excercise 3: Write a program with three variables:
 - one named b of type byte
 - one named smallI of type int32
 - one named bigI of type uint64
Assign each variable the maximum legal value for its type, then Add 1 to each variable. Print out their values.`)
	const value = 10
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615
	fmt.Printf("byte: %d -> %d\n", b, b+1)
	fmt.Printf("int32: %d -> %d\n", smallI, smallI+1)
	fmt.Printf("uint64: %d -> %d\n", bigI, bigI+1)
}

func main() {
	Excercise1()
	PrintLine()
	Excercise2()
	PrintLine()
	Excercise3()
}
