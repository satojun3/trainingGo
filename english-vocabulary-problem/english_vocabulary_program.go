package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	// "strings"
)

func main() {
	// オプションの設定 初期値は w
	var mode *string = flag.String("mode", "w", "This is a written test of English words.")
	// フラグの取得
	flag.Parse()

	// 分岐
	switch *mode {
	case "w":
		fmt.Printf("記述式の英単語プログラムです。\n")
	case "s":
		fmt.Printf("選択式のプログラムです。\n")
	default:
		fmt.Printf("オプションを間違えています。w か s を選択してください。")
	}

	fmt.Printf(inputAnswer("aaa"))

	// testScanner := bufio.NewScanner(os.Stdin)
	// bool := testScanner.Scan()
	// fmt.Println(bool)
	// for testScanner.Scan() {
	// 	if testScanner.Text() == "stop" {
	// 		break
	// 	}
	// 	fmt.Println("It's on an infinite loop.")
	// }
	// fmt.Println("end.")
}

func inputAnswer(correct string) string {
	var responseWord string = ""

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == correct {
			responseWord = "正解です。\n"
			break
		} else if scanner.Text() == "" {
			responseWord = "パスします。\n"
			break
		}
		fmt.Printf("不正解です。\n")
	}
	return responseWord
}
