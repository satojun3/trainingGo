package main

import (
	"bufio"
	"fmt"
	"github.com/satojun3/trainingGo/product-manager/registation"
	"os"
	"strconv"
)

// メニューの表示
func displayMenu() {
	fmt.Println("1. Registation")
	fmt.Println("2. Comfirmation")
	fmt.Println("3. End")
}

var (
	product     registation.Product
	productList []registation.Product
)

func main() {
	// メニューを表示
	displayMenu()

	// 値の入力
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		inputMenuNum, err := strconv.Atoi(scanner.Text()) // TODO: 何で再代入ができるのか調べる
		if err == nil && (inputMenuNum >= 1 && inputMenuNum <= 3) {
			switch inputMenuNum {
			case 1:
				// 登録機能
				product = registation.Register(len(productList))
				productList = append(productList, product)
				fmt.Println(registation.GetLabel())
				fmt.Println(product.Record())
			case 2:
				fmt.Println(registation.GetLabel())
				for _, product := range productList {
					fmt.Println(product.Record())
				}
			case 3:
				fmt.Println("Exit the system.")
				return
			}
		} else {
			fmt.Println("Please enter 1 - 3.")
		}
		displayMenu()
	}
}
