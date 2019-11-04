package main

import (
	"awesomeProject/src/Loops"
	"awesomeProject/src/Maths"
	"awesomeProject/src/Structs"
	"fmt"
	"time"
)

func main(){
	fmt.Println("Hello World")
	fmt.Println(time.Now())
	x := Maths.AddTwo(4,5)
	fmt.Println(x)
	fmt.Println(Maths.Swap("Lewis","Williams"))
	m,j := Maths.OneForTwoExport(5)
	fmt.Println(j)
	fmt.Println(m)

	var age  = 5
	fmt.Println(age)

	arr := Loops.PickLoop(35)
	fmt.Println(arr)
	fmt.Println("Here")
	fmt.Println(Loops.ArraySum([6]int{3,5,6,8,8,8}))
	person := Structs.PersonArray("Lewis", 26)
	fmt.Print(person.Name)
	result := Loops.RangeLoop(3,4,5,6,7)
	fmt.Println(*result)
}
