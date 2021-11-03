package bit_operation

import (
	"fmt"
	"math/big"
)

func addBinary(a string, b string) string {

	ai, _ := new(big.Int).SetString(a, 2)
	bi, _ := new(big.Int).SetString(b, 2)
	fmt.Println(ai)
	ai.Add(ai, bi)
	return ai.Text(2)
}

func main() {
	fmt.Println(addBinary("101", "110"))
}
