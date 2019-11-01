package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


type info struct {
	Count int16
	Pages int8
	Next string
	Prev string
}
type origin struct {
	Name string
	Url string
}

type location struct {
	Name string
	Url string
}
type character struct {
	Id int8
	Name  string
	Status string
	Species string
	Type string
	gender string
	Origin origin
	Location location
	Image  string
	Episode  []string
	Url   string
	Created  string
}

type data struct {
	UserId int
	Id int
	Title string
	Completed bool
}

type Character struct {
	Info info
	Results []character
}


func main() {
	res, err := http.Get("https://rickandmortyapi.com/api/character/")
	if err != nil {
		log.Fatal(err)
	}
	char, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	var Data Character
	err = json.Unmarshal(char, &Data)
	if err != nil {
		fmt.Printf("%s Error:", err)
	}
	fmt.Printf("%+v", Data.Results[6])
}
//"https://jsonplaceholder.typicode.com/todos"
//"https://rickandmortyapi.com/api/character/2

