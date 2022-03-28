package main

import "fmt"

func main() {
	var i, j int = 1, 2
	// 関数の中では、var宣言の代わりに「:=」を代入文として使うことができ、暗黙的な型宣言ができる
	// 関数の外では、キーワードでの宣言（var, functionなど）が必要で、「:=」は利用できない
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}