package main

import "fmt"

func AddMap(name string, i int) map[int]string  {

	var lo = map[int]int{5:7,6:8}

	fmt.Print(lo)

	m := make(map[int]string)

	m[i] = name
	m[7] = "name"
	return m
}
