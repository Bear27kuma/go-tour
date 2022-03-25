package main

import "fmt"

// 関数の戻り値に名前をつけることができる => named return value
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// named return valueを使うと、return文に何も書かなくても良い（naked return）
	// 処理の長い関数でnaked returnを利用すると、コードの可読性が悪くなるため、短い関数でのみ利用すべき
	return
}

func main() {
	fmt.Println(split(17))
}