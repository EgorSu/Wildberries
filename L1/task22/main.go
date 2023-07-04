package main

import (
	"fmt"
	"math/big"
)

func add(val1, val2 *float64) *float64 {
	var res float64 = *val1 + *val2
	return &res
}
func mul(val1, val2 *float64) *float64 {
	var res float64 = *val1 * *val2
	return &res
}
func div(val1, val2 *float64) *float64 {
	var res float64 = *val1 / *val2
	return &res
}
func diff(val1, val2 *float64) *float64 {
	var res float64 = *val1 - *val2
	return &res
}

func main() {
	// с использованием float, но тогда потеряем точность
	a := float64(1 << 150)
	b := float64(1 << 40)

	fmt.Println(*add(&a, &b))
	fmt.Println(*mul(&a, &b))
	fmt.Println(*div(&a, &b))
	fmt.Println(*diff(&a, &b))

	//с использованием пакета big
	c := big.NewInt(int64(1<<63 - 1))
	d := big.NewInt(int64(1<<63 - 3))
	var result big.Int
	fmt.Println(result.Add(c, d))
	fmt.Println(result.Mul(c, d))
	fmt.Println(result.Div(c, d))
	fmt.Println(result.Sub(c, d))
}
