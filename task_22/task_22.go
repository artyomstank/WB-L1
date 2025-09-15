package main

import (
	"fmt"
	"math/big"
)

func main() {
	//scope for a,b in range int64
	{
		var (
			a int64 = 1 << 1
			b int64 = 1 << 1
		)
		fmt.Println("int64")
		fmt.Println(b + a)
		fmt.Println(a - b)
		fmt.Println(b * a)
		fmt.Println(b / a)
	}
	//scope for a,b bigger than 2^20
	{
		var (
			a   = big.NewInt(1<<61 - 1)
			b   = big.NewInt(1<<56 - 2)
			res = big.NewInt(0)
		)
		fmt.Println("big int")
		fmt.Println(res.Add(a, b))
		fmt.Println(res.Sub(b.Neg(b), a))
		fmt.Println(res.Mul(a, b))
		fmt.Println(res.Div(b, a))

	}

}
