package main

import "fmt"


func main() {
	var a int = 42
	//b is pointing to a int and that int is a ref of a
	var b *int = &a
	*b = 34

	fmt.Println(a,*b)

	var arr [2]int
	arr[0] = 23
	arr[1] = 233

	var m *int = &arr[0]
	fmt.Println(arr[0], *m)
	*m = 45
	fmt.Println(arr[0], *m)

	var ms *MyStruct
	stru := &MyStruct{foo: 45}
	ms = stru
	fmt.Println(*ms)

	name := "Luke";
	greeting := "Hello"
	sayHello(&name,&greeting)
	fmt.Println(name)
}
// I am a pointer to a type and that type has the address.

type MyStruct struct {
	foo int
}

func sayHello(name, greeting *string) {
	fmt.Println(*name,*greeting)
	*name = "Yoda"
	fmt.Println(*name)
}
