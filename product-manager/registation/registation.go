package registation

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	TAX           = 0.1
	BUSINESS_CODE = "123456789"
)

type Product struct {
	Id         int
	Name       string
	Price      int
	Tax        float32
	TotalPrice int
	janCode    string
}

// このパッケージ内でしか使用しない
func inputName() (productName string) {
	message := "Please enter product name."
	fmt.Println(message)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		productName = strings.TrimSpace(scanner.Text())
		if productName != "" {
			break
		}
		fmt.Println(message)
	}
	return
}

func inputPrice() (price int) {
	price, err := strconv.Atoi("")

	message := "Please enter the price."
	validationMessage1 := "Please Enter a number."
	validationMessage2 := "Please Enter 1 or more."

	fmt.Println(message)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() != "" {
			price, err = strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err == nil {
				if price > 0 {
					// 正常な値が入力されているため、ループから抜ける
					break
				}
				// 入力値が1未満のため再入力
				fmt.Println(validationMessage2)
			} else {
				// 数値が入力されていないため再入力
				fmt.Println(validationMessage1)
			}
		}
	}
	return
}

func incrementId(latestId int) int {
	return latestId + 1
}

func getJancode(id int) string {
	// 文字列を逆順にする
	reverseCode := reverse(BUSINESS_CODE)
	// ３桁0埋め
	zeroPaddingId := fmt.Sprintf("%03d", id)
	code := reverseCode + zeroPaddingId
	return code
}

func reverse(val string) string {
	var runeArr []rune
	for _, oneText := range val {
		runeArr = append(runeArr, oneText)
	}
	reverseText := ""
	for i := 1; i <= len(runeArr); i++ {
		reverseText += string(runeArr[len(runeArr)-i])
	}
	fmt.Println(reverseText)
	return reverseText
}

func (p Product) Record() string {
	data := "|" + strconv.Itoa(p.Id)
	data += "|" + p.Name
	data += "|" + strconv.Itoa(p.Price)
	data += "|" + strconv.FormatFloat(TAX, 'f', -1, 64)
	data += "|" + strconv.Itoa(p.TotalPrice)
	data += "|" + p.janCode
	data += "|"
	return data
}

func GetLabel() (data string) {
	data = "|id|name|price|tax|total price|JAN code|"
	return
}

// Register は構造体を返す
func Register(id int) Product {
	id = incrementId(id)
	name := inputName()
	price := inputPrice()
	totalPrice := int(float64(price)*TAX) + price
	jancode := getJancode(id)
	product := Product{
		Id:         id,
		Name:       name,
		Price:      price,
		Tax:        TAX,
		TotalPrice: totalPrice,
		janCode:    jancode,
	}
	return product
}
