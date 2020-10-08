package main

import (
	"fmt"
	"strconv"
)

// Person is sample. Definition of a basic structure
type Person struct {
	Age        int
	FirstName  string
	SecondName string
	Sex        string
}

func main() {
	// ケース1
	hoge := Person{
		FirstName:  "Taro",
		SecondName: "Yamada",
		Sex:        "men",
		Age:        65,
	}
	fmt.Println("Name：" + hoge.FirstName + " " + hoge.SecondName)
	fmt.Println("Age:" + strconv.Itoa(hoge.Age))
	fmt.Println("Sex:" + hoge.Sex)
	// 型を調べる
	fmt.Printf("%T\n", hoge)
	fmt.Println(&hoge)

	fmt.Printf("\n")

	// ケース2
	hoge2 := new(Person)
	hoge2.FirstName = "Hanako"
	// hoge2.SecondName = "Tanaka" // 代入をしなくても初期値が入っている => ""
	// hoge2.Age = 65              // 代入をしなくても初期値が入っている => 0
	hoge2.Sex = "female"
	fmt.Println("Name：" + hoge2.FirstName + " " + hoge2.SecondName)
	fmt.Println("Age:" + strconv.Itoa(hoge2.Age))
	fmt.Println("Sex:" + hoge2.Sex)
	// 型を調べる
	fmt.Printf("%T\n", hoge2)
	fmt.Println(&hoge2)

	fmt.Printf("\n")

	// ケース3
	hogeGroup := make([]Person, 2)
	hogeGroup[0] = hoge
	hogeGroup[1] = *hoge2
	for _, v := range hogeGroup {
		fmt.Println("Name：" + v.FirstName + " " + v.SecondName)
		fmt.Println("Age:" + strconv.Itoa(v.Age))
		fmt.Println("Sex:" + v.Sex)
		// 型を調べる
		fmt.Printf("%T\n", v)
		fmt.Println(&v)
	}
}
