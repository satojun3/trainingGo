package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// flag.Parse()
	// args := flag.Args()
	// fmt.Println(args)

	// var ip *int = flag.Int("int", 1234, "help message for flagname")
	// flag.Parse()     // flag定義をした後にParseを実行する
	// fmt.Println(ip)  // pointer型を表示
	// fmt.Println(*ip) // pointerに格納されている値を表示

	var mode *string = flag.String("mode", "w", "This is a written test of English words.")
	flag.Parse()
	if *mode == "w" {
		fmt.Printf("This is a written test of English words.\n")

		japaneseWord := "リンゴ"
		englishWord := "apple"
		fmt.Println(japaneseWord + "を英語にしなさい。")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		// Removing blank spaces
		i := strings.TrimSpace(scanner.Text())

		if i == englishWord {
			fmt.Println("正解")
		} else {
			fmt.Println("不正解、正解は" + englishWord + "です。")
		}
	} else if *mode == "s" {
		fmt.Printf("This is a selective test.\n")
	}
}
