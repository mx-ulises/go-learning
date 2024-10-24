package main

import "fmt"

func PrintLine() {
	fmt.Println("------------------------------------------------------------")
}

func Excercise1() {
	greetings := []string{
		"Hello",
		"Hola",
		"你好",
		"こんにちは",
		"नमस्ते",
	}
	fmt.Println(greetings)
	slice1 := greetings[:2]
	slice2 := greetings[1:4]
	slice3 := greetings[3:]
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}

func Excercise2() {
	note := "Hi 👩 and 👨"
	noteRunes := []rune(note)
	fmt.Println(note)
	fmt.Println(noteRunes[3])
	fmt.Println(string(noteRunes[3]))
}

type Employee struct {
	fisrtName string
	lastName  string
	id        int
}

func Excercise3() {
	employee1 := Employee{"Ulises", "Martinez", 1}
	employee2 := Employee{fisrtName: "Rong", lastName: "Zhang", id: 2}
	var employee3 Employee
	employee3.fisrtName, employee3.lastName, employee3.id = "Dorian", "Echeverria", 3
	fmt.Println(employee1)
	fmt.Println(employee2)
	fmt.Println(employee3)
}

func main() {
	Excercise1()
	PrintLine()
	Excercise2()
	PrintLine()
	Excercise3()
}
