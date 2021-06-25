package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	rs1Letters []rune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
)

func main() {

	// TODO: OD
	// welcome文を作る

	// パスワードの長さを受け取る: int
	fmt.Println("パスワードの文字数を入力してください(１２文字以下):")

	// 誕生日の4桁数字を受け取る: int

	// passwordMaker(文字数, 誕生日)を呼び出してパスワードを生成し、表示する。
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
