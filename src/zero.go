package main

import "fmt"

func main() {
	// 変数に初期値を与えずに宣言すると、ゼロ値（zero value）が与えられる
	var i int		// 数値型（int, floatなど） => 0
	var f float64
	var b bool		// bool型 => false
	var s string	// string型 => ""（空文字列（empty string））
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}