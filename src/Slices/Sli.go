package main

import "fmt"
type Person struct {
		name string
		age int8
}

func Slices(name string, age int8)  {
	var person []Person  // an empty list
	person = append(person, Person{name:name, age: age}, Person{name:name, age: age});
	person = append(person, Person{name:name, age: age});
	fmt.Println(person)
}

func main()  {
	Slices("L",39)
	sli := make([]int,0)
	sli = append(sli,4,7,7)
	fmt.Println(sli)
}