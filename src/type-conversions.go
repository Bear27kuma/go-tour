package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

/*
変数v、型Tがあった場合、T(v)は、変数vを型Tへ変換する

例）
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

よりシンプルに記述することもできる
i := 42
f := float64(i)
u := uint(f)
*/