package Structs

type Person struct {
	Name string
	Age uint8
}

// TODO make a array of people and pass that in.
func PersonArray(name string, age uint8) Person {
	p := Person{name,age}
	return p
}