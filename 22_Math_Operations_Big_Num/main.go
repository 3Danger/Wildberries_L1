package main

import (
	"fmt"
	"math/big"
)

/*
	Разработать программу, которая перемножает,
	делит, складывает, вычитает две числовых переменных a,b,
	значение которых > 2^20.
*/

func main() {
	//  Поскольку обычный int64 не годится для большииих чисел, мы обратимся к пакету big в своей структуре
	bigInteger := NewBigInteger(922337203685477580, 302331203635477500)
	fmt.Println(bigInteger.Add())
	fmt.Println(bigInteger.Sub())
	fmt.Println(bigInteger.Mul())
	fmt.Println(bigInteger.Div())
}

type BigInteger struct {
	a, b *big.Int
}

func (b *BigInteger) Add() *big.Int {
	return new(big.Int).Add(b.a, b.b)
}

func (b *BigInteger) Sub() *big.Int {
	return new(big.Int).Sub(b.a, b.b)
}

func (b *BigInteger) Mul() *big.Int {
	return new(big.Int).Mul(b.a, b.b)
}

func (b *BigInteger) Div() *big.Int {
	return new(big.Int).Quo(b.a, b.b)
}

func NewBigInteger(a, b int64) *BigInteger {
	floats := &BigInteger{new(big.Int), new(big.Int)}
	floats.a.SetInt64(a)
	floats.b.SetInt64(b)
	return floats
}
