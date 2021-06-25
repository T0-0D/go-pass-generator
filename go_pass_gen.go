package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	rs1Letters []rune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
)

func main() { // TODO: OD
	var pass_length int = int(^uint(0) >> 1)
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
		fmt.Print("Type your birthday in 4 letters (ex. Jan 3rd -> 0103): ")
		fmt.Scan(&birth)
	}

	// passwordMaker(文字数, 誕生日)を呼び出してパスワードを生成し、表示する。
	fmt.Print("Your password is here: ")
	fmt.Println(passwordMaker(pass_length, birth))
}

func passwordMaker(pass_length int, birth string) string {

	// ランダム関数の初期化
	rand.Seed(time.Now().UnixNano())
	var password []rune

	// ランダムな文字を一個一個入れていく
	for i := 0; i < pass_length; i++ {
		// appendは配列に要素を追加する、したものを再代入する
		password = append(password, rs1Letters[rand.Intn(len(rs1Letters))])
	}

	// もし最初の文字が大文字じゃなかったら大文字にする
	if password[0] <= 'A' || password[0] >= 'Z' {
		password[0] = rs1Letters[rand.Intn(26)]
	}

	return string(password) + birth
}
