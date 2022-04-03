package main

import (
	"fmt"
	"math/big"
	"testing"
)

func multiply(num1 string, num2 string) string {
	n1 := big.NewInt(0)
	n2 := big.NewInt(0)
	n1.SetString(num1, 10)
	n2.SetString(num2, 10)
	n1.Mul(n1, n2)
	return n1.String()
}

func TestBig(t *testing.T) {
	fmt.Printf("multiply(\"121212\", \"212121\"): %v\n", multiply("121212", "212121"))
	t.Fail()
}
