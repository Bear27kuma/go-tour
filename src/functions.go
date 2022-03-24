package main

import "fmt"

// 変数名の後ろに型名を記述する
func add(x int, y int) int {
	return x + y
}

func main() {
	// int型の引数を2つ取って、add関数を実行する
	fmt.Println(add(42, 13))
}