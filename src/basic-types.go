package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

/*
Go言語の基本型（組み込み型）

bool

string

// int, uint, uintptr型は、32-bitのシステムでは32bitで、64-bitのシステムでは64bit
// サイズ、符号なし（unsigned）整数を使うための特別な理由がない限り、intを使うようにする
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8の別名

rune // int32の別名、Unicodeのコードポイントを表す

float32 float64

complex64 complex128
*/