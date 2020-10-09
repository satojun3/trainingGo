package registation

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// このパッケージ内でしか使用しない
func input() (productName string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		productName = strings.TrimSpace(scanner.Text())
		if productName != "" {
			break
		}
	}
	return
}

func incrementId() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func Register() int {
	fmt.Println("Register the product.")
	fmt.Println("Please enter product name.")
	getId := incrementId()
	fmt.Println(getId())
	fmt.Println(getId())
	fmt.Println(getId())
	fmt.Println(getId())
	return getId()
}
