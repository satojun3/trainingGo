package main

import (
	"flag"
	"fmt"
)

func main() {
	// flag.Parse()
	// args := flag.Args()
	// fmt.Println(args)

	// var ip *int = flag.Int("int", 1234, "help message for flagname")
	// flag.Parse()     // flag定義をした後にParseを実行する
	// fmt.Println(ip)  // pointer型を表示
	// fmt.Println(*ip) // pointerに格納されている値を表示

	var mode *string = flag.String("mode", "typing", "This is a written test of English words.")
	flag.Parse()
	if *mode == "typing" {
		fmt.Printf("This is a written test of English words.\n")
	} else if *mode == "optional" {
		fmt.Printf("This is a selective test.\n")
	}
}
