package main

import "fmt"

// varステートメントでは宣言時に初期化子（initializer）を与えることができる
var i, j int = 1, 2

func main() {
	// 初期化子が与えられた場合、型の記述を省略することができる => その変数は初期子が持つ型になる
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}