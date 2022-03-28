package main

import "fmt"

// 定数（constant）は、constキーワードを使って変数と同じように宣言する
const Pi = 3.14

func main() {
	// 定数は、文字（character）、文字列（string）、boolean、数値（numeric）のみで使える
	// 定数は := を使って宣言できない
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}