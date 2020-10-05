package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// Problem ...
type Problem struct {
	ID          int    `json:"id"`
	Word        string `json:"word"`
	Meaning     string `json:"meaning"`
	Description string `json:"description"`
}

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
	problems := getProblems()
	for _, problem := range problems {
		fmt.Printf(problem.Meaning + " を英語で？\n")
		fmt.Printf(inputAnswer(problem.Word))
	}
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

// JSONファイルの読み込み
func getProblems() []Problem {
	bytes, err := ioutil.ReadFile("problems.json")
	if err != nil {
		fmt.Printf("error:jsonファイルが読み込めない")
	}
	var problems []Problem
	if err := json.Unmarshal(bytes, &problems); err != nil {
		fmt.Printf("error:読み込んだjsonデータを変換できません")
	}
	return problems
}
