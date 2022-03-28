// Goのプログラムはパッケージ（package）で構成され、mainパッケージから開始する
package main

// importでパッケージをインポートする（ここではfmtとmath/rand）
import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}