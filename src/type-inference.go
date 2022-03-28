package main

import "fmt"

func main() {
	v := 42  // change me!
	fmt.Printf("v is of type %T\n", v)
}

/*
明示的な型を指定せずに変数を宣言する場合（:= や var = のいずれか）、変数の型h右側の値もしくは変数から型推論される

型を指定しない数値である場合、変数は値の精度に基づいて、int、float64、complex128の型になる

i := 42            // int
f := 3.142         // float64
g := 0.867 + 0.5i  // complex128
*/