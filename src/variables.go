package main

import "fmt"

// varステートメントで変数を宣言
// 関数の引数同様に、複数の変数の最後に型を記述することで、同じ型の変数を一度に宣言できる
var c, python, java bool

// varステートメントはパッケージ、または関数で利用できる
func main() {
	var i int
	fmt.Println(i, c, python, java)
}