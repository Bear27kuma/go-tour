package main

import "fmt"

// 2つ以上の引数が同じ型である場合、最後の型を残して記述する
// x int, y int => x, y int
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}