package main

import (
	"fmt"
)

// Person is sample. Definition of a basic structure
type Person struct {
	Age        int
	FirstName  string
	SecondName string
	Sex        string
}

func main() {
	hoge := Person{
		Age:        65,
		FirstName:  "Taro",
		SecondName: "Yamada",
		Sex:        "men",
	}
	fmt.Println(hoge)
}
