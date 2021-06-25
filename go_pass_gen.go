package main

import (
	"fmt"
)

func main() { // TODO: OD
	var pass_length int = 13
	var birth string

	// welcome文を作る
	fmt.Println("Welcome to Golang Password Generator!")

	// パスワードの長さを受け取る: int
	for pass_length > 12 {
		fmt.Print("Type the length of your password (12 or less charactars): ")
		fmt.Scan(&pass_length)
	}

	// 誕生日の4桁数字を受け取る: int
	for len(birth) != 4 {
		fmt.Print("Type your birthday in 4 letter (ex. Jan 3rd -> 0103): ")
		fmt.Scan(&birth)
	}

	// passwordMaker(文字数, 誕生日)を呼び出してパスワードを生成し、表示する。
	fmt.Print("Your password is here: ")
	fmt.Println(passwordMaker(pass_length, birth))
}

func passwordMaker(pass_length int, birth string) string {
	// TODO: RYOMA
	// パスワードを生成する！
	return "PASSWORD"
}
