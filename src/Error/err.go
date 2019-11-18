package main

import (
	"errors"
	"fmt"
)

func main() {

	res, err := add(4,0)
	if err != nil {
	fmt.Println(err)
	}
	fmt.Println(res)

}

func add(a,b int ) (int, error) {
	var addition int
	addition = a + b
	if addition == 0 {
	return 0, errors.New("addition can't be 0")
	}
	return addition, nil

}