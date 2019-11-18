package main

import "fmt"

func main() {
channel := make(chan int,2)

channel <- 4

fmt.Println(<- channel)

}
//WHEN THE CHANNEL IS ON THE LEFT HAND SIDE OF THE <- WE ARE PASSING DATA INTO THAT CHANNEL
//WHEN THE VALUE IS ON THE LEFT HAND SIDE OF THE <- WE ARE READING THAT VALUE FROM THE CHANNEL
