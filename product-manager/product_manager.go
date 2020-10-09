package main

import (
	"bufio"
	"fmt"
	"github.com/satojun3/trainingGo/product-manager/registation"
	"os"
	"strconv"
	// "strings"
)

type Product struct {
	Id      int
	Name    string
	Price   int
	janCode string
}

// メニューの表示
func displayMenu() {
	fmt.Println("1. Registation")
	fmt.Println("2. Comfirmation")
	fmt.Println("3. End")
}

// }
// func input() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	strings.TrimSpace(scanner.Text())
// }

// 入力値チェック

func main() {
	// メニューを表示
	displayMenu()

	// 値の入力
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputMenuNum, err := strconv.Atoi(scanner.Text())
		if err == nil && (inputMenuNum >= 1 && inputMenuNum <= 3) {
			switch inputMenuNum {
			case 1:
				// 登録機能
				fmt.Println("id:", registation.Register())
				fmt.Println("登録したよ")
			case 2:
				fmt.Printf("C")
			case 3:
				fmt.Println("Exit the system.")
				return
			}
		}
		displayMenu()
	}
}
