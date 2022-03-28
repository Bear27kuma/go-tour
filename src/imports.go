package main

// 括弧でパッケージのインポートをグループ化 => Factored Import Statement
import (
	"fmt"
	"math"
)
// 以下のように記述することもできる
// import "fmt"
// import "math"

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}